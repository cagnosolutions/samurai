package samurai

import "sync"

type DataStore struct {
	stores map[string]*SafeMap
	sync.RWMutex
}

func DataStoreInstance() *DataStore {
	return &DataStore{
		stores: make(map[string]*SafeMap, 0),
	}
}

func (ds *DataStore) AddStore(name string) bool {
	ds.Lock()
	if _, ok := ds.stores[name]; !ok {
		ds.stores[name] = SafeMapInstance(16)
		ds.Unlock()
		return true
	}
	ds.Unlock()
	return false
}

func (ds *DataStore) DelStore(name string) bool {
	ds.Lock()
	if _, ok := ds.stores[name]; ok {
		delete(ds.stores, name)
		ds.Unlock()
		return true
	}
	ds.Unlock()
	return false
}

func (ds *DataStore) UseStore(name string) (*SafeMap, bool) {
	ds.RLock()
	store, ok := ds.stores[name]
	ds.RUnlock()
	return store, ok
}
