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
	docker compose -v -f $(DC) up --build prometheus -d

.PHONY: down_all
down_all:
	docker compose -v -f $(DC) down -v

.PHONY: pandora
pandora:
	pandora ./requests/pandora_config/flat_ramp_up.yaml

.PHONY: dump_logs
dump_logs:
	curl http://localhost:8081/dump

.PHONY: gen_ammo
gen_ammo:
	go run ./cmd/generate_data/main.go

.PHONY: active_py_venv
active_py_venv:
	source myenv/bin/activate

.PHONY: deactivate_py_venv
deactivate_py_venv:
	deactivate

.PHONY: graph
graph:
	. myenv/bin/activate && \
	python3 ./analize/analize.py && \
	deactivate


