package singleton

import (
	"sync"
)

// wiki: https://en.wikipedia.org/wiki/Singleton_pattern
//
// Restricts instantiation of a type to one object

type singleton map[string]string

var (
	once     sync.Once
	instance singleton
)

func New() singleton {
	if instance == nil {
		once.Do(
			func() {
				instance = make(singleton)
			})
	}
	return instance
}
