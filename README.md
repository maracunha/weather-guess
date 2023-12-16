# The Weather Guess


## Golang

Install Golang

### ☁️  Air - Live reload for Go apps
```bash
go install github.com/cosmtrek/air@latest
```

### Go LSP

Make sure you have golsp installed in your favorite editor.
```bash
go install golang.org/x/tools/gopls@latest
```

### Setting up the server

```bash
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware
touch cmd/main.go
```

SIDE NOTE

if you see the following error, execute `go mod tidy`  after importing echo

# To run

Start server with just:
```bash
air
```
