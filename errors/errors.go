package errors

import (
	"fmt"
)

// "errors"
// "fmt"
// customerr "hello/errors"

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("поле %s: %s", e.Field, e.Message)
}

func ValidateAge(age int) error {

	if age < 0 {
		return &ValidationError{Field: "age", Message: "не может быт отрицательнным"}
	}

	if age > 150 {
		return &ValidationError{Field: "age", Message: "нереалистичное значение"}
	}
	return nil
}

func ProcessUser(age int) error {
	if err := ValidateAge(age); err != nil {
		return fmt.Errorf("processUser: %w", err)
	}

	return nil
}

// func main() {

// 	err := customerr.ProcessUser(-5)

// 	if err != nil {
// 		fmt.Println(err)

// 		var valErr *customerr.ValidationError
// 		if errors.As(err, &valErr) {
// 			fmt.Printf("Проблемное поле: %s\n", valErr.Field)
// 		}
// 	}
// }
