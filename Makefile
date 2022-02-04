root:
	export Root=gitee.com/windcoder/qingyublog


build:
	echo "Building go qingyublog binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64

build-sys:
	echo "Building go qingyublog binary"
	mkdir -p bin/qycms
	mkdir -p bin/qycms/config
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build   -o  bin/qycms/qysystem cmd/qycms-system/qysystem.go
	cp configs/qysystem.yaml bin/qycms/config/qysystem.yaml