app_name = go_webassembly_calculator

to_container:
	echo -ne "\033]0;$(app_name)\007" && docker exec -it $(app_name) bash

run:
	echo -ne "\033]0;$(app_name)\007" && docker-compose up

run_and_build:
	echo -ne "\033]0;$(app_name)\007" && docker-compose up --build

compile:
	GOARCH=wasm GOOS=js go build -o lib.wasm main.go
