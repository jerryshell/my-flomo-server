app_name = my-flomo-server

build-amd64-darwin :
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o $(app_name)-amd64-darwin

build-amd64-linux :
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $(app_name)-amd64-linux

build-amd64-windows :
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o $(app_name)-amd64-windows

build-arm-linux :
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o $(app_name)-arm-linux

build-all : build-amd64-darwin build-amd64-linux build-amd64-windows build-arm-linux

clean :
	rm -f $(app_name)-*
