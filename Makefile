BINARY_NAME=gogetovpn

build:
	mkdir -p bin
	curl -o bin/openvpn-install.sh https://raw.githubusercontent.com/angristan/openvpn-install/master/openvpn-install.sh
	chmod +x bin/openvpn-install.sh
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME} cmd/gogetovpn/main.go

run:
	APP_ENV=production ./${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm -rf ./bin