package main

import (
	"fmt"

	"github.com/Bios-Marcel/gostreamexample/stream"
)

func main() {
	result := stream.
		StreamString([]string{"Hallo", "filterme", "Welt"}).
		Filter(func(value string) bool { return value != "filterme" }).
		Collect()

	if len(result) != 2 {
		fmt.Printf("Length should have been 2, but was %d", len(result))
	} else {
		fmt.Println(result[0], result[1])
	}
}
