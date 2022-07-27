FROM golang:1.16-buster as uwu
ARG DATABASE_URL
ENV DATABASE_URL=${DATABASE_URL}

RUN echo ${DATABASE_URL}

RUN mkdir /app
WORKDIR /app
COPY . .

RUN go run github.com/prisma/prisma-client-go db push
RUN go build

FROM debian:latest as RUNNER

COPY --from=uwu ./app/api .

RUN apt-get update
RUN apt install -y ca-certificates && rm -rf /var/cache/apk/*

CMD ["./api"]