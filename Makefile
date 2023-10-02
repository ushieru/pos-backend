swag:
	swag init
container:
	swag init && podman build . -t point_of_sale
container-run:
	podman run -p 8080:8080 --name point_of_sale_container -d point_of_sale
clean-container:
	podman stop --all && podman rm point_of_sale_container && podman image rm localhost/point_of_sale