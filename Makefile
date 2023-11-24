.PHONY: build
build:
	rm -f go-auth-api
	go build -o go-auth-api cmd/main.go

.PHONY: build_image
build_image:
	docker build -t go-auth-api .

.PHONY: delete_image
delete_image:
	docker rmi go-auth-api:latest

.PHONY: start
start:
	docker-compose up -d

.PHONY: stop
stop:
	docker-compose stop && docker-compose rm -f

.PHONY: start_web
start_web:
	cd webserver && docker-compose up -d

.PHONY: stop_web
stop_web:
	cd webserver && docker-compose stop && docker-compose rm -f

.PHONY: start_auth
start_auth: build
	./go-auth-api

.PHONY: create_test_user
create_test_user:
	bash scripts/register_test.sh

.PHONY: auth_test
auth_test:
	bash scripts/auth_test.sh

.PHONY: public_page_test
public_page_test:
	bash scripts/public_page_test.sh

.PHONY: private_page_test
private_page_test:
	bash scripts/private_page_test.sh
