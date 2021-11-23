FROM golang:1.17

RUN mkdir oauth_api

COPY . oauth_api

WORKDIR oauth_api

RUN go get -d -v ./...

RUN go install -v ./...



CMD ["go", "run", "."]
