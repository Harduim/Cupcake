#! /usr/bin/bash

docker compose -f services/docker-compose.yml --env-file backend/.env up -d