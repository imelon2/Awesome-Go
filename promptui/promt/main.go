package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
)

func main() {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	prompt1 := promptui.Prompt{
		Label:    "is?",
		Validate: validate,
	}

	prompt2 := promptui.Prompt{
		Label:    "is???",
		Validate: validate,
	}
	result1, err := prompt1.Run()
	result2, err := prompt2.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result1)
	fmt.Printf("You choose %q\n", result2)
}
