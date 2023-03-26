package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/geekykant/aslipaydekho/scraper/model"
	"github.com/geekykant/aslipaydekho/scraper/utils"
)

func FetchPopulateCompensationsSinceLastWeek() error {
	lastFetchedTime := time.Now().AddDate(0, 0, -7) //last 7 days (1 week)

	batchFetchCount := uint16(500)
	reqPostOffset := uint16(0)

	//Counter to skip sticky posts (interferes with its old date)
	counterBeforeLeaving := 3

	for {
		paolList, err := FetchCompensationPostsInRange(reqPostOffset, batchFetchCount)
		if err != nil {
			err = fmt.Errorf("Error while fetching - " + err.Error())
			return err
		}

		if len(paolList) == 0 {
			return nil
		}

		for _, paol := range paolList {
			if paol.OriginalPost.CreationDate.Before(lastFetchedTime) {
				counterBeforeLeaving--
				if counterBeforeLeaving == 0 {
					return nil
				}
				continue
			}

			// Insert both - Parsed, Unparsed into MQ
			err := SendOfferLetterToMQ(&paol)
			if err != nil {
				panic(err)
			}
			// fmt.Printf("fetched post created on %s \n", paol.OriginalPost.CreationDate)
		}

		reqPostOffset += batchFetchCount
	}
}

func InitPopulateAllCompensationsToMQ() {
	totalCompensationPostCount, err := fetchCompensationPostsCount()
	if err != nil {
		panic("Aiyoo error - Couldn't get the total compensation post count - " + err.Error())
	}

	fmt.Printf("[*] Found Total Compensations Posts: %d \n", totalCompensationPostCount)

	//Create a channel to listen on go routine calls
	ch := make(chan model.PostAndOfferLetter, 1)
	if err := FetchAllCompensationPosts(ch, totalCompensationPostCount); err != nil {
		panic("Error occoured" + err.Error())
	}

	for i := 0; i < int(totalCompensationPostCount); i++ {
		paol := <-ch
		// Insert both - Parsed, Unparsed into MQ
		err := SendOfferLetterToMQ(&paol)
		if err != nil {
			panic(err)
		}
	}
}

func FetchAllCompensationPosts(ch chan model.PostAndOfferLetter, totalCompensationPostCount uint16) error {
	batchFetchCount := uint16(500)
	reqPostOffset, reqPostCount := uint16(0), batchFetchCount

	for reqPostOffset < totalCompensationPostCount {
		go func(offset uint16, reqCount uint16) {
			paolList, err := FetchCompensationPostsInRange(offset, reqCount)
			if err != nil {
				fmt.Printf("Error while fetching in range %d - %d - "+err.Error()+"\n", offset, offset+reqCount)
				return
			}

			fmt.Printf("fetched %d to %d - totalCount - %d \n", offset, offset+reqCount, len(paolList))
			for _, paol := range paolList {
				ch <- paol
			}

		}(reqPostOffset, reqPostCount)

		reqPostOffset += batchFetchCount
	}

	return nil
}

func FetchCompensationPostsInRange(reqPostOffset uint16, reqPostCount uint16) ([]model.PostAndOfferLetter, error) {
	allPosts, err := fetchCompensationPosts(reqPostOffset, reqPostCount)
	if err != nil {
		return nil, err
	}

	allPostAndOfferLetters := make([]model.PostAndOfferLetter, len(allPosts))
	for i := 0; i < len(allPosts); i++ {
		postAndOfferLetter := new(model.PostAndOfferLetter)

		//Basic cleans post content - removing markdowns
		postAndOfferLetter.PostUrl = fmt.Sprintf("https://leetcode.com/discuss/compensation/" + allPosts[i].ID)
		postAndOfferLetter.OriginalPost = model.PostOutput{
			ID:           allPosts[i].ID,
			CreationDate: utils.GetDateTimeFromEpochMillis(allPosts[i].CreationDate),
			PostContent:  allPosts[i].PostContent,
		}

		//Cleaning & Parsing begins - to extract useful details from the post
		utils.BasicCleanCompensationPostContent(&allPosts[i].PostContent)
		postAndOfferLetter.ParsedOfferLetter = model.ParsePostContent(&allPosts[i])

		allPostAndOfferLetters[i] = *postAndOfferLetter
	}

	return allPostAndOfferLetters, nil
}

