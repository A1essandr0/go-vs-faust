package main

type RawEvent struct {
	Source string 		`json:"source"`
	EventName string 	`json:"event_name"`
	EventStatus string 	`json:"event_status"`
	Created string 		`json:"created"`
	Payout string 		`json:"payout"`
}

type ProcessedEvent struct {
	// these are base fields from raw event
	Source string 		`json:"source"`
	EventName string 	`json:"event_name"`
	EventStatus string 	`json:"event_status"`
	Created string 		`json:"created"`
	Payout string 		`json:"payout"`

    // these are added by processors
    AppId string 		`json:"app_id"`

    EventValue string 	`json:"event_value"`

    ProcessTime string 	`json:"process_time"`
    ProcessDate string 	`json:"process_date"`
    OsName string 		`json:"os_name"`

    MytrackerStatus string 		`json:"my_tracker_status"`
    MytrackerResponse string 	`json:"my_tracker_rsp"`
    AppmetricaStatus string 	`json:"appmetrica_status"`
    AppmetricaResponse string 	`json:"appmetrica_rsp"`
    AppsflyerStatus string 		`json:"appsflyer_status"`
    AppsflyerResponse string 	`json:"appsflyer_rsp"`

    Error string 		`json:"error"`
}