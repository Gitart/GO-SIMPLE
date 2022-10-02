@setlocal enableextensions enabledelayedexpansion
@ECHO off
ECHO.
ECHO :::::::::::::::::::::::::: arg.bat example :::::::::::::::::::::::::::::::
ECHO :: By:  U, 2022-07-29                                                   ::
ECHO :: Version: 1.0                                                         ::
ECHO :: Purpose: Checks the args passed to the batch.                        ::
ECHO ::                                                                      ::
ECHO :: Start by gathering all the args with the %%* in a for loop.          ::
ECHO ::                                                                      ::
ECHO :: Now we use a 'for' loop to search for our keys which are identified  ::
ECHO :: by the text '--'. The function then sets the --arg ^= to the next    ::
ECHO :: arg. "CALL:Function_GetValue" ^<search for --^> ^<each arg^>         ::
ECHO ::                                                                      ::
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

ECHO.

ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
ECHO :: From the command line you could pass... arg.bat --x 90 --y 220       ::
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
ECHO.
ECHO.Checking Args:"%*"

FOR %%a IN (%*) do (
    CALL:Function_GetValue "--","%%a" 
)

ECHO.
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
ECHO :: Now lets check which args were set to variables...                   ::
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
ECHO.
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
ECHO :: For this we are using the CALL:Function_Show_Defined "--x,--y,--z"   ::
ECHO ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
ECHO.
CALL:Function_Show_Defined "--x,--y,--z"
endlocal
goto done

:Function_GetValue

REM First we use find string to locate and search for the text.
echo.%~2 | findstr /C:"%~1" 1>nul

REM Next we check the errorlevel return to see if it contains a key or a value
REM and set the appropriate action.

if not errorlevel 1 (
  SET KEY=%~2
) ELSE (
  SET VALUE=%~2
)
IF DEFINED VALUE (
    SET %KEY%=%~2
    ECHO.
    ECHO ::::::::::::::::::::::::: %~0 ::::::::::::::::::::::::::::::
    ECHO :: The KEY:'%KEY%' is now set to the VALUE:'%VALUE%'                     ::
    ECHO :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
    ECHO.
    ECHO %KEY%=%~2
    ECHO.
    REM It's important to clear the definitions for the key and value in order to
    REM search for the next key value set.
    SET KEY=
    SET VALUE=
)
GOTO:EOF

:Function_Show_Defined 
ECHO.
ECHO ::::::::::::::::::: %~0 ::::::::::::::::::::::::::::::::
ECHO :: Checks which args were defined i.e. %~2
ECHO :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
ECHO.
SET ARGS=%~1
for %%s in (%ARGS%) DO (
    ECHO.
    ECHO :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
    ECHO :: For the ARG: '%%s'                         
    IF DEFINED %%s (
        ECHO :: Defined as: '%%s=!%%s!'                                             
    ) else (
        ECHO :: Not Defined '%%s' and thus has no value.
    )
    ECHO :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
    ECHO.
)
goto:EOF

:done
