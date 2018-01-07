@ECHO OFF
ECHO Office
ECHO API SERVICE
ECHO Main Service
ECHO 07-09.2015 

SETLOCAL
:: start

REM Path to current Programm API Service
SET GOPATH=%CD%
SET GOBIN=%CD%\bin

rem UNIX
rem For Unix
rem SET GOOS=linux
rem SET GOARCH=amd64
rem SET CGO_ENABLED=0 


REM Ini
SET GOROOT=C:\GO
SET PATH=%GOROOT%\BIN;%PATH%;
CLS

TITLE Run "GO" 
COLOR 0f

REM ECHO "  "                 
REM ECHO ....................................................................
REM ECHO gopath = %gopath%
REM ECHO ....................................................................
REM go clean -r -i
REM go install -a
 
REM Îáíîâëåíèå äðàéâåðà Gorethinkdb
rem go get -u github.com/dancannon/gorethink
rem go get -u github.com/googollee/go-socket.io

REM go env
REM ECHO Update Gorethinkdb......................
REM d:\MORION\RETHINKDB\service\hoservice.go
REM 
REM SET versgo=go version
REM SET BUILD_DATE=

REM ECHO %BUILD_DATE%
REM ECHO %versgo%

REM go   help build
rem  pscp -P 2222 -l user1 -pw Password  test.txt  user1@10.0.50.16:/home/user1/1/ 

REM Build process
go build -o D:\SERVICE\service.exe 
CD D:\SERVICE\
service.exe

REM Log
REM service.exe  >> D:\DB\service\log.txt
REM d:\Curl\curl.exe -X POST http://10.0.0.24:5555/admin/adf9bd9ead764bd89a88347f1f3a85b6/2
REM
REM go rem service.go

REM go build -o read_csv_file.exe  read_csv_file.go
REM d:\db\go\bin\log.txt

rem update to server
rem updatetoserver.bat

ECHO %Date% %Time%  %date:~-10,0% Ñåðâåð çàïóùåí >> d:\SERVICE\log.txt 


REM go run selectgo.go
REM go run d:\DB\GO\main.go
@PAUSE
