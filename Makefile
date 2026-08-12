export majorVersion=1
export minorVersion=0
export patchVersion=$(shell git log --format='%h' | wc -l)
export ver=$(majorVersion).$(minorVersion).$(patchVersion)

tag:
	git tag "v$(ver)"

deps:
	go mod download
	go mod verify
	go mod tidy

tools:
	which go
	which golint # go install golang.org/x/lint/golint@latest
	which govulncheck # go install honnef.co/go/tools/cmd/staticcheck@latest
	
fix:
	go fix ./...

lint:
	gofmt -w=true -s=true -l=true ./
	golint ./...
	go vet ./...
	staticcheck -checks=inherit,-SA1019 ./...

check: lint

test:
	go test -v ./...

cover:
	go test --cover ./...

# https://go.dev/blog/govulncheck
# install it by `go install golang.org/x/vuln/cmd/govulncheck@latest`
vuln:
	which govulncheck
	govulncheck ./...

start: run

run:
	go run example/example.go

include make/*.mk
