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

## Day four

Context with timeout is used to limit ipinfo fetching time with an interval of 1 second. Timeout error is handeled correctly and dislayed to a user.


## Day five

Learned how to deal with the [Cobra](https://github.com/spf13/cobra). I'm using positional and keyword arguments with hierarchical command logic and validation:

```
> dailygo 
Dailygo — a golang learning project.

Usage:
  dailygo [command]

Available Commands:
  completion  Generate the autocompletion script
  help        Help about any command
  info        Print geoinfo of an IP address
  initdb      Init a new database
  runserver   Run the dailygo web server

Flags:
  -h, --help   help for dailygo

Use "dailygo [command] --help" for more info
```

