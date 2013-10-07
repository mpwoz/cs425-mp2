all: client

client:
	go build dgroup.go

clean:
	rm dgroup

