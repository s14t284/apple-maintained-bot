build:
	go build -o bin/apple-maintained-bot cmd/server/main.go

deploy:
	make build
	git push heroku main