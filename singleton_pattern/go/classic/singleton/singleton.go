package singleton

var uniqueInstance *Singleton

type Singleton struct {
}

// GetDescription returns a description of Singleton
func (s *Singleton) GetDescription() string {
	return "I'm a classic Singleton!"
}

// GetInstance returns a singleton instance of Singleton
func GetInstance() *Singleton {
	if uniqueInstance == nil {
		uniqueInstance = &Singleton{}
	}
	return uniqueInstance
}
