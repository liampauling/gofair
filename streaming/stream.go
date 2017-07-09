package streaming

import (
	"log"
)

type Stream interface {
	OnSubscribe(ChangeMessage MarketChangeMessage)
	OnResubscribe(ChangeMessage MarketChangeMessage)
	OnHeartbeat(ChangeMessage MarketChangeMessage)
	OnUpdate(ChangeMessage MarketChangeMessage)
}


type MarketStream struct {
	Cache	map[string]MarketCache
}


func (ms *MarketStream) OnSubscribe(changeMessage MarketChangeMessage){
	log.Println(changeMessage)
}


func (ms *MarketStream) OnResubscribe(changeMessage MarketChangeMessage){
	log.Println(changeMessage)
}


func (ms *MarketStream) OnHeartbeat(changeMessage MarketChangeMessage){
	log.Println(changeMessage)
}


func (ms *MarketStream) OnUpdate(changeMessage MarketChangeMessage){
	// todo update clk/initialClk

	for _, marketChange := range changeMessage.MarketChanges {

		if marketCache, ok := ms.Cache[marketChange.MarketId]; ok {
			marketCache.UpdateCache(changeMessage, marketChange)
		} else {
			ms.Cache[marketChange.MarketId] = CreateMarketCache(changeMessage, marketChange)
			log.Println("Created new market cache", marketChange.MarketId)
		}
	}
	// todo output snap
}
