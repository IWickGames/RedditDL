package net

import (
	"io"
	"net/http"
	"os"
	"os/exec"
)

/*
Calls FFMPEG to convert .m3u8 files to mp4 files
*/
func ConvertM3U8(url string, outputFile string) error {
	com := exec.Command("ffmpeg", "-i", url, outputFile)
	err := com.Run()
	return err
}

/*
Downloads an image from a url
*/
func DownloadImage(url string, saveLocation string) error {
	client := &http.Client{}                   // Create a new http client
	r, err := http.NewRequest("GET", url, nil) // Create a new GET request
	if err != nil {
		return err
	}

	// Set the user agent
	// Not required but a good idea
	r.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")

	resp, err := client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close() // Make sure to defer the close so it executes when the function exits

	file, err := os.Create(saveLocation) // Create the file
	if err != nil {
		return err
	}
	defer file.Close() // Closes it when the function returns

	_, err = io.Copy(file, resp.Body) // Copy the bytes from the url into the file
	if err != nil {
		return err
	}

	return nil
}
