build: cmd/*
	CGO_ENABLED=0 go build -o dist/xdo ./cmd/main.go

develop:
	docker-compose -f develop/docker-compose.yaml up -d
	go run cmd/main.go
.PHONY: develop

test:
	./test/test-post.sh localhost:8000
.PHONY: test

loadtest:
	./test/load-test.sh
.PHONY: loadtest
