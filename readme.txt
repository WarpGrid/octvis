build like this:

make sure dockerd is running

then in this dir:
	docker build -t octvis .
to run the container:
	docker run -p 8090:8090 octvis
