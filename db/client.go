package db

import (
	"log"
)

// DataStore is an interface for query ops
type DataStore interface {
}

// Option holds configuration for data store clients
type Option struct {
	TestMode bool
}

// DataStoreFactory holds configuration for data store
type DataStoreFactory func(conf Option) (DataStore, error)

var datastoreFactories = make(map[string]DataStoreFactory)

// Register saves data store into a data store factory
func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Fatalf("Datastore factory %s does not exist.", name)
		return
	}
	if _, ok := datastoreFactories[name]; ok {
		log.Fatalf("Datastore factory %s already registered. Ignoring.", name)
		return
	}
	datastoreFactories[name] = factory
}
