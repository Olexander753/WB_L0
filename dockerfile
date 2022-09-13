FROM golang:1.19.1-alpine3.16 AS build 
RUN apk
WORKDIR /go/WB_L0

COPY Gopkg.lock Gopkg.toml ./  
COPY vendor vendor  
COPY cmd cmcd  
COPY pkg pkg 
COPY postgres postgres

RUN go install ./...

FROM alpine:3.16  
WORKDIR /usr/bin  
COPY --from=build /go/bin .