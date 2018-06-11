FROM golang:1.10.3-alpine
LABEL maintainer="Max Proske <mproske@sfu.ca>"

ENV SOURCES /go/src/github.com/maxproske/cloud-native-go/

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8080

EXPOSE 8080

# Name of the executable being built
ENTRYPOINT cloud-native-go 
