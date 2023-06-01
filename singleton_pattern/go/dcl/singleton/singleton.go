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
	return "I'm a classic Singleton!"
}

// GetInstance returns a singleton instance of Singleton
func GetInstance() *Singleton {
	// double-checked locking
	if uniqueInstance == nil {
		mu.Lock()
		defer mu.Unlock()
		if uniqueInstance == nil {
			uniqueInstance = &Singleton{}
		}
	}
	return uniqueInstance
}
