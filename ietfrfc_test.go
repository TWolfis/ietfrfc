package ietfRfc_test

import (
	"fmt"
	"testing"

	"github.com/TWolfis/ietfRfc"
)

func TestRfc(t *testing.T) {
	rfcNumber := 1234

	// Test if RFC can be fetched
	rfc, err := ietfRfc.GetRFC(rfcNumber)
	if err != nil {
		t.Error(err)
	}

	// Test if RFC Body is not empty
	if len(rfc.Body) == 0 {
		t.Error("RFC body is empty")
	}

	// Test if RFC Number is correct
	if rfc.Metadata.Series[0].Value != fmt.Sprint(rfcNumber) {
		t.Errorf("RFC number is incorrect, expected %d, got %s", rfcNumber, rfc.Metadata.Series[0].Value)
	}
}

func ExampleGetRFC() {
	rfcNumber := 1234

	// Get RFC
	rfc, err := ietfRfc.GetRFC(rfcNumber)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print RFC Body
	fmt.Println(rfc.Body)
}
