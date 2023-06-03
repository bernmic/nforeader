package nfoparser

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"
)

const (
	episodeNfoFile = "episode.nfo"
)

func TestEpisodeMarshal(t *testing.T) {
	episode := Episode{
		Title:          "Title",
		ShowTitle:      "ShowTitle",
		UserRating:     1,
		Top250:         1,
		Season:         1,
		Episode:        1,
		DisplaySeason:  -1,
		DisplayEpisode: -1,
		Outline:        "Outline",
		Plot:           "Plot",
		Tagline:        "Tagline",
		Runtime:        23,
		MPAA:           "12",
		Playcount:      0,
		Lastplayed:     "",
		Id:             "1",
		Credits:        "Credits",
		Director:       "Director",
		Premiered:      "2021-01-01",
		Year:           "2021",
		Status:         "",
		Code:           "",
		Aired:          "2021-03-01",
		Studio:         "Studio",
		Trailer:        "",
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
	episode.Ratings.Rating = append(episode.Ratings.Rating, rating1, rating2)

	thumb1 := Thumb{
		Spoof:   "spoof",
		Aspect:  "aspect",
		Cache:   "cache",
		Preview: "preview",
		Link:    "https://anywhere.com",
	}
	episode.Thumb = append(episode.Thumb, thumb1, thumb1)

	episode.Uniqueid = append(episode.Uniqueid, UniqueId{
		Type: "imdb",
		Id:   "tt123",
	})
	episode.Uniqueid = append(episode.Uniqueid, UniqueId{
		Type: "tmdb",
		Id:   "tt321",
	})
	episode.Genre = append(episode.Genre, "Comedy", "Thriller")

	episode.FileInfo.StreamDetails.Video = append(episode.FileInfo.StreamDetails.Video, StreamVideo{
		Codec:             "hevc",
		Aspect:            "1.85",
		Width:             1920,
		Height:            1080,
		DurationInSeconds: 1230,
		StereoMode:        "",
	})

	episode.Actor = append(episode.Actor, Actor{
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

	b, err := xml.MarshalIndent(episode, " ", "  ")
	if err != nil {
		t.Errorf("Error marshalling episode: %v\n", err)
		t.Fail()
	}
	t.Log(string(b))

	m := Episode{}
	err = xml.Unmarshal(b, &m)
	if err != nil {
		t.Logf("Error unmarshalling episode: %v\n", err)
		t.Fail()
	}
}

func TestEpisodeUnmarshal(t *testing.T) {
	b, err := ioutil.ReadFile(episodeNfoFile)
	if err != nil {
		t.Fatalf("Could not read file '%s': %v\n", episodeNfoFile, err)
	}

	m := Episode{}
	err = xml.Unmarshal(b, &m)
	if err != nil {
		t.Logf("Error unmarshalling episode: %v\n", err)
		t.Fail()
	}
	t.Logf("Title: %s\n", m.Title)
}

func TestEpisodeNfoReader(t *testing.T) {
	f, err := os.Open(episodeNfoFile)
	if err != nil {
		t.Fatalf("Could not read file '%s': %v\n", episodeNfoFile, err)
	}
	m, err := ReadEpisodeNfo(f)
	if err != nil {
		t.Errorf("Error reading episode file: %v\n", err)
		t.Fail()
	}
	t.Logf("Title: %s\n", m.Title)

	// test if not file is not readable
	f, err = os.OpenFile("x.out", os.O_CREATE|os.O_WRONLY, 0644)
	defer os.Remove("x.out")
	defer f.Close()
	m, err = ReadEpisodeNfo(f)
	if err == nil {
		t.Error("Accepted nil as reader")
		t.Fail()
	}
	// test with wrong file format
	f, err = os.Open(movieNfoFile)
	if err != nil {
		t.Fatalf("Could not open file '%s': %v\n", movieNfoFile, err)
	}
	m, err = ReadEpisodeNfo(f)
	if err == nil {
		t.Error("Accepted movie as episode", err)
		t.Fail()
	}
}
