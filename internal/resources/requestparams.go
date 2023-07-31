package resources

type requestParams struct {
	requestID string
}

var instance *requestParams

func getRequestParamsInstance() *requestParams {
	if instance == nil {
		instance = &requestParams{}
	}

	return instance
}

func SetRequestID(reqID string) {
	res := getRequestParamsInstance()
	res.requestID = reqID
}

func RequestID() string {
	res := getRequestParamsInstance()
	return res.requestID
}
