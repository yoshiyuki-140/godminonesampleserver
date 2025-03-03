.PHONY: run kill up

run:
	go build
	./godminonesampleserver &

kill:
	killall godminonesampleserver

up:
	docker compose up -d