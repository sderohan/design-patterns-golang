package main

import "sync"

var (
	lock                = &sync.Mutex{}
	instance *Singleton = nil
)

type Singleton struct {
}

func (state *Singleton) getInstance() *Singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Singleton{}
		}
	}
	return instance
}

func main() {

}
