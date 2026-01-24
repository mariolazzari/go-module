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

### Create random string

```go
package toolkit

import "crypto/rand"

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

// Tools is the type used to instantiate this module.
// All the variables of this type will have access to the methods of this module.
type Tools struct{}

// RandomString generates a random string of the given length n, using randomStringSource as the source characters.
func (t *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(randomStringSource)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)
}
```

### Test module in Go app

```go
package toolkit

import "testing"

func TestTool_RandomString(t *testing.T) {
	var testTools Tools
	s := testTools.RandomString(10)
	if len(s) != 10 {
		t.Errorf("RandomString() expected length of 10, got %d", len(s))
	}
}
```

```sh
go test .
```

## Uploading files

### UploadFiles method
