package main

import (
	"fmt"
)

type Operation interface {
	get_result (int, int) (result int, errorMsg string)
}

type Ope_add struct {
}

type Ope_sub struct {
}

type Ope_mul struct {
}

type Ope_div struct {
}

func (add *Ope_add) get_result(a int, b int) (result int, errorMsg string) {
	return a + b, ""
}

func (sub *Ope_sub) get_result(a int, b int) (result int, errorMsg string) {
	return a - b, ""
}

func (mul *Ope_mul) get_result(a int, b int) (result int, errorMsg string) {
	return a * b, ""
}

func (div *Ope_div) get_result(a int, b int) (result int, errorMsg string) {
	if b==0 {
		errorMsg = "Cannot proceed, the divider is zero."
		return 
	} else {
		return a / b, ""
	}
}


func OperationFactory(oper string) Operation {
	switch oper {
	case "+":
		return &Ope_add{}
	case "-":
		return &Ope_sub{}
	case "*":
		return &Ope_mul{}
	case "/":
		return &Ope_div{}
	default:
		return nil
	}
	
}

func main() {
	Operator_add := OperationFactory("+")
	if result, errorMsg := Operator_add.get_result(8,2); errorMsg == "" {
		fmt.Println("add result is ", result)
	} else {
		fmt.Println("errorMsg is: ", errorMsg)
	}
	Operator_sub := OperationFactory("-")
	if result, errorMsg := Operator_sub.get_result(8,2); errorMsg == "" {
		fmt.Println("sub result is ", result)
	} else {
		fmt.Println("errorMsg is: ", errorMsg)
	}
	Operator_mul := OperationFactory("*")
	if result, errorMsg := Operator_mul.get_result(8,2); errorMsg == "" {
		fmt.Println("mul result is ", result)
	} else {
		fmt.Println("errorMsg is: ", errorMsg)
	}
	Operator_div := OperationFactory("/")
	if result, errorMsg := Operator_div.get_result(8,0); errorMsg == "" {
		fmt.Println("div result is ", result)
	} else {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}