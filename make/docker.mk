# Installation of docker compose is described here
# https://docs.docker.com/compose/install/linux/

docker/up:
	docker compose up -d
	docker ps

docker/down:
	docker compose down

docker/prune:
	docker system prune -a
