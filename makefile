air = ~/go/bin/air

dev:
	${air} --build.cmd "go build -o bin/func ./cmd/main.go" --build.bin "./bin/func"