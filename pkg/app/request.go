package app

import "fmt"

func MakeErrors(errs []string) {
	for _, err := range errs {
		fmt.Println(err)
	}
}
