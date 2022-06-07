FROM golang:1.12-alpine AS build_base

RUN apk add --no-cache git
WORKDIR /tmp/goapp
COPY . .

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go build -o ./out/feedfish cmd/main.go


FROM alpine:3.9 
RUN apk add ca-certificates

COPY --from=build_base /tmp/goapp/out/feedfish /app/feedfish
COPY --from=build_base /tmp/goapp/.pipeline.yml /app/.pipeline.yml

WORKDIR /app
CMD ["/app/feedfish"]
