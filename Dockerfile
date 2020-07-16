FROM golang

ADD . /go/src/github.com/J-Thompson12/planet-express
RUN apt-get update && apt-get install -y make protobuf-compiler
RUN export GO111MODULE=on
RUN go get github.com/golang/protobuf/protoc-gen-go
RUN go get google.golang.org/grpc

RUN go install /go/src/github.com/J-Thompson12/planet-express/ship
EXPOSE 10000

ENTRYPOINT ["/go/bin/ship"]