package unit_testing

import (
	"sync"
	"testing"
)

/*
Concurrency testing of lock-based code is hard, to the extent that provable-correct solutions are difficult to come by.
Ad-hoc manual testing via print statements is not ideal.
The Four Horsemen of the Test-Driven Apocalypse

https://www.bigbeeconsultants.uk/post/2008/four-horsemen/

1. Race Conditions
2. Deadlock
3. Livelock
4. Starvation
*/
func TestCounter(t *testing.T) {
	tests := []struct {
		name     string
		actions  []func(c *Counter)
		expected int
	}{
		{
			name: "Single Increment",
			actions: []func(c *Counter){
				func(c *Counter) { c.Increment() },
			},
			expected: 1,
		},
		{
			name: "Multiple Increments",
			actions: []func(c *Counter){
				func(c *Counter) { c.Increment() },
				func(c *Counter) { c.Increment() },
				func(c *Counter) { c.Increment() },
			},
			expected: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel() // Mark the subtest for parallel execution

			counter := NewCounter()

			// Perform the actions in parallel
			var wg sync.WaitGroup
			for _, action := range test.actions {
				wg.Add(1)
				go func(action func(c *Counter)) {
					defer wg.Done()
					action(counter)
				}(action)
			}
			wg.Wait()

			// Check the final value
			finalValue := counter.Value()
			if finalValue != test.expected {
				t.Errorf("Expected %d, but got %d", test.expected, finalValue)
			}
		})
	}
}
