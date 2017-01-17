DEFAULT: build

build:
	@echo "=> Building binary"
	env GOOS=linux GOARCH=arm GOARM=7 go build -v .

copy: build
	@echo "=> Copy binary"
	scp ./icat-flap root@192.168.0.250:~/
