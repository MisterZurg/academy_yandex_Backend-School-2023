FROM golang:1.20-alpine

# Hardcoded
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD password
ENV POSTGRES_DB db
ENV POSTGRES_HOST 172.17.0.2
ENV POSTGRES_PORT 5432

WORKDIR /usr/src/app

COPY . ./

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
RUN go mod download && go mod verify
RUN mkdir -p /usr/local/bin/
RUN go clean -modcache
RUN go mod tidy
RUN go build -v -o /usr/local/bin/app

CMD ["app"]