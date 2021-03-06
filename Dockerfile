FROM golang:1.14.4-stretch

WORKDIR /go/src

RUN git clone https://github.com/ciromacedo/nf_nwdaf.git

# install libs
RUN apt update
RUN apt install nano
RUN go get -u github.com/urfave/cli
RUN go get -u github.com/calee0219/fatal
RUN go get -u go.mongodb.org/mongo-driver/bson

WORKDIR /go/src/nf_nwdaf


