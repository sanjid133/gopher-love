package util

import "strings"

type Repo struct {
	Platform string
	Owner string
	Name string
}


func ParseRepoUrl(url string) Repo {
	parts := strings.Split(url, "/")
	r := Repo{}
	if len(parts) >0 {
		r.Platform = parts[0]
	}
	if len(parts)>1 {
		r.Owner = parts[1]
	}
	if len(parts)>2 {
		r.Name = parts[2]
	}
	return r
}
