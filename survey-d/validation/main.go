package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

// the questions to ask
var validationQs = []*survey.Question{
	{
		Name:   "valid",
		Prompt: &survey.Input{Message: "Enter 'foo':", Default: "not foo"},
		Validate: func(val interface{}) error {
			// if the input matches the expectation
			if str := val.(string); str != "foo" {
				return fmt.Errorf("You entered %s, not 'foo'.", str)
			}
			// nothing was wrong
			return nil
		},
	},
}

func main() {
	// the place to hold the answers
	answers := struct {
		Name  string
		Valid string
	}{}
	err := survey.Ask(validationQs, &answers)

	if err != nil {
		fmt.Println("\n", err.Error())
	}
}
