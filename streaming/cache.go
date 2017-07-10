package streaming

import "log"


func CreateMarketCache(changeMessage MarketChangeMessage, marketChange MarketChange)(*MarketCache){
	cache := MarketCache{
		changeMessage.PublishTime,
		marketChange.MarketId,
		marketChange.TradedVolume,
		marketChange.MarketDefinition,
		make(map[int64]RunnerCache),
	}
	for _, runnerChange := range marketChange.RunnerChange {
		cache.Runners[runnerChange.SelectionId] = *CreateRunnerCache(runnerChange)
	}
	return &cache
}


func CreateRunnerCache(change RunnerChange)(*RunnerCache){
	log.Println("Creating new runner cache", change.SelectionId)

	// create traded data structure
	var traded Available
	for _, i := range change.Traded {
		traded.Prices = append(
			traded.Prices,
			PriceSize{i[0], i[1]},
		)
	}
	traded.Reverse = false

	// create availableToBack data structure
	var availableToBack Available
	for _, i := range change.AvailableToBack {
		availableToBack.Prices = append(
			availableToBack.Prices,
			PriceSize{i[0], i[1]},
		)
	}
	availableToBack.Reverse = true

	// create availableToLay data structure
	var availableToLay Available
	for _, i := range change.AvailableToLay {
		availableToLay.Prices = append(
			availableToLay.Prices,
			PriceSize{i[0], i[1]},
		)
	}
	availableToLay.Reverse = false

	// create startingPriceBack data structure
	var startingPriceBack Available
	for _, i := range change.StartingPriceBack {
		startingPriceBack.Prices = append(
			startingPriceBack.Prices,
			PriceSize{i[0], i[1]},
		)
	}
	startingPriceBack.Reverse = false

	// create startingPriceLay data structure
	var startingPriceLay Available
	for _, i := range change.StartingPriceLay {
		startingPriceLay.Prices = append(
			startingPriceLay.Prices,
			PriceSize{i[0], i[1]},
		)
	}
	startingPriceLay.Reverse = false

	// create bestAvailableToBack data structure
	var bestAvailableToBack AvailablePosition
	for _, i := range change.BestAvailableToBack {
		bestAvailableToBack.Prices = append(
			bestAvailableToBack.Prices,
			PositionPriceSize{i[0], i[1], i[2]},
		)
	}
	bestAvailableToBack.Reverse = false

	// create bestAvailableToLay data structure
	var bestAvailableToLay AvailablePosition
	for _, i := range change.BestAvailableToLay {
		bestAvailableToLay.Prices = append(
			bestAvailableToLay.Prices,
			PositionPriceSize{i[0], i[1],i[2]},
		)
	}
	bestAvailableToLay.Reverse = false

	// create bestDisplayAvailableToBack data structure
	var bestDisplayAvailableToBack AvailablePosition
	for _, i := range change.BestDisplayAvailableToBack {
		bestDisplayAvailableToBack.Prices = append(
			bestDisplayAvailableToBack.Prices,
			PositionPriceSize{i[0], i[1],i[2]},
		)
	}
	bestDisplayAvailableToBack.Reverse = false

	// create bestDisplayAvailableToLay data structure
	var bestDisplayAvailableToLay AvailablePosition
	for _, i := range change.BestDisplayAvailableToLay {
		bestDisplayAvailableToLay.Prices = append(
			bestDisplayAvailableToLay.Prices,
			PositionPriceSize{i[0], i[1],i[2]},
		)
	}
	bestDisplayAvailableToLay.Reverse = false

	cache := RunnerCache{
		change.SelectionId,
		&change.LastTradedPrice,
		&change.TradedVolume,
		&traded,
		&availableToBack,
		&availableToLay,
		&startingPriceBack,
		&startingPriceLay,
		&bestAvailableToBack,
		&bestAvailableToLay,
		&bestDisplayAvailableToBack,
		&bestDisplayAvailableToLay,
	}
	return &cache
}


type AvailableInterface interface {
	Clear()
	Sort()
	UpdatePrice(int, []float64)
	AppendPrice([]float64)
	RemovePrice(int)
	Update([][]float64)
}


type PriceSize struct {
	Price 		float64
	Size 		float64
}


type PositionPriceSize struct {
	Position 	float64
	Price 		float64
	Size 		float64
}


type AvailablePosition struct {
	Prices		[]PositionPriceSize
	Reverse		bool
}


func (available *AvailablePosition) Clear(){
	available.Prices = nil
}


func (available *AvailablePosition) Sort(){
	// todo
}


func (available *AvailablePosition) UpdatePrice(count int, update []float64){
	available.Prices[count] = PositionPriceSize{update[0], update[1], update[2]}
}


func (available *AvailablePosition) AppendPrice(update []float64){
	available.Prices = append(available.Prices, PositionPriceSize{update[0], update[1], update[2]})
}


func (available *AvailablePosition) RemovePrice(i int){
	s := available.Prices
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	available.Prices = s[:len(s)-1]
}


