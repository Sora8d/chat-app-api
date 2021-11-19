FROM golang:1.17

RUN mkdir rest_api

COPY . rest_api

WORKDIR rest_api

RUN go get -d -v ./...

RUN go install -v ./...



CMD ["go", "run", "."]
