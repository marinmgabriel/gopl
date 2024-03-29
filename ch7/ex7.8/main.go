// Exercise 7.8: Many GUIs provide a table widget with a stateful multi-tier sort: the primary sort key is the most recently clicked column head, the secndary sort key is the second-most recently clicked column head, and so on. Define an implementation of sort.Interface for use by such a table. Comapre that approach wit hrepeated sorting using sort.Stable.
package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func by(fld string) func(x, y *Track) bool {
	switch fld {
	case "Title":
		return func(x, y *Track) bool {
			return x.Title < y.Title
		}
	case "Artist":
		return func(x, y *Track) bool {
			return x.Artist < y.Artist
		}
	case "Album":
		return func(x, y *Track) bool {
			return x.Album < y.Album
		}
	case "Year":
		return func(x, y *Track) bool {
			return x.Year < y.Year
		}
	case "Length":
		return func(x, y *Track) bool {
			return x.Length < y.Length
		}
	default:
		return func(x, y *Track) bool {
			return false
		}
	}
}

func main() {
	printTracks(tracks)
	fmt.Println()

	// Sort by Title
	sort.Sort(customSort{tracks, by("Title")})
	printTracks(tracks)
	sort.Sort(sort.Reverse(customSort{tracks, by("Title")}))
	fmt.Println()

	// Sort by Artist
	sort.Sort(customSort{tracks, by("Artist")})
	printTracks(tracks)
	sort.Sort(sort.Reverse(customSort{tracks, by("Artist")}))
	fmt.Println()

	// Sort by Album
	sort.Sort(customSort{tracks, by("Album")})
	printTracks(tracks)
	sort.Sort(sort.Reverse(customSort{tracks, by("Album")}))
	fmt.Println()

	// Sort by Year
	sort.Sort(customSort{tracks, by("Year")})
	printTracks(tracks)
	sort.Sort(sort.Reverse(customSort{tracks, by("Year")}))
	fmt.Println()

	// Sort by Length
	sort.Sort(customSort{tracks, by("Length")})
	printTracks(tracks)
	sort.Sort(sort.Reverse(customSort{tracks, by("Length")}))
	fmt.Println()
}
