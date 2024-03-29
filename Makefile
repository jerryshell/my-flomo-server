app_name = my-flomo-server

amd64-darwin :
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o $(app_name)-amd64-darwin
	tar -czvf $(app_name)-amd64-darwin.tar.gz $(app_name)-amd64-darwin

amd64-linux :
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $(app_name)-amd64-linux
	tar -czvf $(app_name)-amd64-linux.tar.gz $(app_name)-amd64-linux

amd64-windows :
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o $(app_name)-amd64-windows.exe
	tar -czvf $(app_name)-amd64-windows.tar.gz $(app_name)-amd64-windows.exe

arm-linux :
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o $(app_name)-arm-linux
	tar -czvf $(app_name)-arm-linux.tar.gz $(app_name)-arm-linux

all : amd64-darwin amd64-linux amd64-windows arm-linux

clean :
	rm -f $(app_name)-*

dep :
	go get -u && go mod tidy
