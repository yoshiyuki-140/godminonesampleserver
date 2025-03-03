.PHONY: run kill up get post psqllogin

run:
	go build
	./godminonesampleserver

kill:
	killall godminonesampleserver

up:
	docker compose up -d

get:
	curl localhost:8080

post:
	curl -X POST \
	localhost:8080/tasks \
	-H "Content-Type: application/json" \
	-d '{"task":"Task","isCompleted":false}' | jq

psqllogin:
	psql postgres postgres -h localhost -p 5432 # password -> `password`