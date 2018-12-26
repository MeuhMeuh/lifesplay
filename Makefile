.PHONY: ng start

ng:
	ngrok http --subdomain=${NGROK_SUBDOMAIN} 3555

docker:
	# -d = detached containers
	docker-compose up -d

docker-ssh:
	# -t = tty | -i = interactive mode
	# lifesplay is the container's name as defined in the Dockerfile
	docker exec -ti lifesplay /bin/sh
docker-logs:
	# -f : Tailing
	docker logs lifesplay -f
