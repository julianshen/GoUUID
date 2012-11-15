package main

import (
	"fmt"
	"github.com/julianshen/GoUUID"
)

func main() {
	_uuid, err := uuid.RandomUUID()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(_uuid)
	}

	var str string
	str = _uuid.String()

	myuuid, err := uuid.UUIDFromString(str)

    if err != nil {
        fmt.Println(err)
    }

	fmt.Println(_uuid.version())
}
