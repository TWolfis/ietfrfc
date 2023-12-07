package ietfRfc

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

/*
fetch and handles rfc's
*/

const (
	rfcURL = "https://www.rfc-editor.org/rfc/rfc"
	mdURL  = "https://bib.ietf.org/public/rfc/bibxml/reference.RFC."
)

type Rfc struct {
	Body     string
	Metadata Reference
}

// Global function to retrieve RFC and metadata
func GetRFC(rfcNumber int) (Rfc, error) {
	rfc := Rfc{}
	err := rfc.getR(rfcNumber)
	if err != nil {
		return rfc, err
	}

	err = rfc.getMetadata(rfcNumber)
	if err != nil {
		return rfc, err
	}

	return rfc, nil
}

// Get the RFC text file
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

// Get the metadata
func (r *Rfc) getMetadata(rfcNumber int) error {

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

// GetReference, print out the title, publication date and author(s) of the RFC
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

// Stringer
func (r Rfc) String() string {
	return fmt.Sprintln(r.Body)
}
