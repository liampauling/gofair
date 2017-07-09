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


func (ms MarketStream) OnSubscribe(ChangeMessage MarketChangeMessage){
	log.Println(ChangeMessage)
}


func (ms MarketStream) OnResubscribe(ChangeMessage MarketChangeMessage){
	log.Println(ChangeMessage)
}


func (ms MarketStream) OnHeartbeat(ChangeMessage MarketChangeMessage){
	log.Println(ChangeMessage)
}


func (ms MarketStream) OnUpdate(ChangeMessage MarketChangeMessage){
	// todo update clk/initialClk
	for _, marketChange := range ChangeMessage.MarketChanges {
		log.Println(marketChange, ms.Cache)

		if marketCache, ok := ms.Cache[marketChange.MarketId]; ok {
			marketCache.UpdateCache(marketChange)
		} else {
			ms.Cache[marketChange.MarketId] = MarketCache{}
			log.Println("created cache", marketChange.MarketId)
		}
	}
	// todo output snap
}
