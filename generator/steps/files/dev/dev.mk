_COMPOSE=docker compose -f dev/docker-compose.yml --project-name ${NAMESPACE}

dev-up: ## Up the environment in docker compose
	${_COMPOSE} up -d

dev-down: ## Down the environment in docker compose
	${_COMPOSE} down --remove-orphans

dev-clean: ## Down the environment in docker compose with image cleanup
	${_COMPOSE} down --remove-orphans -v --rmi all
