# gofair

Lightweight golang wrapper for [Betfair API-NG](http://docs.developer.betfair.com/docs/display/1smk3cen4v3lu3yomq5qye0ni)


## use

```golang
config := &gofair.Config{
		"username",
		"password",
		"appKey",
		"/certs/client-2048.crt",
		"/certs/client-2048.key",
		"",
}

trading, err := gofair.NewClient(config)
if err != nil {
    panic(err)
}


fmt.Println(trading.Login())
fmt.Println(trading.KeepAlive())
fmt.Println(trading.Logout())

filter := new(gofair.MarketFilter)
event_types, err := trading.Betting.ListEventTypes(filter)

fmt.Println(event_types)
```
