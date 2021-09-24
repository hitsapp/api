# hits

## Run dev environment
- Clone the repo
- Type `yarn` or `npm i` 
- Type `docker-compose up -d`
- Remove `-example` out of all the `.env-example` files.
- Type `cd api/prisma/ && go run github.com/prisma/prisma-client-go db push`
- Type `cd api`
- Type `go run main.go`
- You're off to the races!

## Stack
```bash
Frontend:
- Next.js
- TypeScript

Backend:
- Go
- Fiber
- Prisma
```

## Maintainers
- @heybereket
- @Looskie
