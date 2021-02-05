FROM golang:1.14.4-stretch

WORKDIR /go/src

RUN git clone https://github.com/ciromacedo/nf_nwdaf.git

# install libs
RUN apt update
RUN apt install nano
RUN go get -u github.com/urfave/cli
RUN go get -u github.com/calee0219/fatal

#RUN cd nf_nwdaf && go run  app.go

# esta porta deve ser a mesma do arquivo de configuração ./config/config.json
EXPOSE 5003
