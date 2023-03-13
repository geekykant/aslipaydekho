package model

type DataNode struct {
	Data Data `json:"data"`
}

type Data struct {
	CategoryTopicList CategoryTopicList `json:"categoryTopicList"`
}

type CategoryTopicList struct {
	Edges []Edge `json:"edges"`
}

type Edge struct {
	Node Node `json:"node"`
}

type Node struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Post Post   `json:"post"`
}

type Post struct {
	ID          string `json:"-"`
	PostContent string `json:"content"`
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
