package nfoparser

import (
	"encoding/xml"
	"io"
	"io/ioutil"
)

type Album struct {
	XMLName       xml.Name `xml:"album"`
	Title          string     `xml:"title,omitempty"`
	Rating         int64      `xml:"rating,omitempty"`
	UserRating     int64      `xml:"userrating,omitempty"`
	Genre          string     `xml:"genre,omitempty"`
	Style          string     `xml:"style,omitempty"`
	Mood           string     `xml:"mood,omitempty"`
	Theme          string     `xml:"theme,omitempty"`
	Compilation    bool       `xml:"compilation,omitempty"`
	Year           string     `xml:"year,omitempty"`
	ReleaseDate    string     `xml:"releasedate,omitempty"`
	Review         string     `xml:"review,omitempty"`
	Type           string     `xml:"type,omitempty"`
	Label          string     `xml:"label,omitempty"`
	AlbumArtistCredits struct {
		MusicBrainzArtistID   string  `xml:"musicBrainzArtistID,omitempty"`
		Artist                string  `xml:"artist,omitempty"`
	} `xml:"albumArtistCredits,omitempty"`
}


func ReadAlbumNfo(r io.Reader) (*Album, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	m := Album{}
	err = xml.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
