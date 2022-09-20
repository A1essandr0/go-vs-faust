FROM golang:1.19-alpine as builder

WORKDIR /install

COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 go build -o /install/go-collector


FROM gcr.io/distroless/static-debian11

WORKDIR /app

COPY --from=busybox:1.35.0-uclibc /bin/sh /bin/sh
COPY --from=busybox:1.35.0-uclibc /bin/sleep /bin/sleep

COPY --from=builder /install/go-collector /app

COPY ./wait-and-start.sh ./