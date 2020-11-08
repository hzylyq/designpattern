package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}
var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Printf("Creting Single Instance Now")
			singleInstance = &single{}
		} else {
			fmt.Printf("Single Instance already created-1")
		}
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return singleInstance
}

func getInstanceOnce() *single {
	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("Creting Single Instance Now")
			singleInstance = &single{}
		})
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return singleInstance
}
