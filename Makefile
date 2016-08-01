all:
	go install go_env/...

run: all
	./bin/server -c config/server.json

test:
	go test -v go_env/...
