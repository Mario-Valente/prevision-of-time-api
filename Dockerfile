FROM golang:1.21-alpine AS build

WORKDIR /go/src/app

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod/ \
         go mod download -x

RUN --mount=type=cache,target=/go/pkg/mod/ \
      CGO_ENABLED=0 go build -o /go/bin/app


FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/app /

CMD ["/app" ]
