# Build API with Go + Gin + Swagger

## swagger

```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

Setup gopath:

```sh
vi ~/.bashrc
export PATH=$PATH:$(go env GOPATH)/bin
source ~/.bashrc
```

Init swagger doc:
`swag init`

Import the swagger package in `router.go`:
```go
	_ "github.com/nathan/gopportunitiesbb/docs"
	swaggerFiles "github.com/swaggo/files"    		 // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" 		 // gin-swagger middleware
```

Add the API document string in `main.go`.

Swagger endpoint:
```sh
http://localhost:8080/swagger/index.html
```

## makefile

```sh
# generate swagger documents
make docs
# run server
make run
# build pgm gopportunitiesbb
make build
```