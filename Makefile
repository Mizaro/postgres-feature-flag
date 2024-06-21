
up:
	docker compose -f ./deploy/docker-compose.yaml -f ./deploy/docker-compose.override.yml up -d
down:
	docker compose -f ./deploy/docker-compose.yaml -f ./deploy/docker-compose.override.yml down --remove-orphans -v

restart: down up
