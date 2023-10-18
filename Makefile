swag:
	swag init -d api -o api/swagger
exe-all:
	make exe-linux && make exe-windows && make exe-silicon
exe-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bins/pos-linux-amd64
exe-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bins/pos-windows-amd64.exe
exe-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bins/pos-darwin-amd64.app
exe-silicon:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bins/pos-silicon-arm64.app
container:
	podman build . -t docker.io/ushieru/total-pos:latest
container-run:
	podman run -p 8080:8080 --name total-pos_container -d total-pos
clean-container:
	podman stop --all && podman rm total-pos_container && podman image rm total-pos