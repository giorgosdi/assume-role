package main

import (
	"fmt"

	"github.com/giorgosdi/assume-role/pkg/api"
)

func main() {

	state := map[string]string{
		"profile": "default",
		"role":    "test",
		"account": "784365650852",
	}

	resp, err := api.AssumeRole(state)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}
