package nfoparser

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"
)

const (
	movieNfoFile = "movie.nfo"
)

func TestMovieMarshal(t *testing.T) {
	movie := Movie{}
	movie.Title = "Movie title"
	movie.OriginalTitle = "Original title"
	movie.SortTitle = "Sort title"

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
	movie.Ratings.Rating = append(movie.Ratings.Rating, rating1, rating2)

	movie.UserRating = 0
	movie.Top250 = 100
	movie.Outline = "Outline"
	movie.Plot = "Plot"
	movie.Tagline = "Tagline"
	movie.Runtime = 90

	thumb1 := Thumb{
		Spoof:   "spoof",
		Aspect:  "aspect",
		Cache:   "cache",
		Preview: "preview",
		Link:    "https://anywhere.com",
	}
	movie.Thumb = append(movie.Thumb, thumb1, thumb1)

	thumb2 := Thumb{
		Preview: "preview",
		Colors:  "colors",
		Link:    "https://fanart.org",
	}
	fanart := Fanart{}
	fanart.Thumb = append(fanart.Thumb, thumb2)
	movie.Fanart = &fanart

	movie.MPAA = "16"
	movie.Playcount = 1
	movie.Lastplayed = ""
	movie.Id = 7
	movie.Uniqueid = append(movie.Uniqueid, UniqueId{
		Type:    "imdb",
		Id:      "tt123",
	})
	movie.Uniqueid = append(movie.Uniqueid, UniqueId{
		Type:    "tmdb",
		Id:      "tt321",
	})
	movie.Genre = append(movie.Genre, "Comedy", "Thriller")

	movie.FileInfo.StreamDetails.Video = append(movie.FileInfo.StreamDetails.Video, StreamVideo{
		Codec:             "hevc",
		Aspect:            1.85,
		Width:             1920,
		Height:            1080,
		DurationInSeconds: 1230,
		StereoMode:        "",
	})

	movie.Actor = append(movie.Actor, Actor{
		Name:  "Willie Wicket",
		Role:  "Hero",
		Order: 1,
		Thumb: "http://thumb.xx",
	}, Actor{
		Name:  "Martha Marshal",
		Role:  "Princess",
		Order: 2,
		Thumb: "https://thumb2.xx",
	})

	b, err := xml.MarshalIndent(movie, " ", "  ")
	if err != nil {
		t.Errorf("Error marshalling movie: %v\n", err)
		t.Fail()
	}
	t.Log(string(b))

	m := Movie{}
	err = xml.Unmarshal(b, &m)
	if err != nil {
		t.Logf("Error unmarshalling movie: %v\n", err)
		t.Fail()
	}
}

func TestMovieUnmarshal(t *testing.T) {
	b, err := ioutil.ReadFile(movieNfoFile)
	if err != nil {
		t.Fatalf("Could not read file '%s': %v\n", movieNfoFile, err)
	}

	m := Movie{}
	err = xml.Unmarshal(b, &m)
	if err != nil {
		t.Logf("Error unmarshalling movie: %v\n", err)
		t.Fail()
	}
	t.Logf("Title: %s\n", m.Title)
}

func TestMovieNfoReader(t *testing.T) {
	f, err := os.Open(movieNfoFile)
	if err != nil {
		t.Fatalf("Could not read file '%s': %v\n", movieNfoFile, err)
	}
	m, err := ReadMovieNfo(f)
	if err != nil {
		t.Errorf("Error reading movie file: %v\n", err)
		t.Fail()
	}
	t.Logf("Title: %s\n", m.Title)
}
