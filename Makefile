build:
	GOOS=linux GOARCH=amd64 go build -o dsa-tools-amd-linux .
	GOOS=linux GOARCH=arm64 go build -o dsa-tools-arm-linux .

start:
	PASSWORT=password go run .

wasm:
	GOOS=js GOARCH=wasm go build -o ./resource/static/wasm/time.wasm ./wasm/time.go