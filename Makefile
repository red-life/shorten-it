generate_mocks:
	mockery --config .mockery.yaml

run_tests:
	go test ./...

build:
	docker-compose build

run:
	make build
	docker-compose up -d