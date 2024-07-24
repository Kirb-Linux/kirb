FROM golang:1.23rc2-alpine3.19

RUN mkdir -p /build

COPY . /build

WORKDIR /build

RUN go build -o kirb.tmp

ENTRYPOINT ["./kirb.tmp"]