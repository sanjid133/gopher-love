package dep

import (
	"testing"
	"fmt"
	"os"
)

func TestParse(t *testing.T) {
	fmt.Println(os.Getwd())
	fmt.Println(GetMyPath())
	//err := Parse("/home/sanjid/go/src/github.com/sanjid133/gopher-love/Gopkg.toml")
	//fmt.Println(err)
}


