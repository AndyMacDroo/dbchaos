FROM golang:1.14

WORKDIR $GOPATH/dbchaos

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...