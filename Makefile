.PHONY: windows

all:
	go build -o myapp app.go


windows:
	GOOS=windows GOARCH=386 go build -o map.exe app.go
