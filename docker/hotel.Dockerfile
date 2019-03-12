FROM golang:1.12-alpine3.9 AS build
RUN apk add --no-cache git
WORKDIR /opt/travel/
COPY . .
RUN CGO_ENABLED=0 go install \
    --ldflags "-s" -a -installsuffix cgo \
    ./cmd/hotel-server/

FROM scratch
USER nobody
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /go/bin/hotel-server /bin/hotel-server
ENTRYPOINT ["/bin/hotel-server"]
