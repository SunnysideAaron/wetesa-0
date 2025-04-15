# startUP the datastore service
# NOTE to see logs take off -d flag OR
# docker logs -f datastore
ds-up:
	@docker compose up --remove-orphans datastore -d

# connect to an already running datastore service
ds-bash:
	@docker compose exec datastore bash

# Shutdown the datastore service
ds-down:
	@docker compose down datastore

# **************************************************

api-build:
#	docker compose build api --no-cache
	@docker compose build api

# run sh in service (starts service)
api-bash:
	@docker compose run --remove-orphans --service-ports api bash

api-watch:
# results in same thing. Leave in case I want later for some kind of troubleshooting.
#	@docker compose run --remove-orphans --service-ports  api bash -c "make watch"
	@docker compose up --remove-orphans api

# **************************************************

all-up:
	@docker compose up --remove-orphans

all-down:
	@docker compose down

# **************************************************

.PHONY: ds-up ds-bash ds-down api-build api-bash api-watch all-up all-down
