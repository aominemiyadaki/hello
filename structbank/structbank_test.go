package structbank

import "testing"

func TestBankAccount(t *testing.T) {

	account := &BankAccount{}

	if account.Balance() != 0 {
		t.Errorf("на старте ожидали 0, по получили %.2f", account.Balance())
	}

	account.Deposit(50)

	expected := 50.0
	actual := account.Balance()

	if actual != float64(expected) {

		t.Errorf("После добавления ожидали %.2f, но получили %.2f", expected, actual)
	}

	err := account.Withdraw(50)

	if err != nil {
		t.Errorf("Не ожидали ошибку при списании, но получили: %v", err)
	}

	expectedWithDraw := 0.0
	actualWithDraw := account.Balance()

	if expectedWithDraw != actualWithDraw {
		t.Errorf("ожидаемый баланс не совпал с балансом списания")
	}

	// Пытаемся снять 100 рублей, хотя на балансе сейчас 0
	err = account.Withdraw(100)

	// Здесь мы НАОБОРОТ ждем, что ошибка появится (то есть err не должен быть nil)
	if err == nil {
		t.Errorf("Ожидали ошибку нехватки средств, но метод вернул nil")
	}
}
