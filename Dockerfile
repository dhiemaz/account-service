FROM golang:1.22-alpine

WORKDIR $GOPATH/src/account_svc
COPY go.mod ./
RUN mkdir -p $GOPATH/src/account_svc

RUN apk add build-base
COPY . $GOPATH/src/account_svc

RUN make build
EXPOSE 8009

CMD ["/go/src/account_svc/account_svc","grpc"]