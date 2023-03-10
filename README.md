# Go Echo Simple API with Clean Architecture

[![Build & Unit Test](https://github.com/zett-8/go-clean-echo/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/zett-8/go-clean-echo/actions/workflows/test.yml)

## 🤓 About this repo
This is a sample of Web API built by Go (Echo) according to *Clean architecture*.  
But I have to say that I'm not a backend specialist nor of Go, so Clean architecture here could be wrong or missing essential concepts. (I'm sorry in that case)

This sample consists of 4 layers, **"Models"**, **"Stores"**, **"Services(Logic)"** and **"Handlers(Framework)"**, although naming might differ in other samples.
Each layer only concerns/handles its inside layer, not the other way around.
This makes it super easy to replace each layer. (also good for testing)

![clean architecture](./utils/img.png)

#### What you might find helpful in this repo.
- [x] [echo](https://github.com/labstack/echo)(Framework) for handling requests
- [x] Clean architecture
- [x] zap([uber-go/zap](https://github.com/uber-go/zap)) for logging
- [x] swagger([swaggo/echo-swagger](https://github.com/swaggo/echo-swagger)) for documentation
- [x] sql-migrate([rubenv/sql-migrate](https://github.com/rubenv/sql-migrate)) for migration
- [x] Database transaction in service layer
- [x] sqlmock ([DATA-DOG/go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)) for testing
- [x] Multi-stage build for optimized docker image
- [x] Hot reload([cosmtrek/air](https://github.com/cosmtrek/air)) for efficient development
- [x] JWT authentication([auth0/go-jwt-middleware](https://github.com/auth0/go-jwt-middleware/)) with [Auth0](https://auth0.com/) for security
- [x] Automated testing on GitHub actions 

<br />

## 👟 How to run
Install.
```shell
git clone git@github.com:zett-8/go-clean-echo.git

cd go-clean-echo
```

Download dependencies.
```shell
go mod download
```

Fill out auth0 config to run with JWT authentication. Or simply disable JWT middleware
```go
// configs/auth0.go
var Auth0Config = Auth0ConfigType{
    Domain:             "****",
    ClientID:           "****",
    Audience:           []string{"****"},
    Issuer:             "****",
    SignatureAlgorithm: validator.RS256,
    CacheDuration:      15 * time.Minute,
}

// or 

// handlers/handlers.go
func SetApi(e *echo.Echo, h *Handlers, m echo.MiddlewareFunc) {
    g := e.Group("/api/v1")
    g.Use(m) // <- Comment out this line
}
```

Run docker.
```shell
docker-compose up
```

<br />

## 🌱 Tips 

### Use multi-stage build
Using multistage build reduces the size of the docker image.

```dockerfile
ARG PORT=8888

# Base image for local development. Use 'air' for hot reload.
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


# The image for build. Set CGO_ENABLE=0 not to build unnecessary binary.
FROM golang:1.18-alpine as builder

ARG PORT
ENV PORT=$PORT

WORKDIR /go/app/builder

COPY --from=base /go/app/base /go/app/builder

RUN CGO_ENABLED=0 go build -o main -ldflags "-s -w"


# The final image to run 
FROM gcr.io/distroless/static-debian11 as production

ARG PORT
ENV PORT=$PORT

WORKDIR /go/app/src

COPY --from=builder /go/app/builder/main /go/app/src/main

EXPOSE $PORT

CMD ["/go/app/src/main"]
```

### Use "Air" for hot reload
We used to have other reloaders such as [realize](https://github.com/oxequa/realize).
But it seems no longer be developed or maintained, so I recommend to use "air"

> [cosmtrek/air](https://github.com/cosmtrek/air).

### Mock a unit for testing
Here is one example to mock "service" for "handler" test.

Let's say this is the service to mock.
```go
// services/services.go
type Services struct {
    AuthorService
    BookService
}

func New(s *stores.Stores) *Services {
    return &Services{
        AuthorService: &authorService{s.AuthorStore},
        BookService:   &bookService{s.BookStore},
    }
}

// services/author.go
type (
    AuthorService interface {
        GetAuthors() ([]models.Author, error)
        DeleteAuthor(id int) error
    }

    authorService struct {
        store stores.AuthorStore
    }
)

func (s *AuthorServiceContext) GetAuthors() ([]models.Author, error) {
	r, err := s.store.Get()
	return r, err
}

func (s *AuthorServiceContext) DeleteAuthor(id int) error {
	err := s.store.DeleteById(id)
	return err
}
```

And in /handlers/author_test.go. Declare Mock struct.
```go
// handlers/author_test.go
type MockAuthorService struct {
	services.AuthorService
	MockGetAuthors       func() ([]models.Author, error)
	MockDeleteAuthorById func(id int) error
}

func (m *MockAuthorService) GetAuthors() ([]models.Author, error) {
	return m.MockGetAuthors()
}

func (m *MockAuthorService) DeleteAuthor(id int) error {
	return m.MockDeleteAuthorById(id)
}
```

Then use it as mockService.
```go
// handlers/author_test.go
func TestGetAuthorsSuccessCase(t *testing.T) {
    s := &MockAuthorService{
        MockGetAuthors: func () ([]models.Author, error) {
            var r []models.Author
            return r, nil
        },
    }
    
    mockService := &services.Services{AuthorService: s}
    
    e := Echo()
    h := New(mockService)
}
```

<br />

## 📃 API Document (Swagger)
```text
http://localhost:8888/swagger/index.html
```

<br />

## ✅ Testing
```shell
make test
```

<br />

## 💜 References
[bxcodec/go-clean-arch](https://github.com/bxcodec/go-clean-arch)  
[onakrainikoff/echo-rest-api](https://github.com/onakrainikoff/echo-rest-api)  
[satishbabariya/go-echo-auth0-middleware](https://github.com/satishbabariya/go-echo-auth0-middleware)  
[xesina/golang-echo-realworld-example-app](https://github.com/xesina/golang-echo-realworld-example-app)  

