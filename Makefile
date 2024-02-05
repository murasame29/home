build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm cmd/app/main.go 
	go build -o app cmd/app/main.go 

run: build
	./app