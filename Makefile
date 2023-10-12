swag:
	swag init
exe-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o point_of_sale
exe-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o point_of_sale.exe
container:
	podman build . -t docker.io/ushieru/total-pos:latest
container-run:
	podman run -p 8080:8080 --name total-pos_container -d total-pos
clean-container:
	podman stop --all && podman rm total-pos_container && podman image rm total-pos