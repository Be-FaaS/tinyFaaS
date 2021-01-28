FROM golang as golang

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 go install -v ./...

FROM scratch

EXPOSE 80/tcp
EXPOSE 6000/udp
EXPOSE 7000/tcp
EXPOSE 8000/tcp

COPY --from=golang /go/bin/reverse-proxy app

ENTRYPOINT [ "./app" ]

