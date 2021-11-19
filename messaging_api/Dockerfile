FROM golang:1.17

RUN mkdir messaging_api

COPY . messaging_api

WORKDIR messaging_api

RUN go get -d -v ./...

RUN go install -v ./...



CMD ["go", "run", "."]
