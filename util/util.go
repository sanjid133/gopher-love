package util

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"os/user"
	"strings"
)

func GetPlatform(name string) string {
	parts := strings.Split(name, ".")
	return parts[0]
}

func EnsureDirectory(directory string) error {
	return os.MkdirAll(directory, 0777)
}

func HomeDirectory() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func CheckGoDirectory() {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	fmt.Println(gopath)
}
