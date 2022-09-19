# FROM golang:1.19-alpine3.16
# RUN mkdir models-service
# WORKDIR /models-service
# COPY ./ ./
# RUN go build -o main ./cmd/main
# CMD ["./cmd/main/main"]

# FROM golang:1.19-alpine3.16 AS build 
# RUN apk
# WORKDIR /models-service
# COPY ./ ./

# RUN go install ./...

# FROM alpine:3.16  
# WORKDIR /usr/bin  
# COPY --from=build /go/bin .


FROM golang:1.19-alpine3.16

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/src/app/cmd/main ./...

CMD ["./cmd/main/main"]