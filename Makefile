CCARMV7=arm-linux-gnueabihf-gcc
CCARM64=aarch64-linux-gnu-gcc

all: mod

mod:
	@echo "Running go mod tidy"
	go mod tidy

386: 
	@echo "Building 386"
	GO111MODULE=on GOOS=linux GOARCH=386 go build -o bin/pulse-oximeter.linux.386 main.go

amd64:
	@echo "Building amd64"
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/pulse-oximeter.linux.amd64 main.go

arm5:
	@echo "Building arm5"
	GO111MODULE=on GOOS=linux GOARCH=arm GOARM=5 go build -o bin/pulse-oximeter.linux.armv5 main.go

arm6:
	@echo "Building arm6"
	GO111MODULE=on GOOS=linux GOARCH=arm GOARM=6 go build -o bin/pulse-oximeter.linux.armv6 main.go

arm7:
	@echo "Building arm7"
	GO111MODULE=on GOOS=linux GOARCH=arm GOARM=7 go build -o bin/pulse-oximeter.linux.armv7 main.go

arm64:
	@echo "Building arm64"
	GO111MODULE=on GOOS=linux GOARCH=arm64 go build -o bin/pulse-oximeter.linux.arm64 main.go

darwin:
	@echo "Building darwin"
	GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -trimpath -o bin/pulse-oximeter.darwin.amd64 -a -tags netgo -ldflags '-s -w' main.go

windows:
	@echo "Building windows"
	GO111MODULE=on GOOS=windows GOARCH=amd64 go build -o bin/pulse-oximeter.windows.amd64.exe -a -tags netgo -ldflags '-w' main.go

build:
	@echo "Building pulse-oximeter"
	mkdir -p bin

install:
	cp bin/pulse-oximeter.linux.arm64 /usr/local/bin/pulse-oximeter
	chmod 755 /usr/local/bin/pulse-oximeter

systemd:
	cp deply/pulse-oximeter.service /etc/systemd/system/pulse-oximeter.service
	chmod 644 /etc/systemd/system/pulse-oximeter.service
	systemctl daemon-reload 
	systemctl enable pulse-oximeter.service

package: pulse-oximeter
	@echo "Packaging"

release: package
	@echo "Release"
