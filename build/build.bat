ren Mac 64bit
GOOS=darwin GOARCH=amd64 go build -o out/mac_64/suddenlyMuscleTraining ./runtime
GOOS=darwin GOARCH=amd64 go build -o out/mac_64/MTSetting ./gui

rem Windows 64bit
CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o out/win_64/suddenlyMuscleTraining.exe ./runtime
CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o out/win_64/MTSetting.exe ./gui

rem Linux 64bit
rem CC=gcc CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o out/linux_64/suddenlyMuscleTraining

