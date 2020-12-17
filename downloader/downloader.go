package downloader

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	fileName    string
	fullURLFile string
)

// DownloadFile writes the content of a GET request to a new file
func DownloadFile(URL string) (string, error) {

	// Build fileName from fullPath
	fileURL, err := url.Parse(URL)
	if err != nil {
		return "", err
	}

	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName = segments[len(segments)-1]

	err = os.Chdir("photos")
	if err != nil {
		fmt.Printf("Error switching to photos directory.\n%v\n", err)
	}

	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	// GET the image
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request. %v", err)
		return "", err
	}

	// Put content on file
	resp, err = client.Get(URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)
	defer file.Close()

	log.Printf("Downloaded image %s with size %d.", fileName, size)

	// Return to the root dir
	err = os.Chdir("..")
	return fileName, err
}
