compose = docker compose
exec_container = $(compose) exec -T app

up:
	$(compose) kill
	$(compose) build
	$(compose) up -d

stop:
	$(compose) stop

down:
	$(compose) down -v

generate-key:
	chmod +x ./generate_key.sh
	./generate_key.sh

logs:
	$(compose) logs -f

exec:
	$(exec_container) $(cmd)

test:
	$(exec_container) go test -v ./...

ssh:
	$(exec_container) /bin/sh

test-coverage:
	$(exec_container) go test -v -coverprofile=tmp/coverage.out ./...
	$(exec_container) go tool cover -html=tmp/coverage.out -o tmp/coverage.html
