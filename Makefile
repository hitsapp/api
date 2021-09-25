api_dev:
	docker compose up -d
	cd api/prisma
	go run github.com/prisma/prisma-client-go db push
