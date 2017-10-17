@echo off
SETLOCAL
:: start

rem Path to current Programm API Service
SET GOPATH=%CD%
SET GOBIN =%CD%\bin

rem ïóòü ê êîìïèëÿòîðó
SET GOROOT=C:\GO
SET PATH=%GOROOT%\BIN;%PATH%;
cls

title Run "GO" 
color 0f

rem echo "  "                 
rem echo ....................................................................
rem echo gopath = %gopath%
rem echo ....................................................................
REM go clean -r -i
REM go install -a
rem go get -u github.com/dancannon/gorethink
rem go get -u "github.com/dancannon/gorethink"
rem go env
rem d:\MON\RETHINKDB\service\service.go

go build -o context.exe 
echo "Compilition.."
context.exe 

rem d:\Curl\curl.exe -X POST http://10.10.10.10:5555/admin/9bd9ead764bd89a88347f1f3a85b6/2
rem go rem hoservice.go
rem go run hoservice.go
rem go build hoservice.go
rem hoservice.exe
rem go build -o read_csv_file.exe  read_csv_file.go
echo %Date% %Time%  %date:~-10,2% Run >> log.txt
echo %Date% %Time%  %date:~-10,2% Run >> log.txt
rem go run selectgo.go
rem go run d:\RETHINKDB\GO\main.go
@pause
