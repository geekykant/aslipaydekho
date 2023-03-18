package fetchapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/geekykant/aslipaydekho/scraper/model"
	"github.com/geekykant/aslipaydekho/scraper/utils"
	"github.com/sethgrid/pester"
)

func FetchLeetCodeCompensationPosts(client *pester.Client) {
	reqPostOffset, reqPostCount := 0, 10
	post, err := fetchPostContent(client, reqPostOffset, reqPostCount)
	if err != nil {
		panic("Some error occoured - " + err.Error())
	}

	//Print the fetched details
	fmt.Println("[*] Found post with post id" + post.ID)
	fmt.Println("[*] URL is https://leetcode.com/discuss/compensation/" + post.ID)
	fmt.Println(post.PostContent)

	//Print after parsing the data
	offerLetter := model.ParsePostContent(post)
	model.PrintOfferLetter(&offerLetter)

	//Insert both - Parsed, Unparsed into MQ
}

func fetchPostContent(client *pester.Client, reqPostOffset int, reqPostCount int) (*model.Post, error) {
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

	post_meta := respData.Data.CategoryTopicList.Edges[7].Node
	post_meta.Post.ID = post_meta.ID

	utils.CleanCompensationPostContent(&post_meta.Post.PostContent)
	return &post_meta.Post, nil
}
