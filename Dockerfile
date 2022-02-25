FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /golang_database_sample

EXPOSE 18080

CMD [ "/golang_database_sample" ]   