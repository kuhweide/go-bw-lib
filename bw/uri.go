package bw

type MatchType int

const (
	MatchTypeDomain MatchType = iota
	MatchTypeHost
	MatchTypeStartsWith
	MatchTypeExact
	MatchTypeRegularExpression
	MatchTypeNever
)

type uri struct {
	MatchType MatchType `json:"match"`
	Uri       string    `json:"uri"`
}

func NewUri(matchType MatchType, uriString string) uri {
	return uri{MatchType: matchType, Uri: uriString}
}
