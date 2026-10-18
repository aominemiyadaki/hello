package errgroup

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

type Result struct {
	ID    string
	Value string
}

type FetchFunc func(ctx context.Context, id string) (Result, error)

func fetchAll(
	ctx context.Context,
	ids []string,
	limit int,
	fetch FetchFunc,
) ([]Result, error) {
	if limit <= 0 {
		return nil, errors.New("limit must be > 0")
	}

	// 1. Создаем группу и дочерний отменяемый контекст
	g, ctx := errgroup.WithContext(ctx)
	// 2. Устанавливаем лимит одновременно работающих горутин
	g.SetLimit(limit)

	results := make([]Result, len(ids))

	for i, id := range ids {
		// Быстрый выход: если кто-то уже упал, не планируем новые задачи
		if ctx.Err() != nil {
			break
		}

		// Захватываем переменные цикла для горутины
		i, id := i, id

		// 3. g.Go сам блокируется, если лимит исчерпан, пока не освободится слот
		g.Go(func() error {
			res, err := fetch(ctx, id)
			if err != nil {
				// При первой ошибке errgroup сама вызовет cancel() для ctx
				return fmt.Errorf("fetch %s: %w", id, err)
			}

			// Безопасно пишем в слайс: каждый индекс уникален
			results[i] = res
			return nil
		})
	}

	// 4. Ждем завершения всех задач и получаем первую ошибку (если была)
	if err := g.Wait(); err != nil {
		return nil, err
	}

	return results, nil
}

func main() {
	limit := 2
	ids := []string{"1", "2", "3", "4", "5"}

	results, err := fetchAll(context.Background(), ids, limit, func(ctx context.Context, id string) (Result, error) {
		if id == "3" {
			return Result{}, fmt.Errorf("bad item")
		}

		select {
		case <-time.After(300 * time.Millisecond):
			return Result{ID: id, Value: "value_" + id}, nil
		case <-ctx.Done():
			return Result{}, ctx.Err()
		}
	})

	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Printf("Успешно: %+v\n", results)
}
