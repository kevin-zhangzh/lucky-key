lucky-linux:
	GOOS=linux GOARCH=amd64 go build -o ./build/lucky ./cmd
lucky:
	go build -o ./build/lucky ./cmd