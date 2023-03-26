package model

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/geekykant/aslipaydekho/scraper/utils"
)

type GraphqlRequest struct {
	OperationName string                 `json:"operationName"`
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
}

type DataNode struct {
	Data Data `json:"data"`
}

type Data struct {
	CategoryTopicList CategoryTopicList `json:"categoryTopicList"`
}

type CategoryTopicList struct {
	Edges    []Edge `json:"edges"`
	TotalNum uint16 `json:"totalNum"`
}

type Edge struct {
	Node Node `json:"node"`
}

type Node struct {
	ID   string    `json:"id"`
	Name string    `json:"name"`
	Post PostInput `json:"post"`
}

type PostInput struct {
	ID           string `json:"-"`
	CreationDate int64  `json:"creationDate"`
	PostContent  string `json:"content"`
}
type PostOutput struct {
	ID           string    `json:"id"`
	CreationDate time.Time `json:"creationDate"`
	PostContent  string    `json:"content"`
}

// Education: Tier 1 College
// Years of Experience: 2
// Prior Experience: Software Engineer
// Date of the Offer: Feb 2022
// Company: TranZact
// Title/Level: SDE-2
// Location: Mumbai
// Salary: 18,00,000
// Relocation/Signing Bonus: 2,00,000
// Stock bonus: 4 LPA (over 4 years)
// Bonus: 2,00,000
// Total comp (Salary + Bonus + Stock): 24 LPA (1 L stocks per year)

type OfferLetter struct {
	Education                string
	YearsOfExperience        string
	PreviousJobTitle         string
	NewJobTitle              string
	DateOfOffer              string
	Company                  string
	Location                 string
	Salary                   string
	RelocationOrSigningBonus string
	StockBonus               string
	Bonus                    string
	TotalCompensation        string
	Benefits                 string
	OtherDetails             string
}

type PostAndOfferLetter struct {
	PostUrl           string      `json:"url"`
	OriginalPost      PostOutput  `json:"post"`
	ParsedOfferLetter OfferLetter `json:"offerLetter"`
}

func ParsePostContent(post *PostInput) OfferLetter {
	offer := OfferLetter{}
	attrPatterns := utils.GetOfferLetterParsingPattern()

	for field, pattern := range attrPatterns {
		re := regexp.MustCompile(pattern)
		match := re.FindStringSubmatch(post.PostContent)
		if match != nil {
			fieldValue := reflect.ValueOf(&offer).Elem().FieldByName(field)
			fieldValue.SetString(strings.TrimSpace(match[1]))
		}
	}

	return offer
}

func PrintOfferLetter(offerLetter *OfferLetter) {
	value := reflect.ValueOf(offerLetter)

	// Iterate over each field of the struct
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldName := value.Type().Field(i).Name

		// Print the field name and value
		fmt.Printf("%s: %v\n", fieldName, field.Interface())
	}
}
