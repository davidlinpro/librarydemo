.PHONY: docker seed clean frontend-dev

docker:
	docker-compose up -d --build

seed:
	cat ./db/library.sql | docker exec -i library_db /usr/bin/mysql -u root --password=library library

clean:
	docker-compose down --rmi all --remove-orphans

frontend-dev:
	cd frontend && yarn install && yarn serve
