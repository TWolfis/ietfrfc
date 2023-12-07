package ietfRfc

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

/*
Download IETF RFCs and their metadata
*/

// The base URLs for retrieving RFC documents and their metadata info
const (
	rfcURL = "https://www.rfc-editor.org/rfc/rfc"
	mdURL  = "https://bib.ietf.org/public/rfc/bibxml/reference.RFC."
)

// Rfc holds the text body of the requested RFC and the metadata concerning the RFC
type Rfc struct {
	Body     string
	Metadata Reference
}

// GetRFC fetches the RFC and RFC Metadata by number
func GetRFC(rfcNumber int) (Rfc, error) {
	rfc := Rfc{}
	err := rfc.getR(rfcNumber)
	if err != nil {
		return rfc, err
	}

	err = rfc.getM(rfcNumber)
	if err != nil {
		return rfc, err
	}

	return rfc, nil
}

// getR fetches the RFC
func (r *Rfc) getR(rfcNumber int) error {
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

// getM fetches the RFC Metadata
func (r *Rfc) getM(rfcNumber int) error {

	// Example: https://bib.ietf.org/public/rfc/bibxml/reference.RFC.791.xml
	murl := fmt.Sprintf("%s%d.xml", mdURL, rfcNumber)

	// Get metadata
	resp, err := http.Get(murl)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(body, &r.Metadata)
	if err != nil {
		return err
	}

	return nil
}

// GetReference, prints out the title, publication date and author(s) of the RFC
func (r Rfc) GetReference() {
	title := r.Metadata.Front.Title
	month := r.Metadata.Front.Date.Month
	year := r.Metadata.Front.Date.Year

	fmt.Println(title)
	fmt.Println(month, year)

	for _, author := range r.Metadata.Front.Authors {
		fmt.Printf("%s ", author.Fullname)
	}
	fmt.Println()
}

// String prints out the RFC Body
func (r Rfc) String() string {
	return fmt.Sprintln(r.Body)
}
