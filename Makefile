.PHONY: windows, test, clean, run

all: run


run:
	go run app.go

windows:
	GOOS=windows GOARCH=386 go build -o map.exe app.go


test:
	go test ./draw


clean:
	rm -fr ./tiles
