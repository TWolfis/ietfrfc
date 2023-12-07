package ietfrfc_test

import (
	"fmt"
	"testing"

	"github.com/TWolfis/ietfrfc"
)

func TestRfc(t *testing.T) {
	rfcNumber := 2328

	// Test if RFC can be fetched
	rfc, err := ietfrfc.GetRFC(rfcNumber)
	if err != nil {
		t.Error(err)
	}

	// Test if RFC Body is not empty
	if len(rfc.Body) == 0 {
		t.Error("RFC body is empty")
	}

	// Test if RFC Title equals OSPF Version 2
	if rfc.Title != "OSPF Version 2" {
		t.Errorf("RFC title is incorrect, expected OSPF Version 2, got %s", rfc.Title)
	}
}

func ExampleGetRFC() {
	rfcNumber := 2328

	// Get RFC
	rfc, err := ietfrfc.GetRFC(rfcNumber)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print RFC Body
	fmt.Println(rfc.Body)
}
