GOOS=linux  GOARCH=amd64  go build -o pfop_linux_amd64   pfop.go
GOOS=linux  GOARCH=386    go build -o pfop_linux_386     pfop.go
GOOS=darwin GOARCH=amd64  go build -o pfop_darwin_amd64  pfop.go
GOOS=darwin GOARCH=386    go build -o pfop_darwin_386    pfop.go
GOOS=windows GOARCH=amd64 go build -o pfop_windows_amd64.exe pfop.go
GOOS=windows GOARCH=386   go build -o pfop_windows_386.exe   pfop.go