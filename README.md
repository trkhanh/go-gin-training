go mod init newsfeeder
go get -u github.com/gin-gonic/gin


go test $path
go test ./...
go test -cover ./...