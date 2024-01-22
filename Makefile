release:
	GOOS=linux GOARCH=386   go build -v -o build/chucknorris_linux_i386 .
	GOOS=linux GOARCH=amd64 go build -v -o build/chucknorros_linux_amd64 .
	GOOS=linux GOARCH=arm   go build -v -o build/chucknorris_linux_armhf .
	GOOS=linux GOARCH=arm64 go build -v -o build/chucknorris_linux_aarch64 .

clean:
	-rm -rf build

.PHONY: c_api clean c_api_android c_api_linux