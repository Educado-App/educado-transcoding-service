@echo off
SETLOCAL

FOR /F "tokens=* USEBACKQ" %%F IN (`git describe --tags %git rev-list --tags --max-count=1%`) DO (
    SET VERSION=%%F
)

FOR /F "tokens=* USEBACKQ" %%F IN (`git rev-list --count HEAD`) DO (
    SET BUILD=%%F
)

FOR /F "tokens=* USEBACKQ" %%F IN (`git rev-parse --short HEAD`) DO (
    SET GITHASH=%%F
)

IF NOT EXIST server\build (
    mkdir server\build
)

go build -o \build\educado.exe -ldflags "-X main.Version=%VERSION% -X main.Build=%BUILD% -X main.GitHash=%GITHASH%"

copy ..\..\.env server\build\
copy ..\..\gcp_credentials.json \build\



echo Built version %VERSION% (Build %BUILD%) and copied .env to server\build\

ENDLOCAL
