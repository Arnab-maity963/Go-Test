package main

import (
	"fmt"
	slp "github.com/Rajdeep-Nemo/sleepycat"
)

func main() {
	age, err := slp.Int(
		slp.Prompt("Enter age: "),
		slp.MaxAttempt(3),
	)
	if err != nil {
		fmt.Println("Invalid....");
	} else {
		fmt.Println("Your age is",age);
	}
}
