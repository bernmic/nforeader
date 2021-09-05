package main

import (
	"flag"
	"fmt"
	"github.com/bernmic/nfoparser"
	"os"
)

func printUsage() {
	fmt.Println("Usage: reader movie|tvshow|episode <file>")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		printUsage()
		return
	}
	if args[0] == "movie" {
		f, err := os.Open(args[1])
		if err != nil {
			fmt.Printf("Error opening %s: %v\n", args[1], err)
			os.Exit(1)
		}
		m, err := nfoparser.ReadMovieNfo(f)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", args[1], err)
			os.Exit(1)
		}
		fmt.Printf("Successfully read movie %s\n", m.Title)
		return
	}

	if args[0] == "tvshow" {
		f, err := os.Open(args[1])
		if err != nil {
			fmt.Printf("Error opening %s: %v\n", args[1], err)
			os.Exit(1)
		}
		m, err := nfoparser.ReadTVShowNfo(f)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", args[1], err)
			os.Exit(1)
		}
		fmt.Printf("Successfully read tvshow %s\n", m.Title)
		return
	}

	if args[0] == "episode" {
		f, err := os.Open(args[1])
		if err != nil {
			fmt.Printf("Error opening %s: %v\n", args[1], err)
			os.Exit(1)
		}
		m, err := nfoparser.ReadEpisodeNfo(f)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", args[1], err)
			os.Exit(1)
		}
		fmt.Printf("Successfully read tvshow %s\n", m.Title)
		return
	}
	printUsage()
	os.Exit(1)
}
