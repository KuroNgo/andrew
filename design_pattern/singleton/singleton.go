package singleton

// Singleton will provide you with a single instance as an object
// and guarantee that there are no duplicated
// You'll use the Singleton pattern in many different situations. For example:
// When you want to use the same connection to a database to make every query
type Singleton interface {
	AddOne() int
}
type singleton struct {
	count int
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

var instance *singleton

// GetInstance At the first call to use the instance,
// it is created and then reused between
// all the parts in the application that need to use that particular behavior
// When you open a Secure Shell SSH connection to a server to do a few tasks,
// and don't want to reopen the connection for each task
// If you need to limit the access to some variable or space, you use a Singleton as the door
// to this variable (we'll see in the following chapters that this is more achievable in Go using channels anyway)
// If you need to limit the number of calls to some places,
// yuu create a Singleton instance to make the calls in the accepted window
func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}
