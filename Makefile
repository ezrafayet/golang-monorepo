build:
	@$(MAKE) -C apps/service-a build
	@$(MAKE) -C apps/service-b build

dev:
	sudo docker-compose -f docker-compose.yml build
	sudo docker-compose -f docker-compose.yml up