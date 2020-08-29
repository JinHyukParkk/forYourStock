all: exec

clean:
	

build:
	@go build .

exec:
	@./shell/start.sh
