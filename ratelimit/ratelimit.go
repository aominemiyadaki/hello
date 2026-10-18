package ratelimit

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://api.site.com/page1",
		"https://api.site.com/page2",
		"https://api.site.com/page3",
		"https://api.site.com/page4",
		"https://api.site.com/page5",
		"https://api.site.com/page6",
		"https://api.site.com/page7",
		"https://api.site.com/page8",
		"https://api.site.com/page9",
		"https://api.site.com/page10",
	}

	sem := make(chan struct{}, 3)

	var wg = sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)
		go fetch(url, sem, &wg)
	}

	wg.Wait()

}

func fetch(url string, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	sem <- struct{}{}

	defer func() {

		<-sem
	}()

	fmt.Println("начали скачивать", url)

	time.Sleep(1 * time.Second)

	fmt.Println("закончили скачивать", url)

}
