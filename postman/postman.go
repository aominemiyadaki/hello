package postman

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Postman(
	wg *sync.WaitGroup,
	pointer chan<- string,
	msg string,
	id int,
	ctx context.Context,
) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Почтальон завершил работу, id:", id)
			return
		default:
			fmt.Println("Почтальон  начал доставлять, id:", id)
			rtime := rand.Intn(5)
			time.Sleep(time.Duration(rtime) * time.Second)
			fmt.Println("Почтальон доставил до почты, id:", id)

			pointer <- msg

			fmt.Println("Я почтальон передал письмо, id:", id, "Содержание письма: ", msg)
		}
	}
}

func PostmanPool(ctx context.Context, count int) <-chan string {

	pointer := make(chan string)
	wg := &sync.WaitGroup{}

	for id := range count {
		wg.Add(1)

		msg := postmanToMail(id)
		go Postman(wg, pointer, msg, id, ctx)

	}

	go func() {
		wg.Wait()
		close(pointer)
	}()

	return pointer
}

func postmanToMail(id int) string {
	ptm := map[int]string{
		0: "zero",
		1: "aboba",
		2: "labuba",
		3: "amoga",
	}

	mail, ok := ptm[id]

	if !ok {
		fmt.Println("писем нету - подметаем пол")
		return "Нет письма, газета, реклама"
	}

	return mail

}
