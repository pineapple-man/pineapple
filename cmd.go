package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"pineapple/classfile"
	"pineapple/classpath"
)

type Cmd struct {
	BasicConfig
	HeapConfig
	StackConfig
	cpOption   string
	XjreOption string
	Class      string
	args       []string
}
type StackConfig struct {
	// 虚拟机栈大小
	Xss string
}
type HeapConfig struct {
	// 设置堆的大小
	Xms string
	// 设置最大堆容量
	Xmx string
}
type BasicConfig struct {
	// 帮助标记
	HelpFlag bool
	// 版本
	VersionFlag bool
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
	flag.StringVar(&cmd.Xms, "Xms", "", "heap size")
	flag.StringVar(&cmd.Xmx, "Xmx", "", "max heap size")
	flag.StringVar(&cmd.Xss, "Xss", "", "max stack size")
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
	classPath := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v \n", classPath, cmd.Class, cmd.args)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	cf := loadClass(className, classPath)
	fmt.Println(cmd.Class)
	printClassInfo(cf)
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf(" %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf(" %s\n", m.Name())
	}
}

func loadClass(className string, classPath *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := classPath.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}
func PrintVersion() {
	println("Version 0.0.1")
}
