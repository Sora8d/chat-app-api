FROM golang:1.17

RUN mkdir users_api

COPY . users_api

WORKDIR users_api

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

CMD ["go", "run", "."]
