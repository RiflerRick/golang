FROM golang:latest

RUN apt update

WORKDIR /go/app
COPY . /go/app

# not sure if building in the dockerfile is supposed to be done or not
RUN go build -o trie . 

RUN cd /go/app

ENTRYPOINT [ "./trie" ]

