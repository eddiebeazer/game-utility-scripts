$Env:GOOS = "windows"; $Env:GOARCH = "amd64"

cd .\update-project-version
go build -o ..\bin\update-project-version.exe
cd ..

cd .\create-multiplayer-config
go build -o ..\bin\create-multiplayer-config.exe
cd..

$Env:GOOS = "linux"; $Env:GOARCH = "amd64"

cd .\p4-workspace-cleanup
go build -o ..\bin\p4-workspace-cleanup
cd ..

cd .\clean-playfab-multiplayer-builds
go build -o ..\bin\clean-playfab-multiplayer-builds

cd ..