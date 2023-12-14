# The Weather Guess


## Golang

Install Golang

### Air

you should really get air installed. If you run into an error, you can always cut and rerun the server, but that is kind of a pain
```bash
go install github.com/cosmtrek/air@latest
```

### Go LSP

make sure you have golsp installed in your favorite editor.
```bash
go install golang.org/x/tools/gopls@latest
```

### Setting up the server

go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware
touch cmd/main.go

SIDE NOTE

if you see the following error, execute `go mod tidy`  after importing echo

Principle: HATEOAS

Hypermedia As The Engine Of Application State
Does that mean HTML is finally a programming language?
Does that mean I am an HTML Engineer?


Its existed for a long time HATEOAS Circa 2011

# To run

start server with

$ air
cjkgg
