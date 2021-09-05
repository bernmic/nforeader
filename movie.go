package nfoparser

import (
	"encoding/xml"
	"io"
	"io/ioutil"
)

type Movie struct {
	XMLName       xml.Name `xml:"movie"`
	Title         string   `xml:"title,omitempty"`
	OriginalTitle string   `xml:"originaltitle,omitempty"`
	SortTitle     string   `xml:"sorttitle,omitempty"`
	Ratings       struct {
		Rating []Rating `xml:"rating,omitempty"`
	} `xml:"ratings"`
	UserRating int64      `xml:"userrating,omitempty"`
	Top250     int64      `xml:"top250,omitempty"`
	Outline    string     `xml:"outline,omitempty"`
	Plot       string     `xml:"plot,omitempty"`
	Tagline    string     `xml:"tagline,omitempty"`
	Runtime    int64      `xml:"runtime,omitempty"`
	Thumb      []Thumb    `xml:"thumb,omitempty"`
	Fanart     *Fanart    `xml:"fanart,omitempty"`
	MPAA       string     `xml:"mpaa,omitempty"`
	Playcount  int64      `xml:"playcount,omitempty"`
	Lastplayed string     `xml:"lastplayed,omitempty"`
	Id         int64      `xml:"id,omitempty"`
	Uniqueid   []UniqueId `xml:"uniqueid,omitempty"`
	Genre      []string   `xml:"genre,omitempty"`
	Tag        []string   `xml:"tag,omitempty"`
	Set        struct {
		Name     string `xml:"name,omitempty"`
		Overview string `xml:"overview,omitempty"`
	} `xml:"set,omitempty"`
	Country   string `xml:"country,omitempty"`
	Credits   string `xml:"credits,omitempty"`
	Director  string `xml:"director,omitempty"`
	Premiered string `xml:"premiered,omitempty"`
	Year      string `xml:"year,omitempty"`
	Status    string `xml:"status,omitempty"`
	Code      string `xml:"code,omitempty"`
	Aired     string `xml:"aired,omitempty"`
	Studio    string `xml:"studio,omitempty"`
	Trailer   string `xml:"trailer,omitempty"`
	FileInfo  struct {
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

type Rating struct {
	Value   float64 `xml:"value,omitempty"`
	Votes   int64   `xml:"votes,omitempty"`
	Name    string  `xml:"name,attr,omitempty"`
	Max     int64   `xml:"max,attr,omitempty"`
	Default bool    `xml:"default,attr,omitempty"`
}

type Fanart struct {
	Thumb []Thumb `xml:"thumb,omitempty"`
}

type Thumb struct {
	Spoof   string `xml:"spoof,omitempty,attr"`
	Aspect  string `xml:"aspect,omitempty,attr"`
	Cache   string `xml:"cache,omitempty,attr"`
	Preview string `xml:"preview,omitempty,attr"`
	Colors  string `xml:"colors,omitempty,attr"`
	Link    string `xml:",chardata"`
}

type UniqueId struct {
	Type    string `xml:"type,omitempty,attr"`
	Default bool   `xml:"default,omitempty,attr"`
	Id      string `xml:",chardata"`
}

type StreamVideo struct {
	Codec             string  `xml:"codec,omitempty"`
	Aspect            float64 `xml:"aspect,omitempty"`
	Width             int64   `xml:"width,omitempty"`
	Height            int64   `xml:"height,omitempty"`
	DurationInSeconds int64   `xml:"durationinseconds,omitempty"`
	StereoMode        string  `xml:"stereomode,omitempty"`
}

type StreamAudio struct {
	Codec    string `xml:"codec,omitempty"`
	Language string `xml:"language,omitempty"`
	Channels int64  `xml:"channels,omitempty"`
}

type StreamSubtitle struct {
	Language string `xml:"language,omitempty"`
}

type Actor struct {
	Name  string `xml:"name,omitempty"`
	Role  string `xml:"role,omitempty"`
	Order int64  `xml:"order,omitempty"`
	Thumb string `xml:"thumb,omitempty"`
}

func ReadMovieNfo(r io.Reader) (*Movie, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	m := Movie{}
	err = xml.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
