package models

import (
	"sync"
)

// Bus struct with a receiver channel called Bus
type Bus struct {
	Bus chan interface{}
}

// The single instance of the Bus
var instance *Bus
var once sync.Once

// GetBus returns the singleton instance, initializing it if necessary
func GetBus() *Bus {
	once.Do(func() {
		instance = &Bus{
			Bus: make(chan interface{}),
		}
	})
	return instance
}
