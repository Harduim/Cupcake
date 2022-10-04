setup-backend-dev:
	docker-compose -f services/docker-compose.yml --env-file backend/.env.example up -d

destroy-backend-dev:
	docker-compose -f services/docker-compose.yml --env-file backend/.env.example down


.PHONY: server test