# UPDATE #2:
Bad news, I have archived this project due to time and money running it. _**Good news**_, AprilNEA has made an alternative that's a drop-in replacement to hits! If you would like to support and use their's, visit https://github.com/AprilNEA/hits.

```diff
- https://hits-app.vercel.app/hits?url=https://yoururl.com/
+ https://hits.aprilnea.com/hits?url=https://yoururl.com/
```

> [!NOTE]
> I do not own nor have control over AprilNEA's websites, domains, or repositories!!

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


# Sponsored By
Hits is brought to you by [Hop](https://hop.io/) - No more configs. No more fuss. Just push your code.


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
`GET https://hits.hop.sh/v1/leaderboard (30s cache, 50reqs/minute)`

## Used on
**Open a PR to add your website!**
- https://github.com/looskie
- https://github.com/alii
- https://github.com/heybereket
- https://github.com/cnrad
- https://github.com/walkator


## Running Locally
```bash
# Run Docker Compose 
$ docker-compose up -d

# Install packages
$ go get

# Start API
$ go run main.go

```

## Frontend
The frontend has been moved to a separate repository:
[https://github.com/hitsapp/web](https://github.com/hitsapp/web)

## Maintainers
- @heybereket
- @Looskie
