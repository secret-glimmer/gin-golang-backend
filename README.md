# Simple Gin Golang Backend

Simple golang backend service using [Gin](https://gin-gonic.com/) framework.

<img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" align="center" width="128" />

## Run Application

1.  Copy `.env.example` file and paste it within same directory
2.  Rename it as `.env`
3.  Run following command

```bash
swag init -g cmd/main.go
```

```bash
go run ./cmd
```

4.  Visit <http://localhost:8000/swagger/index.html> in your web browser to see swagger document
