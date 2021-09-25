api_dev:
	docker compose up -d
	cd api/prisma && go run github.com/prisma/prisma-client-go db push
	cd api && go run main.go

deps:
	cd api && go get -d ./
	cd web && yarn

web_dev:
	cd web && yarn dev

build:
	cd api && go build
	cd web && go build
