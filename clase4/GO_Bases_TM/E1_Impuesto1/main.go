package main

import "fmt"

type MyError struct {
	msg string
}

func main() {
	salary := 150000

	err := MyErrorSalaryTest(salary)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Debe pagar impuesto\n")
}

func (err MyError) Error() string {
	return fmt.Sprintf("%v\n", err.msg)
}

func MyErrorSalaryTest(salary int) error {
	if salary < 150000 {
		return &MyError{
			msg: "error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return nil
}
