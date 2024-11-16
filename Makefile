run:
	docker compose up -d --build

reset:
	@echo "Removing all Docker containers..."
	@docker rm -f $$(docker ps -aq) || true

backend-logs:
	@docker compose logs backend

.PHONY: clean
clean: reset build

