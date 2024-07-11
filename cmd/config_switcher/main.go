package main

import (
	"flag"
	"fmt"
	"os"
)

type cmdArgs struct {
	addFlag              bool
	removeFlag           bool
	setDefaultConfigFlag string
	configPath           string
	configsFolder        string
}

func main() {
	args := cmdArgs{}
	flag.BoolVar(&args.addFlag, "add", false, "Used when you want to add a config file to the program, requires the -config and -folder flags")
	flag.BoolVar(&args.removeFlag, "remove", false, "Used when you want to remove a config file to the program, requires the -config and -folder flags")
	flag.StringVar(&args.configPath, "config", "", "Used to specify the path and name of a configuration file")
	flag.StringVar(&args.configsFolder, "folder", "", "Used to specify the folder where the different config files are located for a specific target config")
	flag.StringVar(&args.setDefaultConfigFlag, "default", "", "Used to specify the default configuration that will be switched")

	flag.Parse()
	nonParsedArgs := flag.Args()

	// Validate the flags
	if args.addFlag && args.removeFlag {
		fmt.Println("Flags -add and -remove cannot be used at the same time")
		os.Exit(2)
	}
	if args.addFlag && (args.configPath == "" || args.configsFolder == "") {
		fmt.Println("Flag -add requires flags -config and -folder")
		os.Exit(2)
	}
	if args.removeFlag && (args.configPath == "" || args.configsFolder == "") {
		fmt.Println("Flag -remove requires flags -config and -folder")
		os.Exit(2)
	}

	if args.setDefaultConfigFlag != "" && (args.addFlag || args.removeFlag || (args.configPath != "") || (args.configsFolder != "")) {
		fmt.Println("Cannot use the -default flag in combination with any other flags")
		os.Exit(2)
	}

	// Get the non-parsed arguments
	if (args.addFlag || args.removeFlag) && (len(nonParsedArgs) > 0) {
		panic("")
	}

	fmt.Println("Hello, World!")
}
