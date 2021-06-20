build-migration:
	go build -o migration cmd/migrate/*.go

migration-create: build-migration
	./migration create $(NAME)

migration-up: build-migration
	./migration up

migration-down: build-migration
	./migration down

migration-redo: build-migration
	./migration redo

migration-status: build-migration
	./migration status

migration-version: build-migration
	./migration version

migration-fix: build-migration
	./migration fix

build: migration-up
	go build -o api cmd/api/main.go

run: build
	./api

# Dev
docker:
	@docker build -f ./Dockerfile.dev -t ranggarifqi/ahsan_muslim_name_generator_api_dev:latest .

run-docker: docker
	@docker-compose up

stop-docker:
	@docker-compose down

# Test
test:
	go test -v ./...

.PHONY: build run test