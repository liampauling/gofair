package streaming

import "log"


func CreateMarketCache(changeMessage MarketChangeMessage, marketChange MarketChange)(MarketCache){
	var cache MarketCache
	cache.PublishTime = changeMessage.PublishTime
	cache.MarketId = marketChange.MarketId
	cache.TradedVolume = marketChange.TradedVolume
	cache.MarketDefinition = marketChange.MarketDefinition
	cache.Runners = make(map[int64]RunnerCache)

	for _, runnerChange := range *marketChange.RunnerChange {
		cache.Runners[runnerChange.SelectionId] = CreateRunnerCache(runnerChange)
	}
	return cache
}


func CreateRunnerCache(change RunnerChange)(RunnerCache){
	log.Println("Created new runner cache", change.SelectionId)
	var cache RunnerCache
	cache.SelectionId = change.SelectionId
	cache.LastTradedPrice = change.LastTradedPrice
	cache.TradedVolume = change.TradedVolume
	//cache.StartingPriceNear = change.StartingPriceNear
	//cache.StartingPriceFar = change.StartingPriceFar

	// create traded data structure
	var traded Available
	for _, i := range change.Traded {
		traded.Prices = append(
			traded.Prices,
			PriceSize{i[0], i[1]},
		)
	}
	traded.DeletionSelect = 1
	traded.Reverse = false
	cache.Traded = &traded

	return cache
}


type PriceSize struct {
	Price 		float64
	Size 		float64
}


type PositionPriceSize struct {
	Position 	int32
	Price 		float64
	Size 		float64
}


type Available struct {
	Prices		[]PriceSize
	DeletionSelect	int32
	Reverse		bool
}


func (available *Available) Clear(){
	available.Prices = nil
}


func (available *Available) Sort(){
	// todo
}


func (available *Available) Update(updates [][]float64){
	for _, update := range updates {
		updated := false
		for count, trade := range available.Prices {
			if trade.Price == update[0] {
				if update[available.DeletionSelect] == 0 {
					available.Prices = remove(available.Prices, count)
					updated = true
					break
				} else {
					available.Prices[count] = PriceSize{update[0], update[1]}
					updated = true
					break
				}
			}
		}
		if updated == false && update[available.DeletionSelect] != 0 {
			available.Prices = append(available.Prices, PriceSize{update[0], update[1]})
		}
	}
}


func remove(s []PriceSize, i int) []PriceSize {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}


type RunnerCache struct {
	SelectionId 	int64
	LastTradedPrice *float64
	TradedVolume 	*float64
	//StartingPriceNear *float64
	//StartingPriceFar *float64
	Traded 		*Available
}


func (cache *RunnerCache) UpdateCache(change RunnerChange) {
	if cache.SelectionId == 10631117 {
		log.Println("new", change.Traded, len(cache.Traded.Prices))
	}
	if change.LastTradedPrice != nil {
		*cache.LastTradedPrice = *change.LastTradedPrice
	}
	if change.TradedVolume != nil {
		*cache.TradedVolume = *change.TradedVolume
	}
	//if change.StartingPriceNear != nil {
	//	*cache.StartingPriceNear = *change.StartingPriceNear
	//}
	//if change.StartingPriceFar != nil {
	//	*cache.StartingPriceFar = *change.StartingPriceFar
	//}
	if change.Traded != nil {
		cache.Traded.Update(change.Traded)
	}
	if cache.SelectionId == 10631117 {
		log.Println(len(cache.Traded.Prices))
	}
}


type MarketCache struct {
	PublishTime		int64
	MarketId		string
	TradedVolume		*float64
	MarketDefinition	*MarketDefinition
	Runners 		map[int64]RunnerCache
}


func (cache *MarketCache) UpdateCache(changeMessage MarketChangeMessage, marketChange MarketChange) {
	cache.PublishTime = changeMessage.PublishTime

	if marketChange.MarketDefinition != nil {
		cache.MarketDefinition = marketChange.MarketDefinition
	}
	if marketChange.TradedVolume != nil {
		cache.TradedVolume = marketChange.TradedVolume
	}
	if marketChange.RunnerChange != nil {
		for _, runnerChange := range *marketChange.RunnerChange {
			if runnerCache, ok := cache.Runners[runnerChange.SelectionId]; ok {
				runnerCache.UpdateCache(runnerChange)
			} else {
				cache.Runners[runnerChange.SelectionId] = CreateRunnerCache(runnerChange)
			}
		}
	}
	tem, _ := cache.Runners[10631117]
	log.Println(tem.SelectionId, *tem.LastTradedPrice, *tem.TradedVolume, len(tem.Traded.Prices))
}
