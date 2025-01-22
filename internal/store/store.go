package store

import "sync"

var DataStore *sync.Map
var DataStoreStat *sync.Map

func InitStore() {
	DataStore = new(sync.Map)
	DataStoreStat = new(sync.Map)
}
