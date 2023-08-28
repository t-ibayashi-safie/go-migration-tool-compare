down:
	docker compose down

run: clean
	docker compose up -d

clean:
	rm -rf ./docker/db/data/*
