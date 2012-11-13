package main

import (
	"fmt"
	uuid "github.com/julianshen/GoUUID"
)

func main() {
	_uuid, err := uuid.RandomUUID()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(_uuid)
	}
}
