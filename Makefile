deps:
	go mod download
	go mod verify
	go mod tidy

# go install golang.org/x/lint/golint@latest
lint:
	gofmt -w=true -s=true -l=true ./
	golint ./...
	go vet ./...

check: lint

test:
	go test -v ./...

cover:
	go test --cover ./...

# https://go.dev/blog/govulncheck
# install it by `go install golang.org/x/vuln/cmd/govulncheck@latest`
vuln:
	which govulncheck
	govulncheck -version
	govulncheck ./...
