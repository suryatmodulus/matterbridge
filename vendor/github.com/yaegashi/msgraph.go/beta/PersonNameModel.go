// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PersonName undocumented
type PersonName struct {
	// ItemFacet is the base model of PersonName
	ItemFacet
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// First undocumented
	First *string `json:"first,omitempty"`
	// Initials undocumented
	Initials *string `json:"initials,omitempty"`
	// Last undocumented
	Last *string `json:"last,omitempty"`
	// LanguageTag undocumented
	LanguageTag *string `json:"languageTag,omitempty"`
	// Maiden undocumented
	Maiden *string `json:"maiden,omitempty"`
	// Middle undocumented
	Middle *string `json:"middle,omitempty"`
	// Nickname undocumented
	Nickname *string `json:"nickname,omitempty"`
	// Suffix undocumented
	Suffix *string `json:"suffix,omitempty"`
	// Title undocumented
	Title *string `json:"title,omitempty"`
	// Pronunciation undocumented
	Pronunciation *YomiPersonName `json:"pronunciation,omitempty"`
}