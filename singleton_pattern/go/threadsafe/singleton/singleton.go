package singleton

import "sync"

var (
	uniqueInstance *Singleton

	mu sync.Mutex
)

type Singleton struct {
}

// GetDescription returns a description of Singleton
func (s *Singleton) GetDescription() string {
	return "I'm a singleton using sync.Mutex!"
}

// GetInstance returns a singleton instance of Singleton
func GetInstance() *Singleton {
	mu.Lock()
	defer mu.Unlock()

	if uniqueInstance == nil {
		uniqueInstance = &Singleton{}
	}
	return uniqueInstance
}
