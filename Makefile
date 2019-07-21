GIT_REF := $(shell git describe --always --tag)
VERSION ?= $(GIT_REF)

SERVICE_NAME := $(shell basename ${CURDIR})

.PHONY: build
build:
	@echo "+ $@"
	CGO_ENABLED=0 go build -o bin/server \
        -ldflags "-w -s" ./cmd

.PHONY: test
test:
	go test -v -race -cover -coverprofile=coverage.out  ./...

.PHONY: rich_test
rich_test:
	richgo test -v -race -cover -coverprofile=coverage.out  ./...

.PHONY: cover
cover: rich_test
	richgo tool cover -html=coverage.out

.PHONY: gen_swagger
gen_swagger:
	swag init -g cmd/server/main.go

.PHONY: up_migrate
up_migrate:
	go run cmd/db/migrate.go up ${arg}

.PHONY: down_migrate
down_migrate:
	go run cmd/db/migrate.go down ${arg}

.PHONY: drop_migrate
drop_migrate:
	go run cmd/db/migrate.go drop

.PHONY: check_migrate_version
check_migrate_version:
	go run cmd/db/migrate.go version

.PHONY: gen_migration
gen_migration:
	go run cmd/db/migrate.go new ${name}

.PHONY: gen_models
gen_models:
	rm -f model/*.xo.go
	xo mysql://root@$(DB_HOST):$(DB_PORT)/go_echo_boilerplate_development --int32-type int64 --uint32-type int64 --template-path xo/model_templates -o model
	rm model/schemamigration.xo.go

.PHONY: up_db
up_db:
	docker-compose -f docker-compose.dev.yml up -d db

.PHONY: down_db
down_db:
	docker-compose -f docker-compose.dev.yml down

