package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

// Apod struct uses the "hdurl" attribute from
// json unmarshalled below through this struct
type Apod struct {
	Hdurl string
}

var (
	// Change the apiKey to your NASA API generated key
	apiKey = "DEMO_KEY"
	url    = "https://api.nasa.gov/planetary/apod?api_key=" + apiKey
)

func main() {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	// Grab full json response from APOD API
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	// Create the struct which only has the Hdurl (json:"hdurl") attribute
	var imgData Apod
	json.Unmarshal([]byte(b), &imgData)

	// Execute the wallpaper command found in same directory...
	// OR you can put the wallpaper.exe file into some directory
	// and set PATH to include that directory
	cmd := exec.Command("wallpaper", "-i", imgData.Hdurl)
	cmd.Stdout = os.Stdout // pipe command output to os output
	cmd.Stderr = os.Stderr // pipe command error output to os error output
	err = cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "command failed to run: %v\n", err)
	}
}
