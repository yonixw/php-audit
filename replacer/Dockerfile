FROM golang:1.20 as builder
WORKDIR /app

COPY go.mod .
#COPY go.sum .
#RUN go mod download

RUN go get gopkg.in/yaml.v3 && cat go.mod


COPY . .

RUN go build -o replacer . && ls


FROM bash
COPY --from=builder /app /app
WORKDIR /app 
CMD ./replacer
