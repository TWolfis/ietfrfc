package ietfrfc_test

import (
	"fmt"
	"testing"

	"github.com/TWolfis/ietfrfc"
)

func TestRfc(t *testing.T) {
	number := 2328

	// Test if RFC can be fetched
	rfc, err := ietfrfc.Get(number)
	if err != nil {
		t.Error(err)
	}

	// Test if RFC Body is not empty
	if len(rfc.Body) == 0 {
		t.Error("RFC body is empty")
	}

	// Test if RFC Title equals OSPF Version 2
	x := "OSPF Version 2"
	if rfc.Title != x {
		t.Errorf("RFC title is incorrect, expected %v, got %v", x, rfc.Title)
	}
}

func ExampleGet() {
	number := 2328

	// Get
	rfc, err := ietfrfc.Get(number)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print RFC Body
	fmt.Println(rfc.Body)
}
