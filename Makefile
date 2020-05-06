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


docker-publish:
	docker build -t smartbestbuys-api-go:0.0.1 -t smartbestbuys-api-go:latest .
	docker tag smartbestbuys-api-go 273540820564.dkr.ecr.ap-southeast-1.amazonaws.com/smartbestbuys-api-go:0.0.1
	export AWS_PROFILE=smart-best-buys-dev
	$(DOCKER_TOKEN)
	docker push 273540820564.dkr.ecr.ap-southeast-1.amazonaws.com/smartbestbuys-api-go:0.0.1

DOCKER_TOKEN := $(shell aws ecr get-login --no-include-email --region ap-southeast-1)