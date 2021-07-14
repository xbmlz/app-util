mkdir .\bin\x86\
go build -ldflags="-w -s" -buildmode=c-shared -o .\bin\x86\AppUtil.dll main.go