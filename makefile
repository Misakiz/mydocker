all: clean bulid install
build:
        go build -o /bin/docker
uninstall:
        rm /bin/docker
clean: unsintall