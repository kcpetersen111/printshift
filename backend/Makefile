build:
	docker compose up -d --build

reset:
	docker rm -f $(docker ps -aq)

.PHONY: clean
clean: reset build

