# UPDATE:
The renewal price for hits.link was very expensive, so we decided to drop the domain. If you'd like to continue using hits please update your URLs to `https://hits-app.vercel.app/hits?url=https://yoururl.com/`


```diff
- https://hits.link/hits?url=https://yoururl.com/
+ https://hits-app.vercel.app/hits?url=https://yoururl.com/
```

Frontend: `https://hits-app.vercel.app`

API: `https://hits.hop.sh/`

# [hits](https://hits-app.vercel.app) - the better hits counter
![Hits](https://hits-app.vercel.app/hits?url=https://github.com/heybereket/hits&bgRight=292B2F)

## API Docs

**Create a new Hit:** <br />
`GET https://hits-app.vercel.app/hits?url=https://yoururl.com (20reqs/minute)`

**Query Params:**
> Valid hex color codes are accepted (remove "#" in the queries that apply) <br />
> Default border type is round

```bash
?json - Send JSON with data instead of SVG
?label - Set label text
?bgLeft - Set the background colour of the label
?bgRight - Set the background colour of the hits number
?color - Set the text color
?border - Choose a border (square or round)
```

**Get Top Hits:** <br />
`GET https://hits.hop.sh/v1/top (30s cache, 50reqs/minute)`

## Used on
**Open a PR to add your website!**
- https://github.com/looskie
- https://github.com/alii
- https://github.com/heybereket
- https://github.com/cnrad
- https://github.com/walkator


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
