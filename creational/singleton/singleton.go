package singleton

import (
	"fmt"
	"sync"
)

// wiki: https://en.wikipedia.org/wiki/Singleton_pattern

/*
 * Singleton is a creational design pattern that ensure a class has only one instance,
 * and provide a global point of access to it.
 */

type single struct{}

var singleInstance *single
var once sync.Once

func GetInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				// creating single instance now...
				fmt.Println("Creating single instance now.")
				singleInstance = &single{}
			})
	} else {
		// single instance already created
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
