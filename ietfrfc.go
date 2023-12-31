package ietfrfc

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

/*
Download IETF RFCs and their metadata
*/

const (
	// URL pointing towards the RFC document
	rfcURL = "https://www.rfc-editor.org/rfc/rfc"
	// URL pointing towards the RFC metadata document
	refURL = "https://www.rfc-editor.org/refs/ref"
)

// rfc holds the text body of the requested RFC and the metadata concerning the RFC
type RFC struct {
	Body     string
	Title    string
	Authors  string
	Metadata string
}

// GetRFC fetches the RFC and RFC Metadata by number
func GetRFC(rfcNumber int) (RFC, error) {

	if rfcNumber < 1 {
		return RFC{}, fmt.Errorf("RFC number must be greater than 0")
	}

	rfc := RFC{}
	err := rfc.getR(rfcNumber)
	if err != nil {
		return rfc, err
	}

	err = rfc.getRef(rfcNumber)
	if err != nil {
		return rfc, err
	}

	return rfc, nil
}

// getR fetches the RFC
func (r *RFC) getR(rfcNumber int) error {
	// Example: https://www.rfc-editor.org/rfc/rfc791.txt
	rurl := fmt.Sprintf("%s%d.txt", rfcURL, rfcNumber)

	// Get rfc Body
	resp, err := http.Get(rurl)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	r.Body = string(body)

	return nil
}

// getRef fetches the RFC Metadata
func (r *RFC) getRef(rfcNumber int) error {
	var rNum string
	// if the rfcNumber is less than 1000 and bigger than 100, add a leading 0
	if rfcNumber >= 100 && rfcNumber < 1000 {
		rNum = fmt.Sprintf("0%d", rfcNumber)
	} else if rfcNumber >= 10 && rfcNumber < 100 {
		rNum = fmt.Sprintf("00%d", rfcNumber)
	} else if rfcNumber < 10 && rfcNumber > 0 {
		rNum = fmt.Sprintf("000%d", rfcNumber)
	} else {
		rNum = fmt.Sprintf("%d", rfcNumber) // if the rfcNumber is bigger than 1000, just add it as is.
	}

	// Example: "https://www.rfc-editor.org/refs/ref1945.txt"
	refurl := fmt.Sprintf("%s%s.txt", refURL, rNum)
	fmt.Println(refurl)
	// Get metadata
	resp, err := http.Get(refurl)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	rfcInfo := string(body)
	r.Metadata = rfcInfo

	// Split the string based on the " character
	splitInfo := strings.Split(rfcInfo, "\"")

	// The first element is the author(s) and the second is the title of the RFC
	r.Authors = splitInfo[0]
	r.Title = splitInfo[1]

	return nil
}

// GetReference, prints out the title, publication date and author(s) of the RFC
func (r *RFC) GetReference() {
	fmt.Println("Title:", r.Title)
	fmt.Println("Authors:", r.Authors)
	fmt.Println("Publication information:", r.Metadata)
	fmt.Println()
}

// String prints out the RFC Body
func (r *RFC) String() string {
	return fmt.Sprintln(r.Body)
}
