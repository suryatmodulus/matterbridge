// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RubricQualityFeedbackModel undocumented
type RubricQualityFeedbackModel struct {
	// Object is the base model of RubricQualityFeedbackModel
	Object
	// QualityID undocumented
	QualityID *string `json:"qualityId,omitempty"`
	// Feedback undocumented
	Feedback *EducationItemBody `json:"feedback,omitempty"`
}