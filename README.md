<div align="center" id="top"> 
  <img src="./.github/app.gif" alt="Go Fiber Starter" />
</div>

<h1 align="center">Go Fiber Starter</h1>

<p align="center">
  <img alt="Github top language" src="https://img.shields.io/github/languages/top/genesysflow/go-fiber-starter?color=56BEB8">

  <img alt="Github language count" src="https://img.shields.io/github/languages/count/genesysflow/go-fiber-starter?color=56BEB8">

  <img alt="Repository size" src="https://img.shields.io/github/repo-size/genesysflow/go-fiber-starter?color=56BEB8">

  <img alt="License" src="https://img.shields.io/github/license/genesysflow/go-fiber-starter?color=56BEB8">

  <img alt="Github issues" src="https://img.shields.io/github/issues/genesysflow/go-fiber-starter?color=56BEB8" />

  <img alt="Github forks" src="https://img.shields.io/github/forks/genesysflow/go-fiber-starter?color=56BEB8" />

  <img alt="Github stars" src="https://img.shields.io/github/stars/genesysflow/go-fiber-starter?color=56BEB8" />
</p>


<p align="center">
  <a href="#dart-about">About</a> &#xa0; | &#xa0; 
  <a href="#sparkles-features">Features</a> &#xa0; | &#xa0;
  <a href="#rocket-technologies">Technologies</a> &#xa0; | &#xa0;
  <a href="#white_check_mark-requirements">Requirements</a> &#xa0; | &#xa0;
  <a href="#checkered_flag-starting">Starting</a> &#xa0; | &#xa0;
  <a href="#memo-license">License</a> &#xa0; | &#xa0;
  <a href="https://github.com/bangadam" target="_blank">Author</a>
</p>

<br>

## :dart: About

Simple and scalable starter kit to build powerful and organized REST projects with Fiber.

## :sparkles: Features

- [x] Logging
- [x] Repository Pattern
- [x] ORM database with Gorm
- [x] Redis cache with go-redis
- [x] Mocking with GoMock
- [x] Api documentation with Swaggo
- [x] Support JWT authentication
- [x] Containerization with Docker compose
- [x] Unit testing with testify
- [x] CI\CD with Github Actions

## :rocket: Technologies

The following tools were used in this project:

- [Go](https://go.dev)
- [PostgreSQL](https://www.postgresql.org)
- [Docker](https://www.docker.com/)
- [Fiber](https://github.com/gofiber/fiber)
- [Gorm](https://gorm.io)
- [Fx](https://github.com/uber-go/fx)
- [Zerolog](https://github.com/rs/zerolog)
- [GoMock](https://github.com/golang/mock)

## :white_check_mark: Requirements

Before starting :checkered_flag:, you need to have [Git](https://git-scm.com), [Go](https://go.dev), [Docker](https://www.docker.com/) and [PostgreSQL](https://www.postgresql.org) installed.

## :checkered_flag: Starting

```bash
# Clone this project
$ git clone https://github.com/genesysflow/go-fiber-starter

# Access
$ cd go-fiber-starter

# Download dependencies
$ go get -v ./...

# Start vite 
$ cd frontend && yarn dev

# Run the project
$ go run cmd/example/main.go

# The server will initialize in the <http://{host}:{port}>
```

## :memo: License

This project is under license from MIT. For more details, see the [LICENSE](LICENSE) file.

Forked from by <a href="https://github.com/bangadam/go-fiber-starter" target="_blank">bangadam/go-fiber-starter</a>

&#xa0;

<a href="#top">Back to top</a>
