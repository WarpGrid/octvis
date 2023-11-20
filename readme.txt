build like this:

make sure dockerd is running

then in this dir:
	docker build -t octvis .
to run the container:
	docker run -p 5020:5020 octvis
