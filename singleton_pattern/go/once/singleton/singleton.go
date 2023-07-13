package singleton

import "sync"

var (
	uniqueInstance *Singleton
	once           sync.Once
)

type Singleton struct {
}

// GetDescription returns a description of Singleton
func (s *Singleton) GetDescription() string {
	return "I'm a singleton using sync.Once!"
}

// GetInstance returns a singleton instance of Singleton
func GetInstance() *Singleton {
	once.Do(func() {
		uniqueInstance = &Singleton{}
	})
	return uniqueInstance
}
