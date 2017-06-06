package gofair


type MarketFilter struct {
	TextQuery		string		`json:"textQuery,omitempty"`
	ExchangeIds		[]string	`json:"exchangeIds,omitempty"`
	EventTypeIds		[]string	`json:"eventTypeIds,omitempty"`
	MarketCountries		[]string	`json:"marketCountries,omitempty"`
	MarketIds		[]string	`json:"marketIds,omitempty"`
}


type Params struct {
	MarketFilter		*MarketFilter 		`json:"filter,omitempty"`
	MaxResults		int			`json:"maxResults,omitempty"`
	Locale			string			`json:"locale,omitempty"`
}
