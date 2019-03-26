package pkg

import (
	"fmt"
	"github.com/sanjid133/gopher-love/pkg/platform"
	"github.com/sanjid133/gopher-love/pkg/platform/github"
	"strings"
)

type Org struct {
	platform string
	name     string
	token    string
}

func ValidateOrganization(token, org string) (Org, error) {
	orgs := strings.Split(org, "/")
	if len(orgs) != 2 {
		return Org{}, fmt.Errorf("Org name %v is not valid. Correct format is github.com/<org>")
	}
	return Org{
		platform: orgs[0],
		name:     orgs[1],
		token:    token,
	}, nil
}

func (o *Org) LoveOrganization() error {
	var ml platform.MakeLove
	var err error
	switch o.platform {
	case github.Platform:
		if ml, err = github.Initialize(o.token); err != nil {
			return err
		}

	}
	return ml.LoveOrganization(o.name)
}
