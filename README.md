# [hits.link](https://hits.link) - the better hits counter

## API Docs

**Create a new Hit:** <br />
`GET https://api.hits.link/v1/hits?url=https://yoururl.com (15reqs/minute)`

**Query Params:**
```bash
# Enter colours without "#"
?bg - Set the background colour of the hits
?json - Send JSON with data instead of SVG
?color - Set the text color
```

**Get Top Hits:** <br />
`GET https://api.hits.link/v1/top (30s cache, 50reqs/minute)`


## Running Locally
```bash
# Copy all .env.example files into your .env file

# Run Docker Compose 
$ docker-compose up -d

# Sync Database
$ cd api/prisma && go run github.com/prisma/prisma-client-go db push

# Start Go API
$ cd api && go run main.go

# Install dependencies
$ cd web && yarn

# Run web
$ cd web && yarn dev
```

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
