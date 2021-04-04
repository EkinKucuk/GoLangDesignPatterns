package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type SingletonObject struct {
	id    string
	value int
}

var singletonInstance *SingletonObject

func getInstance() *SingletonObject {

	if singletonInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singletonInstance == nil {
			fmt.Println("Creating new singleton instance")
			singletonInstance = &SingletonObject{id: "1", value: 1}
		} else {
			fmt.Println("Singleton instance already created")
		}
	} else {
		fmt.Println("Singleton instance already created")
	}
	return singletonInstance
}
