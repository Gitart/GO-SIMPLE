@ECHO OFF
SETLOCAL
SET GOPATH=%CD%
CLS
TITLE Test "GO" - Run "%1"
COLOR 70
ECHO ................................................................
ECHO
ECHO
go run %1
ECHO -----------------------------------------------------------------
@PAUSE
