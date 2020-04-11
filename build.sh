mkdir out

# Mac 64bit
GOOS=darwin GOARCH=amd64  go build -o out/suddenlyMuscleTraining main.go

# Windows 64bit
GOOS=windows GOARCH=amd64 go build -o out/suddenlyMuscleTraining.exe main.go
