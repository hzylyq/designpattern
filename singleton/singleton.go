package singleton

import "sync"

type singleton struct{}

var lock = &sync.Mutex{}

var singleInstance *singleton

func getInstance() *singleton {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &singleton{}
		}
	}
	return singleInstance
}
