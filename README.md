# apod-wallpaper
set NASA's astronomy picture of the day as your desktop background

## Dependencies
Uses the wallpaper.exe file built from source on 6/30/2020 from github.com/hirenchauhan2/wallpaper

## Setup
This works on Windows 10.

Check out github.com/hirenchauhan2/wallpaper for the source on the ```wallpaper``` command; I had to build that from source using hirenchauhan2's instructions on the README. I dropped the executable in the same directory as the apod.go source file.

Don't just keep using the DEMO_KEY, it has a 30 req/IP/hr limit and a 50 req/IP/day limit as well.

Generate your own NASA API key at https://api.nasa.gov. The rate limit after doing so is 1000 req/IP/hr and you can check out the other cool APIs there too.

In ```apod.go```, change the apiKey and you're all set.

    var (
        // Change the apiKey to your NASA API generated key
        apiKey = "DEMO_KEY"
        url    = "https://api.nasa.gov/planetary/apod?api_key=" + apiKey
    )


Build and run and check your wallpaper. Should be different every day.

    > cd apod-wallpaper
    > go build apod.go
    > apod

Try using the command in Windows Task Manager to automate it every day, so that you have a different wallpaper each time.