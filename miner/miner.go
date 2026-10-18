package miner

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Miner(
	ctx context.Context,
	pointer chan<- int,
	id int,
	wg *sync.WaitGroup,
) {

	defer wg.Done()

	for {
		count := rand.Intn(20)
		// если хотим чтобы текущая итерация по добыче прекратилась
		// select {
		// case <-ctx.Done():
		// 	fmt.Println("Шахтер завершил работу, id:", id)
		// 	return
		// case <-time.After(1 * time.Second):
		// 	fmt.Println("Шахетр Закончил добывать, id:", id)
		// }

		// select {
		// case <-ctx.Done():
		// 	fmt.Println("Шахтер завершил работу, id:", id)
		// 	return
		// case pointer <- count:

		// }

		// если хотим чтобы текущая итерация по добыче продолжилась и завершилась

		select {
		case <-ctx.Done():
			fmt.Println("Шахтер завершил работу, id:", id)
			return
		default:
			fmt.Println("Шахтер Начал добывать, id:", id)
			rtime := rand.Intn(5)
			time.Sleep(time.Duration(rtime) * time.Second)
			fmt.Println("Шахетр Закончил добывать, id:", id)

			pointer <- count
		}

	}
}

func MinerPool(ctx context.Context, minerCount int) <-chan int {

	pointer := make(chan int)
	wg := &sync.WaitGroup{}

	for id := range minerCount {
		wg.Add(1)
		go Miner(ctx, pointer, id, wg)
	}

	go func() {
		wg.Wait()
		close(pointer)
	}()

	return pointer
}

// main

// var coal atomic.Int64
// var mail []string

// mtx := sync.Mutex{}

// minerContext, minerCancel := context.WithCancel(context.Background())
// postmanContext, postmanCancel := context.WithCancel(context.Background())

// go func() {
// 	time.Sleep(3 * time.Second)
// 	minerCancel()
// 	fmt.Println("РАБОЧИЙ ДЕНЬ ШАХТЕРОВ ОКОНЧЕН!")
// }()

// go func() {
// 	time.Sleep(6 * time.Second)
// 	postmanCancel()
// 	fmt.Println("РАБОЧИЙ ДЕНЬ ШАХТЕРОВ ОКОНЧЕН!")
// }()

// minerPointer := miner.MinerPool(minerContext, 1000)
// mailPointer := postman.PostmanPool(postmanContext, 1000)

// wg := &sync.WaitGroup{}

// wg.Add(1)
// go func() {
// 	defer wg.Done()
// 	for v := range minerPointer {
// 		coal.Add(int64(v))
// 	}
// }()

// wg.Add(1)
// go func() {
// 	defer wg.Done()
// 	for v := range mailPointer {
// 		mtx.Lock()
// 		mail = append(mail, v)
// 		mtx.Unlock()
// 	}
// }()

// wg.Wait()

// fmt.Println("Добытого угля", coal.Load())

// mtx.Lock()
// fmt.Println("Доставленных писем", len(mail))
// mtx.Unlock()
