package singleton

import (
	"github.com/sirupsen/logrus"
	"sync"
)

var lock = &sync.Mutex{}

type sharedResources struct {
	requestID string
}

var singleInstance *sharedResources

func getSharedResourcesInstance() *sharedResources {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			logrus.Debug("Creating sharedResources instance now.")
			singleInstance = &sharedResources{}
		}
	}

	return singleInstance
}

func SetRequestID(reqID string) {
	res := getSharedResourcesInstance()
	res.requestID = reqID
}

func RequestID() string {
	res := getSharedResourcesInstance()
	return res.requestID
}
