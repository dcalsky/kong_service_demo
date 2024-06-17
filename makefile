test:
	echo "print env:"
	env
	go test -ldflags="-s=false" -gcflags=all=-l -v ./...  -coverprofile=coverage.out.tmp
	cat coverage.out.tmp | grep -v "_mock.go" > coverage.out
	go tool cover -html=coverage.out -o coverage.html


build:
	bash ./build.sh

run: build
	./output/boot.sh