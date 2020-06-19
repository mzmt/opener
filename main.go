package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	flag.Parse()
	name := flag.Arg(0)

	app := make(map[string]string, 10)
	app["twitter"] = "https://twitter.com/home"
	app["gmail"] = "https://mail.google.com/"
	app["trans"] = "https://translate.google.co.jp/"
	app["github"] = "http://github.com/"
	app["reddit"] = "https://www.reddit.com/"
	app["facebook"] = "https://www.facebook.com/"
	app["amazon"] = "https://www.amazon.co.jp/"
	app["netflix"] = "https://www.netflix.com/jp/"

	if name == "list" {
		for key := range app {
			fmt.Println(key)
		}
		os.Exit(0)
	}

	if app[name] == "" {
		panic("App not found.")
	}

	// only for mac OS
	err := exec.Command("open", app[name]).Start()
	if err != nil {
		panic(err)
	}
}

func alias(name string) {
	// alias translate
	// e.g.    l -> list
}
