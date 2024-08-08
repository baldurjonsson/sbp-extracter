package main

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

const DATAFILE = "dataFile.txt"

var KEYS = map[int]string{
	-1: "-",
	0:  "A",
	1:  "A#",
	2:  "B",
	3:  "C",
	4:  "C#",
	5:  "D",
	6:  "D#",
	7:  "E",
	8:  "F",
	9:  "F#",
	10: "G",
	11: "G#",
	12: "F#m",
	13: "Gm",
	14: "G#m",
	15: "Am",
	16: "Bbm",
	17: "Bm",
	18: "Cm",
	19: "C#m",
	20: "Dm",
	21: "D#m",
	22: "Em",
	23: "Fm",
}

type DataFile struct {
	Songs []Song `json:"songs"`
}

type Song struct {
	Author   string `json:"author"`
	Name     string `json:"name"`
	Key      int    `json:"key"`
	KeyShift int    `json:"keyShift"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: sbp-extracter <sbp file>")
		os.Exit(1)
	}
	reader, err := zip.OpenReader(os.Args[1])

	if err != nil {
		panic(err)
	}
	defer reader.Close()
	for _, file := range reader.File {
		if file.Name == DATAFILE {
			fileReader, err := file.Open()
			if err != nil {
				panic(err)
			}
			fileBuffer := bufio.NewReader(fileReader)
			_, err = fileBuffer.ReadBytes('\n') // File version is in the first line, ignored.
			if err != nil {
				panic(err)
			}
			var dataFile DataFile
			err = json.NewDecoder(fileBuffer).Decode(&dataFile)
			if err != nil {
				panic(err)
			}

			for _, song := range dataFile.Songs {
				key := song.Key + song.KeyShift
				if (song.Key < 12 && key >= 12) || (song.Key >= 12 && key >= 24) {
					key -= 12
				}
				fmt.Printf("%s\t%s\t%s\n", song.Author, song.Name, KEYS[key])
			}

		}
	}
}
