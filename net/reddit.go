package net

import (
	"strings"
)

/*
Searches for media links in a block of html supplied
	Reddits media urls are
		- preview.redd.it
		- i.redd.it
		- v.redd.it
*/
func GetMediaUrls(data string) []string {
	var media []string
	var cache string

	testLinks := strings.Split(data, "src=\"https://") // Get a list of items beginning with a media link
	for _, value := range testLinks {                  // Loop through all links
		if strings.HasPrefix(value, "preview.redd.it") { // Check if the link starts with preview.redd.it
			cache = strings.Split(value, "\"")[0] // Cut out the url from the src=""

			/*
				This checks to make sure the preview link is for media
				Reddit uses preview.redd.it to also store stuff such as badge images

				Bage urls look like this "preview.redd.it/award_images/something.png"
				While media urls look like this "https://preview.redd.it/something.jpg"

				When you split it by "/" it should return a length of two only if it is a media url
			*/
			if len(strings.Split(cache, "/")) != 2 {
				continue
			}

			cache = strings.ReplaceAll(cache, "amp;", "") // Cleans up encoded characters in the url
			media = append(media, "https://"+cache)
		}

		if strings.HasPrefix(value, "i.redd.it") {
			cache = strings.Split(value, "\"")[0]
			media = append(media, "https://"+cache)
		}

		if strings.HasPrefix(value, "v.redd.it") {
			cache = strings.Split(value, "?")[0]
			media = append(media, "https://"+cache)
		}
	}

	return media
}

/*
Supplied a url will return a string containing the type of media the url leads to
Possible values are IMAGE, VIDEO, UNKNOWN
*/
func GetMediaType(url string) string {
	if strings.HasPrefix(url, "https://i.redd.it") {
		return "IMAGE"
	}
	if strings.HasPrefix(url, "https://preview.redd.it") {
		return "IMAGE"
	}
	if strings.HasPrefix(url, "https://v.redd.it") {
		return "VIDEO"
	}
	return "UNKNOWN"
}
