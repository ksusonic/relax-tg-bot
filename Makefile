.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/main.go

run: build
	./.bin/bot

build-image:
	docker build -t relax-tg:v0.1 .

start-container:
	docker run --name relax-tg --env-file .env relax-tg:v0.1
