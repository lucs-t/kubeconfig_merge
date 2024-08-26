mac_intel:
	GOOS=darwin CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o merge_conf main.go
	chmod +x merge_conf
mac_apple:
	GOOS=darwin CGO_ENABLED=0 GOARCH=arm64 go build -ldflags="-s -w" -o merge_conf main.go
	chmod +x merge_conf
linux:
	GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o merge_conf main.go
	chmod +x merge_conf
