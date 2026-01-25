FROM golang:1.25-alpine AS build

WORKDIR /app
COPY . .

RUN go build -ldflags="-w -s"

FROM alpine:latest AS run

WORKDIR /app

COPY --from=build /app/curls.sh ./curls.sh

EXPOSE 3000

CMD ["./curls.sh"]
