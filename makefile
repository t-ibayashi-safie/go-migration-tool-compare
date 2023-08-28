down:
	docker compose down

run: clean
	docker compose up -d

clean:
	rm -rf ./docker/db/data/*

mysql:
	docker compose exec mysql mysql -u root
