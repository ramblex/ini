// Provides a simple command to parse an ini file and output it in a more
// parsable format.
package main

import (
	"flag"
	"fmt"

	"bitbucket.org/innovaeng/ini"
)

var iniPath string
var outputType string

func init() {
	flag.StringVar(&iniPath, "ini", "/opt/innova/innova.conf", "Path to ini file")
	flag.StringVar(&outputType, "output-type", "string", "Output type")
}

func main() {
	flag.Parse()
	ini, err := ini.ReadIni(iniPath)

	if err != nil {
		fmt.Print(err)
	}

	if outputType == "string" {
		fmt.Print(ini)
	}
}
