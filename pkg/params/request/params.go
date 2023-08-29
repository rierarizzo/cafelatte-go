package request

type requestParams struct {
	requestId string
}

var instance *requestParams

func getRequestParamsInstance() *requestParams {
	if instance == nil {
		instance = &requestParams{}
	}

	return instance
}

func SetRequestId(reqId string) {
	res := getRequestParamsInstance()
	res.requestId = reqId
}

func Id() string {
	res := getRequestParamsInstance()
	return res.requestId
}
