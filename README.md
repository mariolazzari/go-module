# Building a module in Go

## Project setup

```sh
mkdir toolkit
mkdir app
cd toolkit
go mod init github.com/mariolazzari/go-module/toolkit
cd ../app
go mod init myapp
# init go workspace
cd ..
go work init toolkit app
```

```go
package toolkit

type Tools struct{}
```

## Getting started with the module

###
