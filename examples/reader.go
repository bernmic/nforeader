package main

import (
	"flag"
	"fmt"
	"github.com/bernmic/nfoparser"
	"os"
	"path/filepath"
	"strings"
)

func printUsage() {
	fmt.Println("Usage: reader movie|tvshow|episode <file|dir>")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		printUsage()
		return
	}
	if args[0] == "movie" {
		s, err := os.Stat(args[1])
		if err != nil {
			fmt.Printf("Error getting info about %s: %v\n", args[1], err)
			os.Exit(1)
		}
		if s.IsDir() {
			_, err := readAllMovies(args[1])
			if err != nil {
				os.Exit(1)
			}
			return
		}
		m, err := readMovieNfo(args[1])
		if err == nil {
			fmt.Printf("Successfully read movie %s\n", m.Title)
		}
		return
	}

	if args[0] == "tvshow" {
		s, err := os.Stat(args[1])
		if err != nil {
			fmt.Printf("Error getting info about %s: %v\n", args[1], err)
			os.Exit(1)
		}
		if s.IsDir() {
			_, err := readAllTVShows(args[1])
			if err != nil {
				os.Exit(1)
			}
			return
		}
		m, err := readTVShowNfo(args[1])
		if err == nil {
			fmt.Printf("Successfully read tvshow %s\n", m.Title)
		}
		return
	}

	if args[0] == "episode" {
		s, err := os.Stat(args[1])
		if err != nil {
			fmt.Printf("Error getting info about %s: %v\n", args[1], err)
			os.Exit(1)
		}
		if s.IsDir() {
			_, err := readAllEpisodes(args[1])
			if err != nil {
				os.Exit(1)
			}
			return
		}
		m, err := readEpisodeNfo(args[1])
		if err == nil {
			fmt.Printf("Successfully read episode %s - S%02dE%02d - %s\n", m.ShowTitle, m.Season, m.Episode, m.Title)
		}
		return
	}
	printUsage()
	os.Exit(1)
}

func readMovieNfo(filename string) (*nfoparser.Movie, error) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening %s: %v\n", filename, err)
		return nil, err
	}
	m, err := nfoparser.ReadMovieNfo(f)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", filename, err)
		return nil, err
	}
	return m, nil
}

func readAllMovies(dirname string) ([]*nfoparser.Movie, error) {
	movies := make([]*nfoparser.Movie, 0)
	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && strings.HasSuffix(path, "movie.nfo") {
			m, err := readMovieNfo(path)
			if err == nil {
				movies = append(movies, m)
				fmt.Printf("Successfully read movie %s\n", m.Title)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error scanning dir %s: %v\n", dirname, err)
		return nil, err
	}
	return movies, nil
}

func readTVShowNfo(filename string) (*nfoparser.TVShow, error) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening %s: %v\n", filename, err)
		return nil, err
	}
	m, err := nfoparser.ReadTVShowNfo(f)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", filename, err)
		return nil, err
	}
	return m, nil
}

func readAllTVShows(dirname string) ([]*nfoparser.TVShow, error) {
	tvShows := make([]*nfoparser.TVShow, 0)
	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && strings.HasSuffix(path, "tvshow.nfo") {
			m, err := readTVShowNfo(path)
			if err == nil {
				tvShows = append(tvShows, m)
				fmt.Printf("Successfully read tvshow %s\n", m.Title)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error scanning dir %s: %v\n", dirname, err)
		return nil, err
	}
	return tvShows, nil
}

func readEpisodeNfo(filename string) (*nfoparser.Episode, error) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening %s: %v\n", filename, err)
		return nil, err
	}
	m, err := nfoparser.ReadEpisodeNfo(f)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", filename, err)
		return nil, err
	}
	return m, nil
}

func readAllEpisodes(dirname string) ([]*nfoparser.Episode, error) {
	tvShows := make([]*nfoparser.Episode, 0)
	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && strings.HasSuffix(path, ".nfo") && !strings.HasSuffix(path, "tvshow.nfo") {
			m, err := readEpisodeNfo(path)
			if err == nil {
				tvShows = append(tvShows, m)
				fmt.Printf("Successfully read episode %s - S%02dE%02d - %s\n", m.ShowTitle, m.Season, m.Episode, m.Title)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error scanning dir %s: %v\n", dirname, err)
		return nil, err
	}
	return tvShows, nil
}
