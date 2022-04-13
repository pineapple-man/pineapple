package classpath

import (
	"os"
	"path/filepath"
)

// Classpath 存在三个字段，分别存放三种类路径
type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreoption, cpOption string) *Classpath {
	classpath := &Classpath{}
	classpath.parseBootAndExtClasspath(jreoption)
	classpath.parseUserClasspath(cpOption)
	return classpath
}
func (c *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := c.bootClasspath.ReadClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := c.extClasspath.ReadClass(className); err == nil {
		return data, entry, err
	}
	return c.userClasspath.ReadClass(className)

}

func (c *Classpath) String() string {
	return c.userClasspath.String()
}

func (c *Classpath) parseBootAndExtClasspath(jreoption string) {
	jreDir := getJreDir(jreoption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	c.bootClasspath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	c.extClasspath = newWildcardEntry(jreExtPath)
}

// 优先使用用户输入的 -Xjre 选项作为 jre 目录，
// 如果没有输入该选项，则在当前目录下寻找 jre 目录
// 如果找不到，尝试通过 JAVA_HOME 环境变量寻找，如果环境变量也找不到，则异常返回
func getJreDir(jreoption string) string {
	if jreoption != "" && exists(jreoption) {
		return jreoption
	}
	if exists("./jre") {
		return "./jre"
	}
	if javaHome := os.Getenv("JAVA_HOME"); javaHome != "" {
		return filepath.Join(javaHome, "jre")
	}
	panic("Can not find jre folder")
}

// 用于判断目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (c *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	c.userClasspath = newEntry(cpOption)
}
