.PHONY: test

test:
	CGO_ENABLED=1 \
	go test -coverprofile=coverage.out ./... 
