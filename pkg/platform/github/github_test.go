package github

import (
	"fmt"
	"os"
	"testing"
)

func TestInitialize(t *testing.T) {
	l, err := Initialize(os.Getenv("GITHUB_TOKEN"))
	fmt.Println(err)
	l.LoveOrganization("kubernetes")
}
