run:
	docker compose up -d --build

reset:
	@echo "Removing all Docker containers..."
	@docker rm -f $$(docker ps -aq) || true

exec-db:
	docker compose exec -it postgres psql -U user -p 5432 -h localhost printshift

backend-logs:
	@docker compose logs backend

.PHONY: clean
clean: reset build

