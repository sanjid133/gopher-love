install:
	go install ./...

test: install
	go test -v

fmt:
	gofmt -s -w pkg cmd util

tags:
	find ./ -name '*.go' -print0 | xargs -0 gotags > TAGS

push:
	git push origin HEAD

revendor:
	dep ensure -v


