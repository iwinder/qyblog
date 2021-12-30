root:
	export Root=gitee.com/windcoder/qingyublog


build:
	echo "Building go qingyublog binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64