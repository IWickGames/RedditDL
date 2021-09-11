# Reddit Downloader (Reddit-dl)
Reddit-dl a tool to download images and video from reddit

---

## Usage
Reddit-DL is made to be easy to use
> redditdl [redditurl] [filename]

### Example
> redditdl https://www.reddit.com/r/funny/comments/pm9oip/respect_the_pup/ RespectThePup

This will download the video into a file named `RespectThePup.mp4`
Reddit-DL will automaticly convert to mp4 using FFMPEG and images will always be saved in the jepg format

---

## Requirements
Reddit-dl only requires that ffmpeg is installed and aviable by running `ffmpeg` once this is complete you can download videos fine. Images **do not** need ffmpeg to download, they will work nativly.

## About
This project is just kinda for fun and if there is a demand I will work on keeping it updated but if you want to help in cleaning up the code (yes I am new to GoLang and this is just me learning more about it) or improving it then feel free to make a pull request. I will try to keep checking for new requests.

## Building
To build this project you require [GoLang](https://golang.org/dl/). After you install it you install it in this root directory with `go.mod` and `main.go` execute this command "`go build .`" It will then output a file for your operating system that you can run.