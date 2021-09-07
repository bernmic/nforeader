package nfoparser

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"
)

const (
	tvshowNfoFile = "tvshow.nfo"
)

func TestTVShowMarshal(t *testing.T) {
	tvshow := TVShow{
		Title:          "Title",
		OriginalTitle:  "Original title",
		ShowTitle:      "Show title",
		UserRating:     0,
		Top250:         1,
		Season:         1,
		Episode:        1,
		DisplaySeason:  -1,
		DisplayEpisode: -1,
		Outline:        "Outline",
		Plot:           "Plot",
		Tagline:        "Tagline",
		Runtime:        1000,
		MPAA:           "23",
		Playcount:      1,
		Lastplayed:     "",
		Id:             123,
		Premiered:      "",
		Year:           "2021",
		Status:         "",
		Code:           "",
		Aired:          "2021-01-01",
		Studio:         "Studio",
		DateAdded:      "2021-05-01",
	}

	rating1 := Rating{
		Value:   1.001,
		Votes:   1,
		Name:    "imdb",
		Max:     10,
		Default: true,
	}
	rating2 := Rating{
		Value: 2.002,
		Votes: 2,
		Name:  "tmdb",
		Max:   10,
	}
	tvshow.Ratings.Rating = append(tvshow.Ratings.Rating, rating1, rating2)

	thumb1 := Thumb{
		Spoof:   "spoof",
		Aspect:  "aspect",
		Cache:   "cache",
		Preview: "preview",
		Link:    "https://anywhere.com",
	}
	tvshow.Thumb = append(tvshow.Thumb, thumb1, thumb1)

	thumb2 := Thumb{
		Preview: "preview",
		Colors:  "colors",
		Link:    "https://fanart.org",
	}
	fanart := Fanart{}
	fanart.Thumb = append(fanart.Thumb, thumb2)
	tvshow.Fanart = &fanart

	tvshow.Uniqueid = append(tvshow.Uniqueid, UniqueId{
		Type: "imdb",
		Id:   "tt123",
	})
	tvshow.Uniqueid = append(tvshow.Uniqueid, UniqueId{
		Type: "tmdb",
		Id:   "tt321",
	})
	tvshow.Genre = append(tvshow.Genre, "Comedy", "Thriller")

	tvshow.Actor = append(tvshow.Actor, Actor{
		Name:  "Willie Wicket",
		Role:  "Hero",
		Order: 1,
		Thumb: "http://thumb.xx",
	}, Actor{
		Name:  "Martha Marshal",
		Role:  "Princess",
		Order: 2,
		Thumb: "http://thumb2.xx",
	})

	b, err := xml.MarshalIndent(tvshow, " ", "  ")
	if err != nil {
		t.Errorf("Error marshalling tvshow: %v\n", err)
		t.Fail()
	}
	t.Log(string(b))

	m := TVShow{}
	err = xml.Unmarshal(b, &m)
	if err != nil {
		t.Logf("Error unmarshalling tvshow: %v\n", err)
		t.Fail()
	}
}

func TestTVShowUnmarshal(t *testing.T) {
	b, err := ioutil.ReadFile(tvshowNfoFile)
	if err != nil {
		t.Fatalf("Could not read file '%s': %v\n", tvshowNfoFile, err)
	}

	m := TVShow{}
	err = xml.Unmarshal(b, &m)
	if err != nil {
		t.Logf("Error unmarshalling tvshow: %v\n", err)
		t.Fail()
	}
	t.Logf("Title: %s\n", m.Title)
}

func TestTVShowNfoReader(t *testing.T) {
	f, err := os.Open(tvshowNfoFile)
	if err != nil {
		t.Fatalf("Could not read file '%s': %v\n", tvshowNfoFile, err)
	}
	m, err := ReadTVShowNfo(f)
	if err != nil {
		t.Errorf("Error reading tvshow file: %v\n", err)
		t.Fail()
	}
	t.Logf("Title: %s\n", m.Title)

	// test if not file is not readable
	f, err = os.OpenFile("x.out", os.O_CREATE|os.O_WRONLY, 0644)
	defer os.Remove("x.out")
	defer f.Close()
	m, err = ReadTVShowNfo(f)
	if err == nil {
		t.Error("Accepted nil as reader")
		t.Fail()
	}
	// test with wrong file format
	f, err = os.Open(episodeNfoFile)
	if err != nil {
		t.Fatalf("Could not open file '%s': %v\n", episodeNfoFile, err)
	}
	m, err = ReadTVShowNfo(f)
	if err == nil {
		t.Error("Accepted episode as tvshow", err)
		t.Fail()
	}
}
