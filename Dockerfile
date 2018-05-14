FROM golang:1.10
COPY . .
RUN go get -d
RUN CGO_ENABLED=0 go build main.go

FROM scratch
COPY --from=0 /go/main /karousel
CMD ["/karousel"]
