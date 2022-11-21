setup-backend-dev:
	docker compose -f services/docker-compose.yml --env-file backend/.env up -d

destroy-backend-dev:
	docker compose -f services/docker-compose.yml --env-file backend/.env down --rmi local

build-image:
	bash services/backend/build.sh 


.PHONY: server test