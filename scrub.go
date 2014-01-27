package main

import (
	"os"
	"fmt"
	"strings"
	"path/filepath"
	id3 "github.com/mikkyang/id3-go"
)

func search_for_key_slash_pattern(title string) string {
	return_value := ""
	for j := 12; j >= 1; j-- {
		prefix1 := fmt.Sprintf("%dA/", j)
		if strings.HasPrefix(title, prefix1) {
			return_value = prefix1
			break
		}
		prefix2 := fmt.Sprintf("%dB/", j)
		if strings.HasPrefix(title, prefix2) {
			return_value = prefix2
			break
		}
	}
	return return_value
}

func search_for_key_pattern(title string) string {
	return_value := ""
	found_flag := false
	for j := 12; j >= 1; j-- {
		for i := 1; i <= 10; i++ {
			prefix := fmt.Sprintf("%dA - %d - ", j, i)
			if strings.Contains(title, prefix) {
				return_value = prefix
				found_flag = true
				break
			}
		}

		if found_flag {
			break
		}
		
		for k := 1; k <= 10; k++ {
			prefix := fmt.Sprintf("%dB - %d - ", j, k)
			if strings.Contains(title, prefix) {
				return_value = prefix
				found_flag = true
				break
			}
		}
		
		if found_flag {
			break
		}
	}
	return return_value
}

func read_tags(filename string) {
	set_flag := false
	mp3_file, err := id3.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open %s: %s\n", filename, err)
		return
	} else {
		defer mp3_file.Close()
		title := mp3_file.Title()
		prefix := search_for_key_slash_pattern(title)
		if prefix != "" {
			fmt.Println("Prefix: ", prefix)
			new_title := strings.TrimPrefix(title, prefix)
			if set_flag {
				mp3_file.SetTitle(new_title)
				fmt.Println("New Title:", mp3_file.Title())
			} else {
				fmt.Println("*New Title:", new_title)
			}
		}
		
		title = mp3_file.Title()
		prefix1 := search_for_key_pattern(title)
		if prefix1 != "" {
			fmt.Println("Prefix: ", prefix1)
			new_title := strings.TrimPrefix(title, prefix1)
			if set_flag {
				mp3_file.SetTitle(new_title)
				fmt.Println("New Title:", mp3_file.Title())
			} else {
				fmt.Println("*New Title:", new_title)
			}
		}
		
		if prefix != "" || prefix1 != "" {
			fmt.Println("Path:        ", filename)
			fmt.Println("ID3 Version: ", mp3_file.Version())
			fmt.Println("Artist:      ", mp3_file.Artist())
			fmt.Println("Title:       ", mp3_file.Title())
			fmt.Println("Album:       ", mp3_file.Album())
			fmt.Println("Year:        ", mp3_file.Year())
			fmt.Println("Genre:       ", mp3_file.Genre())
			
			bpm_frame := mp3_file.Frame("TBPM")
			if bpm_frame != nil {
				bpm := bpm_frame.String()
				fmt.Println("BPM:         ", bpm)
			}
			
			key_frame := mp3_file.Frame("TKEY")
			if key_frame != nil {
				key := key_frame.String()
				fmt.Println("KEY:         ", key)
			}
			
			comments := mp3_file.Comments()
			for _, comment := range comments {
				if strings.HasPrefix(comment, "eng") {
					tokens := strings.Split(comment, ":")
					if len(tokens) == 2 {
						fmt.Println("Comment [eng]:", tokens[1])
					}
				}
			}
			
			fmt.Println()
		}
	}
}

func visit(path string, f os.FileInfo, err error) error {
	if !f.IsDir() {
		if strings.HasSuffix(f.Name(), ".mp3") {
			read_tags(path)
		}
	}
	return nil
} 

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s [path to mp3s]\n", os.Args[0])
		return
	}
	root_path := os.Args[1:][0]
	err := filepath.Walk(root_path, visit)
	fmt.Println(err)
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
