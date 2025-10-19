DC := ./deployment/docker-compose.yml
# DB_ENV := ./configs/db_config.env


.PHONY: run_gin_app
run_gin_app:
# --no-cache
	docker compose -f $(DC) up --build gin-app -d

.PHONY: down_gin_app
down_gin_app:
	docker compose -f $(DC) down -v gin-app

.PHONY: run_prom
run_prom:
	docker compose -f $(DC) up --build prometheus -d 

.PHONY: down_prom
down_prom:
	docker compose -f $(DC) down -v prometheus

.PHONY: run_cad
run_cad:
	docker compose -f $(DC) up --build cadvisortest -d

.PHONY: down_cad
down_cad:
	docker compose -f $(DC) down -v cadvisortest

.PHONY: stop_all
stop_all:
	docker compose -f $(DC) stop

.PHONY: pandora
pandora:
	pandora ./requests/pandora_config/flat_ramp_up.yaml

.PHONY: dump_logs
dump_logs:
	curl http://localhost:8080/dump

.PHONY: gen_ammo
gen_ammo:
	go run ./cmd/generate_data/main.go

.PHONY: graph
graph:
	. myenv/bin/activate && \
	python3 ./analize/analize.py && \
	deactivate


