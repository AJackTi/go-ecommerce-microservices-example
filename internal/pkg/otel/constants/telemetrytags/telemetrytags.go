package telemetrytags

type app struct {
	Request           string
	RequestName       string
	RequestType       string
	RequestResultName string
	RequestResult     string
	CommandName       string
	Command           string
	CommandResultName string
	CommandResult     string
	QueryName         string
	Query             string
	QueryResultName   string
	QueryResult       string
	EventName         string
	Event             string
	EventResultName   string
	EventResult       string
}

var App = app{
	Request:           "app.request",
	RequestName:       "app.request_name",
	RequestType:       "app.request_type",
	RequestResultName: "app.request_result_name",
	RequestResult:     "app.request_result",
	CommandName:       "app.command_name",
	Command:           "app.command",
	CommandResultName: "app.command_result_name",
	CommandResult:     "app.command_result",
	QueryName:         "app.query_name",
	Query:             "app.query",
	QueryResultName:   "app.query_result_name",
	QueryResult:       "app.query_result",
	EventName:         "app.event_name",
	Event:             "app.event",
	EventResultName:   "app.event_result_name",
	EventResult:       "app.event_result",
}

type exceptions struct {
	EventName  string
	Type       string
	Message    string
	Stacktrace string
}

var Exceptions = exceptions{
	EventName:  "exception",
	Type:       "exception.type",
	Message:    "exception.message",
	Stacktrace: "exception.stacktrace",
}
