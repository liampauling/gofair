package streaming

import "log"


func CreateMarketCache(changeMessage MarketChangeMessage, marketChange MarketChange)(MarketCache){
	var cache MarketCache
	cache.PublishTime = changeMessage.PublishTime
	cache.MarketId = marketChange.MarketId
	cache.TotalMatched = marketChange.TradedVolume
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
	return cache
}


type RunnerCache struct {
	SelectionId 	int64
	LastTradedPrice *float64
}


func (cache *RunnerCache) UpdateCache(change RunnerChange) {
	if change.LastTradedPrice != nil {
		*cache.LastTradedPrice = *change.LastTradedPrice
	}
}


type MarketCache struct {
	PublishTime		int64
	MarketId		string
	TotalMatched		*float64
	MarketDefinition	*MarketDefinition
	Runners 		map[int64]RunnerCache
}


func (cache *MarketCache) UpdateCache(changeMessage MarketChangeMessage, marketChange MarketChange) {
	cache.PublishTime = changeMessage.PublishTime

	if marketChange.MarketDefinition != nil {
		cache.MarketDefinition = marketChange.MarketDefinition
	}
	if marketChange.TradedVolume != nil {
		cache.TotalMatched = marketChange.TradedVolume
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
	tem, _ := cache.Runners[13219181]
	log.Println(tem, tem.SelectionId, *tem.LastTradedPrice)
}
