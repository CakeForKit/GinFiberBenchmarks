DC := ./deployment/docker-compose.yml
# DB_ENV := ./configs/db_config.env


.PHONY: run_gin_app
run_gin_app:
# --no-cache
	docker compose -v -f $(DC) up --build gin-app -d

.PHONY: down_gin_app
down_gin_app:
	docker compose -v -f $(DC) down -v gin-app

.PHONY: run_prom
run_prom:
# --no-cache
# 	docker compose -v -f $(DC) build app_craftplace
	docker compose -v -f $(DC) up --build prometheus -d

.PHONY: gen_ammo
gen_ammo:
	go run ./cmd/generate_data/main.go
# 	pandora ./requests/pandora_config/flat_config.yaml


