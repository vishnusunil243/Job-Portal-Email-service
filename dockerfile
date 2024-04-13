FROM golang:1.21.5-bullseye AS build

RUN apt-get update

WORKDIR /app

COPY . .

RUN go mod download

WORKDIR /app/cmd

RUN go build -o email-service

FROM busybox:latest

WORKDIR /email/cmd

COPY --from=build /app/cmd/email-service .

COPY --from=build /app/.env /email

EXPOSE 8087

CMD ["./email-service"]