func fetchCompensationPosts(reqPostOffset uint16, reqPostCount uint16) ([]model.PostInput, error) {
	client := GetPesterCientInstance()
	url := "https://leetcode.com/graphql"
	variables := map[string]interface{}{
		"categories": []string{"compensation"},
		"first":      reqPostCount,
		"orderBy":    "newest_to_oldest",
		"skip":       reqPostOffset,
		"query":      "",
		"tags":       []string{},
	}

	payload := model.GraphqlRequest{
		OperationName: "categoryTopicList",
		Query: `query categoryTopicList($categories: [String!]!, $first: Int!, $orderBy: TopicSortingOption, $skip: Int, $query: String, $tags: [String!]) {
			categoryTopicList(categories: $categories, orderBy: $orderBy, skip: $skip, query: $query, first: $first, tags: $tags) {
				...TopicsList
				__typename
			}
		}
		fragment TopicsList on TopicConnection { edges { node { id title commentCount viewCount pinned post { id content voteCount creationDate isHidden author { username isActive nameColor activeBadge { displayName icon __typename } __typename } status __typename } __typename } cursor __typename } __typename }`,
		Variables: variables,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	res, err := client.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		err := fmt.Errorf("Error sending request:" + err.Error())
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		err := fmt.Errorf("Error reading response body:" + err.Error())
		return nil, err
	}

	if res.StatusCode != 200 {
		err := fmt.Errorf("Status code was not 200, Instead " + fmt.Sprint(res.StatusCode))
		return nil, err
	}

	var respData model.DataNode
	unmarshErr := json.Unmarshal(body, &respData)
	if unmarshErr != nil {
		return nil, unmarshErr
	}

	fetchedPostsCount := len(respData.Data.CategoryTopicList.Edges)
	allPosts := make([]model.PostInput, fetchedPostsCount)

	for i := 0; i < fetchedPostsCount; i++ {
		post_meta := respData.Data.CategoryTopicList.Edges[i].Node
		post_meta.Post.ID = post_meta.ID
		allPosts[i] = post_meta.Post
	}

	return allPosts, nil
}

func fetchCompensationPostsCount() (uint16, error) {
	client := GetPesterCientInstance()
	url := "https://leetcode.com/graphql"
	variables := map[string]interface{}{
		"categories": []string{"compensation"},
		"first":      15,
		"orderBy":    "newest_to_oldest",
		"skip":       0,
		"query":      "",
		"tags":       []string{},
	}

	reqBody := model.GraphqlRequest{
		OperationName: "categoryTopicList",
		Query: `query categoryTopicList($categories: [String!]!, $first: Int!, $orderBy: TopicSortingOption, $skip: Int, $query: String, $tags: [String!]) {
          categoryTopicList(categories: $categories, orderBy: $orderBy, skip: $skip, query: $query, first: $first, tags: $tags) {
            totalNum
          }
        }`,
		Variables: variables,
	}

	payloadBytes, err := json.Marshal(reqBody)
	if err != nil {
		return 0, err
	}

	res, err := client.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		err := fmt.Errorf("Error sending request:" + err.Error())
		return 0, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		err := fmt.Errorf("Error reading response body:" + err.Error())
		return 0, err
	}

	if res.StatusCode != 200 {
		err := fmt.Errorf("Status code was not 200, Instead " + fmt.Sprint(res.StatusCode))
		return 0, err
	}

	var respData model.DataNode
	unmarshErr := json.Unmarshal(body, &respData)
	if unmarshErr != nil {
		return 0, unmarshErr
	}

	return respData.Data.CategoryTopicList.TotalNum, nil
}
