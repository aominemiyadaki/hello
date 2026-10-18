package wordcounter

import (
	"reflect"
	"testing"
)

func TestWordCounter(t *testing.T) {
	// 1. Описываем структуру нашего тест-кейса
	type testCase struct {
		name     string         // имя теста для удобства чтения логов
		input    string         // что подаем на вход функции
		expected map[string]int // что ожидаем получить на выходе
	}

	// 2. Создаем «таблицу» (срез) с разными сценариями
	tests := []testCase{
		{
			name:  "Простая строка",
			input: "Привет я учу язык Go",
			expected: map[string]int{
				"привет": 1, "я": 1, "учу": 1, "язык": 1, "go": 1,
			},
		},
		{
			name:  "Разный регистр и повторения",
			input: "Go go GO Язык go",
			expected: map[string]int{
				"go": 4, "язык": 1,
			},
		},
		{
			name:  "Очистка от знаков препинания",
			input: "Привет, привет! Ты учишь Go?",
			expected: map[string]int{
				"привет": 2, "ты": 1, "учишь": 1, "go": 1,
			},
		},
		{
			name:     "Пустая строка",
			input:    "   ",
			expected: map[string]int{}, // ждем пустую мапу
		},
	}

	// 3. Запускаем цикл по всем тест-кейсам
	for _, tc := range tests {
		// t.Run запускает каждый сценарий как отдельный подтест
		t.Run(tc.name, func(t *testing.T) {
			actual := WordCounter(tc.input)

			// В Go мапы нельзя сравнивать через ==
			// Для глубокого сравнения структур используют reflect.DeepEqual
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("WordCounter(%q)\nПолучили: %v\nОжидали:  %v", tc.input, actual, tc.expected)
			}
		})
	}
}
