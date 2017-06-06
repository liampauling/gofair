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


type eventType struct {
	Id	string	`json:"id"`
	Name	string 	`json:"name"`
}


type eventTypeResult struct {
	MarketCount	int		`json:"marketCount"`
	EventType	*eventType 	`json:"eventType"`
}


func (b *Betting) ListEventTypes(filter *MarketFilter) ([]eventTypeResult, error) {
	// create url
	url := createUrl(api_url, "betting/json-rpc/v1")

	// build request
	method := "SportsAPING/v1.0/" + "listEventTypes"
	params := new(Params)
	params.MarketFilter = filter

	var results []eventTypeResult

	// make request
	b.Request(url, method, params, &results)
	return results, nil
}
