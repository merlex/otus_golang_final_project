FROM golang:1.22

RUN mkdir -p /opt/integration_tests
WORKDIR /opt/integration_tests

RUN export GO111MODULE=on

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

CMD ["go", "test"]
