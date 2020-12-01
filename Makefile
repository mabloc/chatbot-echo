build:
	chmod +x ./run.sh
	go mod tidy && go mod verify && go build -o ./bin/bot-echo main.go