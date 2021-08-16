.PHONY: compile
compile:
	CGO_ENABLED=0 GOOS=linux go build -a -o ./svc ./cmd/main.go