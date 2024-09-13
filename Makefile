MAKEFLAGS += --silent

export ENV_FILE = .env
ifneq ("$(wildcard $(ENV_FILE))","")
	include $(ENV_FILE)
	export $(shell sed 's/=.*//' $(ENV_FILE))
endif

DATABASE_URL = postgresql://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)@127.0.0.1:$(POSTGRES_EXTERNAL_PORT)/$(POSTGRES_DATABASE)?sslmode=disable
MIGRATION_TOOL = migrate/migrate:v4.12.2
MIGRATION_DIR = $(PWD)/migrations

.PHONY: *

# Serving
start:
	docker-compose up -d
restart: clean start
stop:
	docker-compose down
clean:
	docker-compose down --rmi local -v
app-rebuild:
	docker-compose up -d --no-deps --build app

# Migrations
create-migration:
	docker run -v $(MIGRATION_DIR):/migrations -w /migrations -u $(shell id -u) $(MIGRATION_TOOL) create -ext sql $(name)
migrate-up:
	docker run -v $(MIGRATION_DIR):/migrations --network host $(MIGRATION_TOOL) -path=/migrations/ -database '$(DATABASE_URL)' up
migrate-down:
	docker run -v $(MIGRATION_DIR):/migrations --network host $(MIGRATION_TOOL) -path=/migrations/ -database '$(DATABASE_URL)' down 1
migrate-drop:
	docker run -v $(MIGRATION_DIR):/migrations --network host $(MIGRATION_TOOL) -path=/migrations/ -database '$(DATABASE_URL)' down -all
database-recreate: migrate-drop migrate-up

# Other
ps:
	docker-compose ps
