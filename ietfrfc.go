package ietfrfc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/caltechlibrary/bibtex"
	"github.com/pkg/errors"
)

/*
Download IETF RFCs and their metadata
*/

// rfc holds the text body of the requested RFC and the metadata concerning the RFC
type RFC struct {
	Series       string `json:"series,omitempty"`
	Number       string `json:"number,omitempty"`
	Howpublished string `json:"howpublished,omitempty"`
	Publisher    string `json:"publisher,omitempty"`
	Doi          string `json:"doi,omitempty"`
	URL          string `json:"url,omitempty"`
	Author       string `json:"author,omitempty"`
	Title        string `json:"title,omitempty"`
	Pagetotal    string `json:"pagetotal,omitempty"`
	Year         string `json:"year,omitempty"`
	Month        string `json:"month,omitempty"`
	Day          string `json:"day,omitempty"`
	Abstract     string `json:"abstract,omitempty"`
	Body         string `json:"body,omitempty"`
}

// GetRFC fetches the RFC and RFC Metadata by number
func Get(Number int) (*RFC, error) {
	rfc := RFC{}
	if Number < 1 {
		return &rfc, errors.New("RFC number must be greater than 0")
	}

	c1 := make(chan error)
	c2 := make(chan error)

	go rfc.getRFC(Number, c1)
	go rfc.getRef(Number, c2)

	for i := 0; i < 2; i++ {
		select {
		case err := <-c1:
			if err != nil {
				return &rfc, err
			}
		case err := <-c2:
			if err != nil {
				return &rfc, err
			}
		}
	}

	return &rfc, nil
}

// getRFC fetches the RFC by Number
func (r *RFC) getRFC(Number int, c chan error) {
	// Example: https://www.rfc-editor.org/rfc/rfc791.txt
	rfcURL := fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", Number)

	// Get rfc Body
	resp, err := http.Get(rfcURL)
	if err != nil {
		c <- err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c <- err
	}

	r.Body = string(body)

	c <- nil
}

// getRef fetches the RFC Metadata
func (r *RFC) getRef(Number int, c chan error) {
	refURL := fmt.Sprintf("https://datatracker.ietf.org/doc/rfc%d/bibtex/", Number)
	resp, err := http.Get(refURL)
	if err != nil {
		c <- err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c <- err
	}

	bt, err := bibtex.Parse(body)
	if err != nil {
		c <- err
	}

	// Get first element from parser and from there extract the tags
	if len(bt) < 1 {
		c <- errors.New("no tags found")
		return
	}
	t := bt[0]

	// set chars to trim
	trimChars := "{}"

	r.Series = strings.Trim(t.Tags["series"], trimChars)
	r.Number = strings.Trim(t.Tags["number"], trimChars)
	r.Howpublished = strings.Trim(t.Tags["howpublished"], trimChars)
	r.Publisher = strings.Trim(t.Tags["publisher"], trimChars)
	r.Doi = strings.Trim(t.Tags["doi"], trimChars)
	r.URL = strings.Trim(t.Tags["url"], trimChars)
	r.Author = strings.Trim(t.Tags["author"], trimChars)
	r.Title = strings.Trim(t.Tags["title"], trimChars)
	r.Pagetotal = strings.Trim(t.Tags["pagetotal"], trimChars)
	r.Year = strings.Trim(t.Tags["year"], trimChars)
	r.Month = strings.Trim(t.Tags["month"], trimChars)
	r.Day = strings.Trim(t.Tags["day"], trimChars)
	r.Abstract = strings.Trim(t.Tags["abstract"], trimChars)

	c <- nil
}

// String prints out the RFC Body
func (r *RFC) String() string {
	jsonData, _ := json.MarshalIndent(r, "", "  ") // Use "  " for two space indentation

	// Print the pretty-printed JSON
	return string(jsonData)
}
