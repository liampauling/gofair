package gofair

import "time"

type eventType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type eventTypeResult struct {
	MarketCount int       `json:"marketCount"`
	EventType   eventType `json:"eventType"`
}

type competition struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type competitionResult struct {
	MarketCount       int         `json:"marketCount"`
	CompetitionRegion string      `json:"competitionRegion"`
	Competition       competition `json:"competition"`
}

type timeRange struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

type timeRangeResult struct {
	MarketCount int       `json:"marketCount"`
	TimeRange   timeRange `json:"timeRange"`
}

type event struct {
	Id          string `json:"id"`
	OpenDate    string `json:"openDate"`
	TimeZone    string `json:"timezone"`
	CountryCode string `json:"countryCode"`
	Name        string `json:"name"`
	Venue       string `json:"venue"`
}

type eventResult struct {
	MarketCount int   `json:"marketCount"`
	Event       event `json:"event"`
}

type marketTypeResult struct {
	MarketCount int    `json:"marketCount"`
	MarketType  string `json:"marketType"`
}

type countryResult struct {
	MarketCount int    `json:"marketCount"`
	CountryCode string `json:"countryCode"`
}

type venueResult struct {
	MarketCount int    `json:"marketCount"`
	Venue       string `json:"venue"`
}

type marketCatalogueDescription struct {
	BettingType        string    `json:"bettingType"`
	BSPMarket          bool      `json:"bspMarket"`
	DiscountAllowed    bool      `json:"discountAllowed"`
	MarketBaseRate     float32   `json:"marketBaseRate"`
	MarketTime         time.Time `json:"marketTime"`
	MarketType         string    `json:"marketType"`
	PersistenceEnabled bool      `json:"persistenceEnabled"`
	Regulator          string    `json:"regulator"`
	Rules              string    `json:"rules"`
	RulesHasDate       bool      `json:"rulesHasDate"`
	SuspendDate        time.Time `json:"suspendTime"`
	TurnInPlayEnabled  bool      `json:"turnInPlayEnabled"`
	Wallet             string    `json:"wallet"`
	EachWayDivisor     float32   `json:"eachWayDivisor"`
	Clarifications     string    `json:"clarifications"`
}

type metadata struct {
	RunnerId int `json:"runnerId"`
}

type runnerCatalogue struct {
	SelectionId  int     `json:"selectionId"`
	RunnerName   string  `json:"runnerName"`
	SortPriority int     `json:"sortPriority"`
	Handicap     float32 `json:"handicap"`
	//Metadata		*metadata	`json:"metadata"`  //todo
}

type marketCatalogue struct {
	MarketId                   string                     `json:"marketId"`
	MarketName                 string                     `json:"marketName"`
	TotalMatched               float32                    `json:"totalMatched"`
	MarketStartTime            time.Time                  `json:"marketStartTime"`
	Competition                competition                `json:"competition"`
	Event                      event                      `json:"event"`
	EventType                  eventType                  `json:"eventType"`
	MarketCatalogueDescription marketCatalogueDescription `json:"description"`
	Runners                    []runnerCatalogue          `json:"runners"`
}

func (b *Betting) ListEventTypes(filter MarketFilter) ([]eventTypeResult, error) {
	// create url
	url := createUrl(api_betting_url, "listEventTypes/")

	// build request
	params := new(Params)
	params.MarketFilter = filter

	var response []eventTypeResult

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (b *Betting) ListCompetitions(filter MarketFilter) ([]competitionResult, error) {
	// create url
	url := createUrl(api_betting_url, "listCompetitions/")

	// build request
	params := new(Params)
	params.MarketFilter = filter

	var response []competitionResult

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (b *Betting) ListTimeRanges(filter MarketFilter, granularity string) ([]timeRangeResult, error) {
	// create url
	url := createUrl(api_betting_url, "listTimeRanges/")

	// build request
	params := new(Params)
	params.MarketFilter = filter
	params.Granularity = granularity

	var response []timeRangeResult

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (b *Betting) ListEvents(filter MarketFilter) ([]eventResult, error) {
	// create url
	url := createUrl(api_betting_url, "listEvents/")

	// build request
	params := new(Params)
	params.MarketFilter = filter

	var response []eventResult

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (b *Betting) ListMarketTypes(filter MarketFilter) ([]marketTypeResult, error) {
	// create url
	url := createUrl(api_betting_url, "listMarketTypes/")

	// build request
	params := new(Params)
	params.MarketFilter = filter

	var response []marketTypeResult

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (b *Betting) ListCountries(filter MarketFilter) ([]countryResult, error) {
	// create url
	url := createUrl(api_betting_url, "listCountries/")

	// build request
	params := new(Params)
	params.MarketFilter = filter

	var response []countryResult

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (b *Betting) ListVenues(filter MarketFilter) ([]venueResult, error) {
	// create url
	url := createUrl(api_betting_url, "listVenues/")

	// build request
	params := new(Params)
	params.MarketFilter = filter

	var response []venueResult

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (b *Betting) ListMarketCatalogue(filter MarketFilter, marketProjection []string, sort string, maxResults int) (
	[]marketCatalogue, error) {
	// create url
	url := createUrl(api_betting_url, "listMarketCatalogue/")

	// build request
	params := new(Params)
	params.MarketFilter = filter
	params.MarketProjection = marketProjection
	params.Sort = sort
	params.MaxResults = maxResults

	var response []marketCatalogue

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}
