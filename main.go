package main

import (
	"fmt"
	"os"
	"redditdl/net"
	"strings"
)

func main() {
	/*
		Check to make sure all the arguments are entered
		If not display an error

		Arguments
		0: Default golang arguments (the executables location)
		1: First user argument (the reddit url)
		2: Second user argument (the location of the saved media)
	*/
	if len(os.Args) != 3 {
		fmt.Println("ERROR: Invalid arguments")
		fmt.Println("Usage:")
		fmt.Println("  redditdl <redditURL> <fileName (no file extenchion)>")
		os.Exit(1)
	}

	// Make sure the url is a reddit url
	if !strings.HasPrefix(os.Args[1], "https://www.reddit.com/r") {
		fmt.Println("ERROR: Invalid reddit URL")
		fmt.Println("  EX: https://www.reddit.com/r/funny/comments/pm4u4a/i_photoshop_animals_into_random_objects_heres_my/")
		os.Exit(1)
	}

	// Download the page
	fmt.Println("[Reddit] Downloading page")
	_, data, err := net.Get(os.Args[1])
	if err != nil {
		fmt.Println("ERROR: Failed to request url")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Parce the media links out of the page
	fmt.Println("[Reddit] Parcing media urls")
	links := net.GetMediaUrls(data)
	if len(links) == 0 {
		fmt.Println("ERROR: Could not locate any media urls")
		os.Exit(1)
	}

	// Reddit-DL automaticly uses the first link as the main media url
	// This is selected and displayed here
	fmt.Println("[Parsor] Using:", links[0])

	/*
		Find the media type of the media url
		Possible types are: IMAGE, VIDEO

		UNKNOWN is returned if a type cannot be found
	*/
	switch mediaType := net.GetMediaType(links[0]); mediaType {
	case "IMAGE":
		fmt.Println("[Download] Media type: IMAGE")
		fmt.Println("[Download] Starting download of media (this may take some time)")
		// Download the image as a jpg
		err := net.DownloadImage(links[0], os.Args[2]+".jpg")
		if err != nil {
			fmt.Println("ERROR: Failed to download image")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println("Successfully downloaded", os.Args[2]+".jpg")
		os.Exit(0)

	case "VIDEO":
		fmt.Println("[Download] Media type: VIDEO")
		fmt.Println("[Download] Starting download of media (this may take some time)")
		/*
			Reddit uses .m3u8 files to store videos
			Reddit-DL uses FFMPEG to convert to a .mp4 file
		*/
		err := net.ConvertM3U8(links[0], os.Args[2]+".mp4")
		if err != nil {
			fmt.Println("ERROR: Failed to call FFMPEG (is it installed?)")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println("Successfully downloaded", os.Args[2]+".mp4")
		os.Exit(0)

	// Unknown type was detected so display an error
	case "UNKNOWN":
		fmt.Println("ERROR: Media type could not be identified")
		os.Exit(1)
	}
}
