setlocal
set GOOS=windows
set CGO_ENABLED=0
go build -ldflags "-s" -v -o typedgen.exe
