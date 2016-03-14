# Password Hashing
A simple password hash application by Go, the hash can be use for WordPress

## Resouces
- [gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) for router
- [phpass](http://www.openwall.com/phpass/) for the main idea
- [Golang's port of PHP Portable Hash](https://gist.github.com/georgerb/f0ef84cf487e019e32f6)

## Installation

`go get github.com/tatthien/go-hash`

## Usage

```
cd $GOPATH\src\tatthien\go-hash
go run main.go
```

After running the app, a web server will be started at port 8080 (you can edit port number in code). Then visit [http://localhost:8080/hash/your-password](#) to see the result.

![HashPassword](http://www.tatthien.com/wp-content/uploads/2016/03/hash.gif)


