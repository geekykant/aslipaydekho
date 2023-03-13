package utils

func GetOfferLetterParsingPattern() map[string]string {
	return map[string]string{
		"Education":                `(?i)education:\s*(.+)`,
		"YearsOfExperience":        `(?i)years of experience:\s*(.+)`,
		"PreviousJobTitle":         `(?i)prior experience:\s*(.+)`,
		"NewJobTitle":              `(?i)date of the offer:\s*(.+)`,
		"DateOfOffer":              `(?i)company:\s*(.+)`,
		"Company":                  `(?i)title/level:\s*(.+)`,
		"Location":                 `(?i)location:\s*(.+)`,
		"Salary":                   `(?i)salary:\s*(.+)`,
		"RelocationOrSigningBonus": `(?i)relocation/signing bonus:\s*(.+)`,
		"StockBonus":               `(?i)stock bonus:\s*(.+)`,
		"Bonus":                    `(?i)bonus:\s*(.+)`,
		"TotalCompensation":        `(?i)total comp \(salary \+ bonus \+ stock\):\s*(.+)`,
		"Benefits":                 `(?i)benefits:\s*(.+)`,
		"OtherDetails":             `(?i)other details:\s*(.+)`,
	}
}
