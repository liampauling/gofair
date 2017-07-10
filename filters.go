package gofair


type TimeRangeFilter struct {
	From 	string 	`json:"from,omitempty"`
	To 	string 	`json:"to,omitempty"`
}


type MarketFilter struct {
	TextQuery		string			`json:"textQuery,omitempty"`
	EventTypeIds		[]string		`json:"eventTypeIds,omitempty"`
	MarketCountries		[]string		`json:"marketCountries,omitempty"`
	MarketIds		[]string		`json:"marketIds,omitempty"`
	EventIds		[]string		`json:"eventIds,omitempty"`
	CompetitionIds		[]string		`json:"competitionIds,omitempty"`
	BSPOnly			bool			`json:"bspOnly,omitempty"`
	TurnInPlayEnabled	bool			`json:"turnInPLayEnabled,omitempty"`
	InPlayOnly		bool			`json:"inPlayOnly,omitempty"`
	MarketBettingTypes	[]string		`json:"marketBettingTypes,omitempty"`
	MarketTypeCOdes		[]string		`json:"marketTypeCOdes,omitempty"`
	MarketStartTime		TimeRangeFilter		`json:"marketStartTime,omitempty"`
	WithOrders		string			`json:"withOrders,omitempty"`
}


type Params struct {
	MarketFilter		MarketFilter 		`json:"filter,omitempty"`
	MaxResults		int			`json:"maxResults,omitempty"`
	Granularity		string			`json:"granularity,omitempty"`
	MarketProjection	[]string		`json:"marketProjection,omitempty"`
	Sort			string			`json:"sort,omitempty"`
	Locale			string			`json:"locale,omitempty"`
}
