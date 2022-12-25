# What is this?

This is a learning project I'm willing to grow up adding new feature every day each with some well-know library, a framework or a concept. My aim is to drive a train.

## How to run this program

1. Run the project: `PORT=1600 go run cmd/web/main.go`

2. Open this url in your browser: [http:/127.0.0.1:1600/](http:/127.0.0.1:1600/)

<hr>


## Day one

Created a web application with [Gin](https://github.com/gin-gonic/gin) using it's routers and templates. JSON unmarshalling and HTTP requests also has been covered.

## Day two

Created a Domain/IP database with [gorm](https://gorm.io). Domains are resolved on-the-fly and are stored into the database. Gin's HTTP args and POST params are covered too.

## Day three

Configured logging with <a href="https://github.com/uber-go/zap">zap</a>. Tuned up logging to&nbsp;both file and stdout, with Caller and Stacktrace options.
