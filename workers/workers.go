package workers

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}

	jobs := make(chan int, 100)
	result := make(chan int, 100)

	for range 100 {
		jobs <- rand.Intn(10)
	}

	close(jobs)

	for i := range 3 {
		wg.Add(1)
		go worker(i, jobs, result, &wg)
	}

	wg.Wait()

	close(result)

	for i := range result {
		fmt.Println(i)
	}
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {

		fmt.Printf("Воркер %d взял задачу %d\n", id, job)
		time.Sleep(5 * time.Millisecond)
		results <- job * 2

	}
}
