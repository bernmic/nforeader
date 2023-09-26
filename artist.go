package nfoparser

import (
	"encoding/xml"
	"io"
)

type Artist struct {
	XMLName        xml.Name `xml:"Artist"`
	Name           string   `xml:"name,omitempty"`
	SortName       string   `xml:"sortname,omitempty"`
	Type           string   `xml:"type,omitempty"`
	Gender         string   `xml:"gender,omitempty"`
	Disambiguation string   `xml:"disambiguation,omitempty"`
	Genre          string   `xml:"genre,omitempty"`
	Style          string   `xml:"style,omitempty"`
	Mood           string   `xml:"mood,omitempty"`
	YearsActive    string   `xml:"yearsactive,omitempty"`
	Born           string   `xml:"born,omitempty"`
	Formed         string   `xml:"formed,omitempty"`
	Biography      string   `xml:"biography,omitempty"`
	Died           string   `xml:"died,omitempty"`
	Disbanded      string   `xml:"disbanded,omitempty"`
}

func ReadArtistNfo(r io.Reader) (*Artist, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	m := Artist{}
	err = xml.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
