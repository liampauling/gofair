package gofair


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
	url := createUrl(api_betting_url, "listEventTypes/")

	// build request
	params := new(Params)
	params.MarketFilter = filter

	var results []eventTypeResult

	// make request
	err := b.Request(url, params, &results)

	return results, err
}
