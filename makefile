build:
	env GOOS=linux GOARCH=arm GOARM=6 go build -o test_rs232 test_build.go
	env GOOS=linux GOARCH=arm GOARM=6 go build -o rs232-controller main.go