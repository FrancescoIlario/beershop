FROM gcr.io/distroless/base-debian10 as prod
CMD ["/app"]

FROM golang:1.15.2 as build
WORKDIR /app
RUN apt-get update \
    && apt-get install -y --no-install-recommends make

COPY go.mod .
RUN go mod download \
    && go mod verify

COPY . .

RUN make build

FROM prod 
COPY --from=build /app/bin/beershop /app