FROM golang:1.18-alpine as base

ENV PORT=8888
ENV GO_ENV=development

WORKDIR /app/go/base

COPY go.mod .
COPY go.sum .

RUN apk add build-base
RUN go mod download
RUN go install github.com/cosmtrek/air@latest

COPY . .

FROM golang:1.18-alpine as builder

ENV PORT=8888

WORKDIR /app/go/builder

COPY --from=base /app/go/base /app/go/builder

RUN CGO_ENABLED=0 go build main.go

FROM alpine as production

WORKDIR /app/go/src

RUN apk add --no-cache ca-certificates
COPY --from=builder /app/go/builder/main /app/go/src/main

EXPOSE 8888

CMD ["/app/go/src/main"]
