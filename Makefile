compile:
	echo "Compiling the Svelte ClI app"
	GOOS=freebsd GOARCH=386 go build -o bin/svelte-cli-freebsd main.go
	GOOS=linux GOARCH=386 go build -o bin/svelte-cli-linux main.go
	GOOS=windows GOARCH=386 go build -o bin/svelte-cli-windows main.go