func (available *AvailablePosition) Update(updates [][]float64){
	for _, update := range updates {
		updated := false
		for count, trade := range available.Prices {
			if trade.Price == update[0] {
				if update[2] == 0 {
					available.RemovePrice(count)
					updated = true
					break
				} else {
					available.UpdatePrice(count, update)
					updated = true
					break
				}
			}
		}
		if updated == false && update[2] != 0 {
			available.AppendPrice(update)
		}
	}
}


type Available struct {
	Prices		[]PriceSize
	Reverse		bool
}


func (available *Available) Clear(){
	available.Prices = nil
}


func (available *Available) Sort(){
	// todo
}


func (available *Available) UpdatePrice(count int, update []float64){
	available.Prices[count] = PriceSize{update[0], update[1]}
}


func (available *Available) AppendPrice(update []float64){
	available.Prices = append(available.Prices, PriceSize{update[0], update[1]})
}


func (available *Available) RemovePrice(i int){
	s := available.Prices
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	available.Prices = s[:len(s)-1]
}


func (available *Available) Update(updates [][]float64){
	for _, update := range updates {
		updated := false
		for count, trade := range available.Prices {
			if trade.Price == update[0] {
				if update[1] == 0 {
					available.RemovePrice(count)
					updated = true
					break
				} else {
					available.UpdatePrice(count, update)
					updated = true
					break
				}
			}
		}
		if updated == false && update[1] != 0 {
			available.AppendPrice(update)
		}
	}
}


//func UpdateTwo(available AvailableInterface, updates [][]float64) {
//	for _, update := range updates {
//		updated := false
//		for count, trade := range available.Prices {
//			if trade.Price == update[0] {
//				if update[1] == 0 {
//					available.RemovePrice(count)
//					updated = true
//					break
//				} else {
//					available.UpdatePrice(count, update)
//					updated = true
//					break
//				}
//			}
//		}
//		if updated == false && update[1] != 0 {
//			available.AppendPrice(update)
//		}
//	}
//}


type RunnerCache struct {
	SelectionId 			int64
	LastTradedPrice 		*float64
	TradedVolume 			*float64
	//StartingPriceNear 		*float64
	//StartingPriceFar 		*float64
	Traded 				*Available
	AvailableToBack 		*Available
	AvailableToLay 			*Available
	StartingPriceBack 		*Available
	StartingPriceLay		*Available
	BestAvailableToBack		*AvailablePosition
	BestAvailableToLay		*AvailablePosition
	BestDisplayAvailableToBack	*AvailablePosition
	BestDisplayAvailableToLay	*AvailablePosition
}


func (cache *RunnerCache) UpdateCache(change RunnerChange) {
	if change.LastTradedPrice != 0 {
		*cache.LastTradedPrice = change.LastTradedPrice
	}
	if change.TradedVolume != 0 {
		*cache.TradedVolume = change.TradedVolume
	}
	//if change.StartingPriceNear != 0 {
	//	*cache.StartingPriceNear = change.StartingPriceNear
	//}
	//if change.StartingPriceFar != 0 {
	//	*cache.StartingPriceFar = change.StartingPriceFar
	//}
	if len(change.Traded) > 0 {
		cache.Traded.Update(change.Traded)
		//UpdateTwo(cache.Traded)
	}
	if len(change.AvailableToBack) > 0 {
		cache.AvailableToBack.Update(change.AvailableToBack)
	}
	if len(change.AvailableToLay) > 0 {
		cache.AvailableToLay.Update(change.AvailableToLay)
	}
	if len(change.StartingPriceBack) > 0 {
		cache.StartingPriceBack.Update(change.StartingPriceBack)
	}
	if len(change.StartingPriceLay) > 0 {
		cache.StartingPriceLay.Update(change.StartingPriceLay)
	}
	if len(change.BestAvailableToBack) > 0 {
		cache.BestAvailableToBack.Update(change.BestAvailableToBack)
		//UpdateTwo(cache.BestAvailableToBack)
	}
	if len(change.BestAvailableToLay) > 0 {
		cache.BestAvailableToLay.Update(change.BestAvailableToLay)
	}
	if len(change.BestDisplayAvailableToBack) > 0 {
		cache.BestDisplayAvailableToBack.Update(change.BestDisplayAvailableToBack)
	}
	if len(change.BestDisplayAvailableToLay) > 0 {
		cache.BestDisplayAvailableToLay.Update(change.BestDisplayAvailableToLay)
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
		for _, runnerChange := range marketChange.RunnerChange {
			if runnerCache, ok := cache.Runners[runnerChange.SelectionId]; ok {
				runnerCache.UpdateCache(runnerChange)
			} else {
				cache.Runners[runnerChange.SelectionId] = *CreateRunnerCache(runnerChange)
			}
		}
	}
	tem, _ := cache.Runners[7424945]
	log.Println(tem.SelectionId, *tem.LastTradedPrice, len(tem.Traded.Prices), *tem.TradedVolume,
	len(tem.AvailableToBack.Prices))
}
