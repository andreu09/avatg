include .env
export

.PHONY: fmt generate generate-go test update-deps lint  build run

fmt:
	go fmt ./...

generate-go:
	go generate ./...

wire:
	wire ./internal/app/di

test:
	go test -race -v -count=1  ./...

lint:
	golangci-lint run --disable-all -E errcheck

update-deps:
	go mod tidy
	go mod verify

build: fmt

build-for-docker: build
	docker rmi ${PROJECTNAME} -f
	docker build -t ${PROJECTNAME} ./build

build:
	${BUILD_CMD} -o ${BUILD_DIR}${PROJECTNAME} cmd/main.go

run: build
	mkdir -p ${BUILD_DIR}${LOG_DIR}
	cd ${BUILD_DIR}; ./${PROJECTNAME}

run-for-docker: build-for-docker
	docker run --rm ${PROJECTNAME}:latest
