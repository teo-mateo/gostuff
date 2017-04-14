package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const chunkSize = 4 * 1024
const url = "https://focusmusic.fm/api/tracks.php?offset={ofs}&timestamp={ts}&channel=classical"
const trackCount = 20

//Track ...
type Track struct {
	Title     string `json:"title"`
	Artist    string `json:"artist"`
	URL       string `json:"url"`
	Permalink string `json:"permalink"`
}

func main() {

	flagDownload := flag.Bool("download", false, "true if you want to download")
	flagFile := flag.String("json", "", "preloaded json; default: tracks.json")
	flag.Parse()

	var jsons []string
	if *flagFile != "" {
		jsons = fetchListFromFile(*flagFile)
	} else {
		jsons = fetchListFromWeb()
	}

	tracks := make([]Track, 0)

	for _, j := range jsons {
		fmt.Println(j)
		var t Track
		err := json.Unmarshal([]byte(j), &t)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(t)
		tracks = append(tracks, t)
	}

	if *flagDownload {
		for _, t := range tracks {
			download(&t)
		}
	}

}

func download(track *Track) {

	//MAKE HTTP CALL
	response, err := http.Get(track.URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer response.Body.Close()

	//SETUP DOWNLOADER
	downloader := bufio.NewReader(response.Body)
	chunk := make([]byte, chunkSize)

	totalBytes := response.ContentLength
	bytesSoFar := 0

	//SETUP FILE SAVE
	outputFile := fmt.Sprintf("%s - %s.mp3", track.Artist, track.Title)
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	writer := bufio.NewWriter(f)
	defer f.Close()

	//GET CHUNKS UNTIL DEATH
	for {
		n, err := downloader.Read(chunk)
		if err != nil {
			if io.EOF == err {
				break
			}
			log.Fatal(err.Error())
		}

		//AND WRITE TO FILE
		writer.Write(chunk)

		bytesSoFar += n

		pc := int(bytesSoFar) * 100 / int(totalBytes)
		fmt.Printf("\rread %d bytes of %d (%d%%)", bytesSoFar, totalBytes, pc)
	}
}

func fetchListFromFile(file string) (jsons []string) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err.Error())
	}
	return strings.Split(string(bytes), "\n")
}

func fetchListFromWeb() (jsons []string) {
	tracks := make([]string, 1)

	for track := 1; track <= trackCount; track++ {
		json := getNextTrack(&tracks, track)
		fmt.Println(json)
		time.Sleep(2 * time.Second)
	}
	return tracks
}

func getNextTrack(tracks *[]string, track int) (json string) {
	ts := time.Now().Unix()
	u := strings.Replace(url, "{ofs}", strconv.Itoa(track), 1)
	u = strings.Replace(u, "{ts}", strconv.Itoa(int(ts)), 1)

	response, err := http.Get(u)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		content := strings.Replace(string(body), `\/`, `/`, -1)
		*tracks = append(*tracks, content)
		return content
	}

	return "-"

}
