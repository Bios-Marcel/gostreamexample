package gostreamexample_test

import (
	"fmt"
	"testing"

	"github.com/Bios-Marcel/gostreamexample/stream"
)

func TestLib(t *testing.T) {
	result := stream.
		StreamString([]string{"Hallo", "filterme", "Welt"}).
		Filter(func(value string) bool { return value != "filterme" }).
		Collect()

	if len(result) != 2 {
		t.Errorf("Length should have been 2, but was %d", len(result))
	}

	fmt.Println(result[0], result[1])
}
