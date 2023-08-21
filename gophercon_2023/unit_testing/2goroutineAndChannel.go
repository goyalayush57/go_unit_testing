package unit_testing

import (
	"fmt"
	"sync"
)

// Task represents a task to be processed.
type Task struct {
	ID   int
	Data string
}

func worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		fmt.Printf("Worker %d processing task %d: %s\n", id, task.ID, task.Data)
		// Simulate processing time
		//time.Sleep(time.Second)
	}
}

type process struct {
	numWorkers int
	numTasks   int
	wg         *sync.WaitGroup
}

func (p *process) DoStuffConcurrently() {
	//const numWorkers = 3
	//const numTasks = 10

	tasks := make(chan Task, p.numTasks)

	// Start worker goroutines
	for i := 1; i <= p.numWorkers; i++ {
		p.wg.Add(1)
		go worker(i, tasks, p.wg)
	}

	// Send tasks to the channel
	for i := 1; i <= p.numTasks; i++ {
		tasks <- Task{ID: i, Data: fmt.Sprintf("Task %d", i)}
	}

	// Close the tasks channel to signal workers to stop
	close(tasks)

	// Wait for all workers to finish
	p.wg.Wait()
	fmt.Println("All tasks processed.")
}
