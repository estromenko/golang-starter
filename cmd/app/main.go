package main

import (
	"flag"
	"golang-starter/internal/app"
	"log"
)

var (
	config string
)

func init() {
	flag.StringVar(&config, "config", "", "Path to config file")
}

func main() {
	flag.Parse()
	log.Fatal(app.Run(config))
}
