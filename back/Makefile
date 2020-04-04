run: gen
	go run -tags dev .

gen:
	go generate

static: gen
	env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -tags netgo -installsuffix netgo -ldflags '-w -extldflags "-static"' -o hyback .

windows: gen
	env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o hyback.exe main.go

user:
	http POST :8888/users/ username="root" name="Administrator" password="toor"
