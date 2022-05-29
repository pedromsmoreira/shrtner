.PHONY: infrastructure shrtener run test schema

include .env
$(eval export $(shell sed -ne 's/ *#.*$$//; /./ s/=.*$$// p' .env))

infrastructure:
	@docker-compose up -d cockroachdb

shrtener:
	@docker-compose up --build --abort-on-container-exit --exit-code-from shrtener shrtener

schema:
	@docker-compose up --build --abort-on-container-exit --exit-code-from schema schema

run: shrtener

test:
	go test -v ./...