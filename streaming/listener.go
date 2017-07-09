package streaming


type Listener struct {
	MarketStream	*MarketStream
	OrderStream	Stream
	UniqueId	int64
}


func (l *Listener) AddMarketStream() {
	l.MarketStream = new(MarketStream)
	l.MarketStream.Cache = make(map[string]MarketCache)
}


func (l *Listener) AddOrderStream() {
	// todo
}


func (l *Listener) OnData(ChangeMessage MarketChangeMessage) {
	//todo check unique id
	//todo error handler

	switch *ChangeMessage.Operation {
	case "connection":
		l.onConnection(ChangeMessage)
	case "status":
		l.onStatus(ChangeMessage)
	case "mcm":
		l.onChangeMessage(l.MarketStream, ChangeMessage)
	case "ocm":
		l.onChangeMessage(l.OrderStream, ChangeMessage)
	}
}


func (l *Listener) onConnection(ChangeMessage MarketChangeMessage) {
	// todo
}


func (l *Listener) onStatus(ChangeMessage MarketChangeMessage) {
	// todo
}


func (l *Listener) onChangeMessage(Stream Stream, ChangeMessage MarketChangeMessage) {
	switch ChangeMessage.ChangeType {
	case "SUB_IMAGE":
		Stream.OnSubscribe(ChangeMessage)
	case "RESUB_DELTA":
		Stream.OnResubscribe(ChangeMessage)
	case "HEARTBEAT":
		Stream.OnHeartbeat(ChangeMessage)
	default:
		Stream.OnUpdate(ChangeMessage)
	}
}
