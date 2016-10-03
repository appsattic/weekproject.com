package main

import (
	"fmt"
	"log"
	"os"

	temple "../"
)

func line() {
	fmt.Println("-------------------------------------------------------------------------------")
}

func main() {
	log.Println("Started")
	defer log.Println("Finished")

	line()
	tmpl, err := temple.NewTemple("templates", "base.html", false)
	log.Printf("tmpl=%#v, %#v\n", tmpl, err)
	line()

	trials := []string{"index.html", "about.html", "page.html", "post.html", "index.html"}

	for _, name := range trials {
		this, err := tmpl.Get(name)
		fmt.Printf("--- Template ---\n\n")
		fmt.Printf("this: %s, %#v, %#v\n\n", name, this, err)
		fmt.Printf("--- Output ---\n\n")
		this.Execute(os.Stdout, nil)
		line()
	}

	// and one that doesn't exist
	that, err := tmpl.Get("blah.html")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("that: %#v, %#v\n", that, err)
	that.Execute(os.Stdout, nil)
	line()
}
