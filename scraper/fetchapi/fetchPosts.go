package fetchapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/geekykant/aslipaydekho/scraper/model"
	"github.com/geekykant/aslipaydekho/scraper/utils"
)

func StartFetchInsertNewCompensationPost() {
	reqPostOffset, reqPostCount := 0, 100

	allPostAndOfferLetters, err := FetchCompensationPostsInRange(reqPostOffset, reqPostCount)
	if err != nil {
		err := fmt.Errorf("Error while fetching AllPostAndOfferLetters - " + err.Error())
		panic(err)
	}

	//Insert all posts/offerletters to MQ
	fmt.Println(allPostAndOfferLetters)
}

func FetchCompensationPostsInRange(reqPostOffset int, reqPostCount int) ([]model.PostAndOfferLetter, error) {
	allPosts, err := fetchCompensationPosts(reqPostOffset, reqPostCount)
	if err != nil {
		return nil, err
	}

	allPostAndOfferLetters := make([]model.PostAndOfferLetter, len(allPosts))
	for i := 0; i < len(allPosts); i++ {
		postAndOfferLetter := new(model.PostAndOfferLetter)

		//Basic cleans post content - removing markdowns
		postAndOfferLetter.PostUrl = fmt.Sprintf("https://leetcode.com/discuss/compensation/" + allPosts[i].ID)
		postAndOfferLetter.OriginalPost = allPosts[i]
		utils.BasicCleanCompensationPostContent(&allPosts[i].PostContent)

		//Parsing begins - to extract useful details from the post
		postAndOfferLetter.ParsedOfferLetter = model.ParsePostContent(&allPosts[i])

		allPostAndOfferLetters[i] = *postAndOfferLetter
	}

	return allPostAndOfferLetters, nil
}

func fetchCompensationPosts(reqPostOffset int, reqPostCount int) ([]model.Post, error) {
	client := GetPesterCientInstance()
	url := "https://leetcode.com/graphql"
	payload := fmt.Sprintf(
		`{"operationName":"categoryTopicList","variables":{"orderBy":"newest_to_oldest","query":"","skip":%d,"first":%d,"tags":[],"categories":["compensation"]},"query":"query categoryTopicList($categories: [String!]!, $first: Int!, $orderBy: TopicSortingOption, $skip: Int, $query: String, $tags: [String!]) {\n  categoryTopicList(categories: $categories, orderBy: $orderBy, skip: $skip, query: $query, first: $first, tags: $tags) {\n    ...TopicsList\n    __typename\n  }\n}\n\nfragment TopicsList on TopicConnection {\n  totalNum\n  edges {\n    node {\n      id\n      title\n      commentCount\n      viewCount\n      pinned\n      tags {\n        name\n        slug\n        __typename\n      }\n      post {\n        id\n        content\n        voteCount\n        creationDate\n        isHidden\n        author {\n          username\n          isActive\n          nameColor\n          activeBadge {\n            displayName\n            icon\n            __typename\n          }\n          profile {\n            userAvatar\n            __typename\n          }\n          __typename\n        }\n        status\n        coinRewards {\n          ...CoinReward\n          __typename\n        }\n        __typename\n      }\n      lastComment {\n        id\n        post {\n          id\n          author {\n            isActive\n            username\n            __typename\n          }\n          peek\n          creationDate\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    cursor\n    __typename\n  }\n  __typename\n}\n\nfragment CoinReward on ScoreNode {\n  id\n  score\n  description\n  date\n  __typename\n}"}`,
		reqPostOffset,
		reqPostCount,
	)

	payloadBytes := []byte(payload)
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
	allPosts := make([]model.Post, fetchedPostsCount)

	for i := 0; i < fetchedPostsCount; i++ {
		post_meta := respData.Data.CategoryTopicList.Edges[i].Node
		post_meta.Post.ID = post_meta.ID
		allPosts[i] = post_meta.Post
	}

	return allPosts, nil
}
