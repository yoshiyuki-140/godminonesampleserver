.PHONY: run kill up get post psqllogin

run:
	go build
	./godminonesampleserver

kill:
	killall godminonesampleserver

up:
	docker compose up -d

psqllogin:
	psql postgres postgres -h localhost -p 5432 # password -> `password`


register:
	curl -X POST \
	localhost:8080/account/register \
	-H "Content-Type: application/json" \
	-d '{"name":"kuro","password":"pass"}'

login:
	curl -X GET \
	localhost:8080/account/login \
	-H "Content-Type: application/json" \
	-d '{"name":"kuro","password":"pass"}'

logout:
	curl -X POST \
	localhost:8080/account/logout \
	-H "Content-Type: application/json" \
	-d '{"session_id":"999999"}'


getAllTasks:
	curl -X GET \
	localhost:8080/tasks \
	-H "Content-Type: application/json" | jq

getTask:
	curl -X GET \
	localhost:8080/tasks/1 \
	-H "Content-Type: application/json" | jq

createTask:
	curl -X POST \
	localhost:8080/tasks \
	-H "Content-Type: application/json" \
	-d '{ "session_id": "999999","task": { "task": "kuro", "is_completed": false } }' | jq

updateTask:
	curl -X PUT \
	localhost:8080/tasks/7 \
	-H "Content-Type: application/json" \
	-d '{ "session_id": "999999","task": { "task": "kuro", "is_completed": true } }' | jq


