DEV_URL := http://localhost:1323
PRD_URL := https://ojosama.herokuapp.com
ifeq ($(ENV),prd)
URL := $(PRD_URL)
else
URL := $(DEV_URL)
endif

ojosama-web: go.* *.go
	go vet .
	go fmt .
	go build .

.PHONY: ping
ping:
	curl $(URL)/api/ping
	curl $(URL)/api/version
	curl \
		-X POST \
		-H 'Content-Type: application/json' \
		-d '{"Text":"これはハーブです！"}' \
		$(URL)/api/ojosama

.PHONY: deploy
deploy:
	git push heroku main
