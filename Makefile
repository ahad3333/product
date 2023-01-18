run:
	go run cmd/main.go

swag-gen:
	swag init -g api/api.go -o api/docs

migration-up:
	migrate -path ./migrations/postgres/ -database 'postgres://postgres:0003@localhost:3003/catalog?sslmode=disable' up

