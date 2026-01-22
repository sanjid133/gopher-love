package pkg

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/sanjid133/gopher-love/pkg/system"
	"github.com/sanjid133/gopher-love/util"
)

func FollowUser(url string) error {
	parts := strings.Split(url, "/")
	if len(parts) != 2 {
		return fmt.Errorf("user name %v is not valid. Correct format is github.com/<user>", url)
	}
	ctx := context.Background()

	org := util.GetPlatform(parts[0])
	platform, err := GetPlatform(org, ctx)
	if err != nil {
		return err
	}
	follow, err := platform.Initialize(system.Config)
	if err != nil {
		return err
	}

	isFollowed, err := follow.IsFollowed(parts[1])
	if err != nil {
		return err
	}
	if !isFollowed {
		if err := follow.SendFollow(parts[1]); err != nil {
			return err
		}
		log.Println(parts[1], ":", "Following Done...")
	} else {
		log.Println(parts[1], ":", "Already Followed...")
	}
	return nil
}
