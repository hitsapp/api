# [hits.link](https://hits.link) - the better hits counter
![Hits](https://hits.link/hits?url=https://github.com/heybereket/hits&bg=292B2F)

## API Docs

**Create a new Hit:** <br />
`GET https://api.hits.link/v1/hits?url=https://yoururl.com (15reqs/minute)`

**Query Params:**
> Valid hex color codes are accepted (remove "#" in the queries that apply)
> Default border type is round

```bash
?json - Send JSON with data instead of SVG
?label - Set label text
?bg - Set the background colour of the hits
?color - Set the text color
?border - Choose a border (square or round)
?font - Set the font family (supports any font from Google Fonts)
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
