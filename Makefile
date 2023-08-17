build:
	@go build -o ./bin/main

run: build
	@./bin/main

test:
	go test -v ./...

install:
	cd frontend&&npm install&&npm run build

build-front:
	cd frontend&&npm run build