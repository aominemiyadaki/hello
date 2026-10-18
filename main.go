package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}

	var nums []*int

	for i := 0; i < 10; i++ {
		nums = append(nums, &i)
	}

	wg.Add(len(nums))
	for _, v := range nums {

		go func() {
			defer wg.Done()
			fmt.Println(*v)
		}()
	}

	wg.Wait()
}
