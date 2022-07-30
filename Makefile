test:
	docker-compose exec echo-server go test ./...

swagger:
	swag init