package ietfRfc

import "encoding/xml"

type Reference struct {
	XMLName xml.Name     `xml:"reference"`
	Anchor  string       `xml:"anchor,attr"`
	Target  string       `xml:"target,attr"`
	Front   Front        `xml:"front"`
	Series  []SeriesInfo `xml:"seriesInfo"`
}

type Front struct {
	XMLName xml.Name `xml:"front"`
	Title   string   `xml:"title"`
	Authors []Author `xml:"author"`
	Date    Date     `xml:"date"`
}

type Author struct {
	XMLName  xml.Name `xml:"author"`
	Fullname string   `xml:"fullname,attr"`
	Initials string   `xml:"initials,attr"`
	Surname  string   `xml:"surname,attr"`
}

type Date struct {
	XMLName xml.Name `xml:"date"`
	Month   string   `xml:"month,attr"`
	Year    string   `xml:"year,attr"`
}

type SeriesInfo struct {
	XMLName xml.Name `xml:"seriesInfo"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}
