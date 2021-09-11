package net

import (
	"io/ioutil"
	"net/http"
)

/*
	Function for downloading the html from the webpage
	Returns the http status, the html code and an error status that will be set to nil if everything is ok
*/
func Get(url string) (status string, data string, err error) {
	client := &http.Client{}                   // Define a new http client
	r, err := http.NewRequest("GET", url, nil) // Create a new GET request
	if err != nil {
		return "", "", err
	}

	// Set the user agent
	// This may not be required but is good to keep
	r.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")

	// Execute the GET request and download the website
	resp, err := client.Do(r)
	if err != nil {
		return "", "", err
	}

	// Decodes ReadCloser into a bytes array
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	return resp.Status, string(bytes), nil
}
