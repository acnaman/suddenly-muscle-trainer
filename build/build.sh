mkdir -p out/win_64
mkdir -p out/mac_64
mkdir -p out/linux_64

# Mac 64bit
GOOS=darwin GOARCH=amd64  go build -o out/mac_64/suddenlyMuscleTraining main.go

# Windows 64bit
GOOS=windows GOARCH=amd64 go build -o out/win_64/suddenlyMuscleTraining.exe main.go

# Linux 64bit
GOOS=linux GOARCH=amd64 go build -o out/linux_64/suddenlyMuscleTraining main.go
