package main

import "restapi/src/app"

func main() {
	app := &app.App{}
	app.Initialize()
	app.Start()
}
