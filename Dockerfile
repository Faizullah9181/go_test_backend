FROM golang:1.16-alpine

WORKDIR /go_test_backend

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . ./
RUN ls
RUN go build -o /go_test_backend_app main.go

EXPOSE 8090:3030
# Run
CMD [ "/go_test_backend_app" ]


ENV    MYSQL_HOST: ${MYSQL_HOST}
ENV    MYSQL_USER: ${MYSQL_USER}
ENV    MYSQL_PASSWORD: ${MYSQL_PASSWORD}
ENV    MYSQL_DBNAME: ${MYSQL_DBNAME}
ENV    JWT_SECRET: lolo