ARG PORT=8888

FROM golang:1.18-alpine as base

ARG PORT
ENV PORT=$PORT
ENV GO_ENV=development

WORKDIR /go/app/base

COPY go.mod .
COPY go.sum .

RUN apk add build-base
RUN go mod download
RUN go install github.com/cosmtrek/air@latest

COPY . .

FROM golang:1.18-alpine as builder

ARG PORT
ENV PORT=$PORT

WORKDIR /go/app/builder

COPY --from=base /go/app/base /go/app/builder

RUN CGO_ENABLED=0 go build -o main -ldflags "-s -w"

FROM gcr.io/distroless/static-debian11 as production

ARG PORT
ENV PORT=$PORT

WORKDIR /go/app/src

COPY --from=builder /go/app/builder/main /go/app/src/main

EXPOSE $PORT

CMD ["/go/app/src/main"]
