dev:
	@echo "=============starting locally============="
	docker-compose -f resources/docker/docker-compose.yaml up --build

test:
	export GIN_MODE=test && richgo test -failfast ./app/... -coverprofile .coverage.txt
	go tool cover -func .coverage.txt

down:
	docker-compose -f resources/docker/docker-compose.yaml down

clean: down
	@echo "=============cleaning up============="
	docker system prune -f
	docker volume prune -f
	docker images prune -f

format:
	go fmt ./app/...

mockgen:
	mockgen -source=./app/repositories/$(module)/init.go -destination=./app/mocks/$(module)/repo.go
	mockgen -source=./app/use_cases/$(module)/init.go -destination=./app/mocks/$(module)/use_case.go

mockgen-all:
    mockgen module=user