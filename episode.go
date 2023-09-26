package nfoparser

import (
	"encoding/xml"
	"io"
)

type Episode struct {
	XMLName   xml.Name `xml:"episodedetails"`
	Title     string   `xml:"title,omitempty"`
	ShowTitle string   `xml:"showtitle,omitempty"`
	Ratings   struct {
		Rating []Rating `xml:"rating,omitempty"`
	} `xml:"ratings"`
	UserRating     float64    `xml:"userrating,omitempty"`
	Top250         int64      `xml:"top250,omitempty"`
	Season         int64      `xml:"season,omitempty"`
	Episode        int64      `xml:"episode,omitempty"`
	DisplaySeason  int64      `xml:"displayseason,omitempty"`
	DisplayEpisode int64      `xml:"displayepisode,omitempty"`
	Outline        string     `xml:"outline,omitempty"`
	Plot           string     `xml:"plot,omitempty"`
	Tagline        string     `xml:"tagline,omitempty"`
	Runtime        int64      `xml:"runtime,omitempty"`
	Thumb          []Thumb    `xml:"thumb,omitempty"`
	MPAA           string     `xml:"mpaa,omitempty"`
	Playcount      int64      `xml:"playcount,omitempty"`
	Lastplayed     string     `xml:"lastplayed,omitempty"`
	Id             string     `xml:"id,omitempty"`
	Uniqueid       []UniqueId `xml:"uniqueid,omitempty"`
	Genre          []string   `xml:"genre,omitempty"`
	Credits        string     `xml:"credits,omitempty"`
	Director       string     `xml:"director,omitempty"`
	Premiered      string     `xml:"premiered,omitempty"`
	Year           string     `xml:"year,omitempty"`
	Status         string     `xml:"status,omitempty"`
	Code           string     `xml:"code,omitempty"`
	Aired          string     `xml:"aired,omitempty"`
	Studio         string     `xml:"studio,omitempty"`
	Trailer        string     `xml:"trailer,omitempty"`
	FileInfo       struct {
		StreamDetails struct {
			Video    []StreamVideo    `xml:"video,omitempty"`
			Audio    []StreamAudio    `xml:"audio,omitempty"`
			Subtitle []StreamSubtitle `xml:"subtitle,omitempty"`
		} `xml:"streamdetails,omitempty"`
	} `xml:"fileinfo,omitempty"`
	Actor  []Actor `xml:"actor,omitempty"`
	Resume struct {
		Position float64 `xml:"position,omitempty"`
		Total    float64 `xml:"total,omitempty"`
	} `xml:"resume,omitempty"`
	DateAdded string `xml:"dateadded,omitempty"`
}

func ReadEpisodeNfo(r io.Reader) (*Episode, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	m := Episode{}
	err = xml.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
