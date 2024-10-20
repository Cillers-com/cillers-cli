package lib

import (
	"encoding/xml"
	"fmt"
)

type ChangeProposal struct {
	XMLName      xml.Name     `xml:"change-proposal"`
	Description  Description  `xml:"description"`
	Specification Specification `xml:"specification"`
	CodeReview   CodeReview   `xml:"code-review"`
}

type Description struct {
	ChangeSummary string   `xml:"change-summary"`
	ChangeDetails []string `xml:"change-detail"`
}

type Specification struct {
	FilesToBeCreated []File `xml:"file-to-be-created"`
	FilesToBeUpdated []File `xml:"file-to-be-updated"`
	FilesToBeDeleted []File `xml:"file-to-be-deleted"`
}

type File struct {
	Path    string `xml:"path,attr"`
	Content string `xml:",chardata"`
}

type CodeReview struct {
	PositiveFeedback        string `xml:"positive_feedback"`
	ImprovementSuggestions  string `xml:"improvement_suggestions"`
	CodeQualityAssessment   string `xml:"code_quality_assessment"`
}

func ParseAnthropicResponse(xmlResponse string) (*ChangeProposal, error) {
	var proposal ChangeProposal
	err := xml.Unmarshal([]byte(xmlResponse), &proposal)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling XML: %w", err)
	}
	return &proposal, nil
}
