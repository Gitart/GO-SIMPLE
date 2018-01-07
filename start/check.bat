rem Control enviroument

rem Path to current Programm API Service
SET GOPATH=D:\DB\service
SET GOBIN=D:\DB\service\bin

rem путь к компилятору
SET GOROOT=C:\GO
SET PATH=%GOROOT%\BIN;%PATH%;

go version
go env
go list
pause
