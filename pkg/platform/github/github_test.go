package github

import (
	"testing"
	"fmt"
	"os"
)

func TestInitialize(t *testing.T) {
	l,err := Initialize(os.Getenv("GITHUB_TOKEN"))
	fmt.Println(err)
	l.LoveOrganization("kubernetes")
}
