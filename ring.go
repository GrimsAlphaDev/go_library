package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5)

	// Manual
	// data.Value = "Geralt"

	// data = data.Next()

	// data.Value = "Yennefer"

	// data = data.Next()

	// data.Value = "Triss"

	// data = data.Next()

	// data.Value = "Ciri"

	// data = data.Next()

	// data.Value = "Jaskier"

	// iteration
	for i := 1; i <= data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value)
	})
}
