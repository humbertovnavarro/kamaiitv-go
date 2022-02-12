FROM golang:alpine
ENV GO111MODULE=on
EXPOSE 8080
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go get github.com/cosmtrek/air
COPY . .
ENTRYPOINT ["air", "-c", ".air.toml"]
