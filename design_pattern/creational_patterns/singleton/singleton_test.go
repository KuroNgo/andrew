package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil {
		// Test of acceptance criteria 1 failed
		t.Error("expected pointer to Singleton after calling GetInstance(), not nil")
	}
	expectedCounter := counter1
	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", currentCount)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		// Test 2 fail
		t.Error("Expected same instance in counter2 but it got a difference instance")
	}
	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("")
	}

}
