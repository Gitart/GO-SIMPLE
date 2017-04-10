rem go run main.go
rem go  build -o main.exe
rem go tool pprof -gif main profile.ini>cpu.svg
rem go test -bench=.
rem go run main.go

rem go tool pprof -pdf main.exe
rem go tool pprof http://localhost:8080/debug/

rem go test -bench=BenchmarkRegex -cpuprofile cpu.out

rem go tool pprof -svg  cpu.out > cpu.svg

REM Path to current Programm API Service
SET GOPATH=%CD%
SET GOBIN=%CD%\bin


rem UNIX
rem Êîìïèëÿöèÿ ïîä UNIX ïëàòôîðìó
rem SET GOOS=linux
rem SET GOARCH=amd64
rem SET CGO_ENABLED=0 


REM ïóòü ê êîìïèëÿòîðó
SET GOROOT=C:\GO
SET PATH=%GOROOT%\BIN;%PATH%;
CLS


go build -o test.exe
test.exe

pause
