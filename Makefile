build:
	cd ./src && \
	GOOS=linux GOARCH=amd64 go build -o ../bin/linux-amd64 main.go; \
	GOOS=linux GOARCH=arm64 go build -o ../bin/linux-arm64 main.go; \
	GOOS=windows GOARCH=amd64 go build -o ../bin/win-amd64.exe main.go; \
	GOOS=windows GOARCH=arm64 go build -o ../bin/win-arm64.exe main.go; \
	GOOS=darwin GOARCH=arm64 go build -o ../bin/darwin-arm64 main.go; \
	GOOS=darwin GOARCH=amd64 go build -o ../bin/darwin-amd64 main.go

run:
	cd ./src && go run main.go

run-visualizer:
	cd ./src/plot && python3 plot.py