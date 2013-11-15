.PHONY: windows

all:
	go build app.go


windows:
	GOOS=windows GOARCH=386 go build -o map.exe app.go
