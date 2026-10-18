package semaforerr

/*
 стопим работу если в одной из горутин произошла ошибка

*/

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
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

	// 1. Создаем отменяемый дочерний контекст
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	results := make([]Result, len(ids))
	sem := make(chan struct{}, limit) // Семафор для ограничения параллельности
	var wg sync.WaitGroup

	var once sync.Once
	var firstErr error

	for i, id := range ids {
		// 2. Быстрая проверка: если кто-то уже упал, не запускаем новые задачи
		if ctx.Err() != nil {
			break
		}

		wg.Add(1)
		go func(idx int, itemID string) {
			defer wg.Done()

			// 3. Захват слота в семафоре с поддержкой отмены по контексту
			select {
			case sem <- struct{}{}:
			case <-ctx.Done():
				return
			}
			defer func() { <-sem }()

			// 4. Повторная проверка перед тяжелой работой
			if ctx.Err() != nil {
				return
			}

			// 5. Вызываем fetch с ДОЧЕРНИМ контекстом
			res, err := fetch(ctx, itemID)
			if err != nil {
				// Ровно один раз сохраняем первую ошибку и отменяем контекст для всех
				once.Do(func() {
					firstErr = err
					cancel()
				})
				return
			}

			// 6. Записываем результат только при успехе
			results[idx] = res
		}(i, id)
	}

	// 7. Ждем завершения всех запущенных горутин
	wg.Wait()

	// 8. Если была хоть одна ошибка — возвращаем nil и ошибку
	if firstErr != nil {
		return nil, firstErr
	}

	return results, nil
}

func main() {
	limit := 2
	ids := []string{"1", "2", "3", "4", "5"}

	results, err := fetchAll(context.Background(), ids, limit, func(ctx context.Context, id string) (Result, error) {
		if id == "3" {
			return Result{}, fmt.Errorf("error in item %s", id)
		}

		select {
		case <-time.After(200 * time.Millisecond):
			return Result{ID: id, Value: "val_" + id}, nil
		case <-ctx.Done():
			return Result{}, ctx.Err()
		}
	})

	if err != nil {
		fmt.Printf("Обработка завершилась с ошибкой: %v\n", err)
		return
	}

	fmt.Printf("Успех: %+v\n", results)
}
