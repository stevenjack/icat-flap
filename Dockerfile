FROM hypriot/rpi-golang
MAINTAINER Steven Jack <stevenmajack@gmail.com>

RUN go get github.com/stianeikeland/go-rpio
RUN go get github.com/brian-armstrong/gpio

RUN mkdir -p /goroot1.5/src/github.com/stevenjack/icat-flap
ADD . /goroot1.5/src/github.com/stevenjack/icat-flap
WORKDIR /goroot1.5/src/github.com/stevenjack/icat-flap

RUN go build

ENTRYPOINT ["./icat-flap"]
