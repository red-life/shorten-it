generate_mocks:
	mockery --config .mockery.yaml

run_tests:
	go test ./...