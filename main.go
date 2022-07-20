package main

import "github.com/davidenq/tweets-timeline-challenge/app"

func main() {
	app := &app.App{}
	app.
		LoadConfig().
		LoadServices().
		LoadDomain().
		LoadAPI().
		Init()
}
