FROM golang:1.12

LABEL maintainer="Alexandru Agrapine <alexandru@agrapine.com>"


WORKDIR $GOPATH/src/github.com/agrapine/glow

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

CMD ["glow"]