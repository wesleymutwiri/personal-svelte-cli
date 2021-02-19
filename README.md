# Svelte CLI

This is a simple Golang script to initialize svelte projects rather than cramming npx and degit commands all the time.

In future I'll probably convert it into a Vue cli for svelte clone. Hopefully.


## Usage
Clone the repository into your machine using 
```bash
$ git clone https://github.com/wesleymutwiri/personal-svelte-cli.git
```
Change directory into the folder downloaded
```bash
$ cd personal-svelte-cli
```
Either 

1. Install using the Dockerfile provided and run the following commands for using the app from the dockerfile
```bash
$ docker build -t svelte-cli .

$ docker run -it svelte-cli create <name_and_full_url_of_folder_to_be_created>
```
    Still in progress

2. Install using the makefile provided
```bash 
$ make compile

$ /bin/svelte-cli-<current_os> create <name_and_full_url_of_folder_to_be_created>
```
3. Manually by building the go file if Go is installed on your machine
```bash
$ go build .
$ ./personal-svelte-cli create <name_and_full_url_of_folder_to_be_created>
```

4. Using the go run command for fast compilation
```bash
$ go run main.go create <name_and_full_url_of_folder_to_be_created>
```