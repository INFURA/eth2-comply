FROM golang:1.15-alpine

WORKDIR /eth2-comply

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY pkg pkg

RUN go build -o /usr/bin/eth2-comply cmd/eth2-comply/main.go

COPY tests tests

ENTRYPOINT ["eth2-comply", "--testsRoot", "tests"]
