package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	HelpFlag    bool
	VersionFlag bool
	cpOption    string
	XjreOption  string
	Class       string
	args        []string
}

func ParseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.HelpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.HelpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.VersionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.Class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func PrintUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
	fmt.Printf("\t or %s [-options] -jar jarfile [args]\n", os.Args[0])
}
func StartJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.Class, cmd.args)
}
func PrintVersion() {
	println("Version 0.0.1")
}
