build:
	docker build -t k3d-myregistry.localhost:12345/act1:latest .
	docker push k3d-myregistry.localhost:12345/act1:latest