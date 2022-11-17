package main

func GetZQuery() zQuery {
	var instance zQuery
	instance.Response = 200
	instance.Body.Result = "done"
	instance.Body.RequestedReturnCode = 200
	return instance
}

type body struct {
	Result              string `json:"result"`
	RequestedReturnCode int    `json:"requestedreturncode"`
}

type zQuery struct {
	Response int
	Body     body
}

type zErr struct {
	Message string `json:"message"`
}
