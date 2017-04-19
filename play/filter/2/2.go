package main

import (
	"fmt"
	"io/ioutil"

	"github.com/pravj/geopattern"
)

// Prints pattern's SVG string for a specific pattern
func main() {
	//args := map[string]string{"generator": "squares"}
	args := map[string]string{"phrase": "M", "color": "#f9b"}
	gp := geopattern.Generate(args)
	fmt.Println(gp)
	ioutil.WriteFile("./o2.svg", []byte(gp), 0777)
}
