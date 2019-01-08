FROM golang:1.11.2-stretch as builder

WORKDIR /auth

COPY . ./

# Building using -mod=vendor, which will utilize the v
RUN  && CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o /src/cmd/app/app /src/cmd/app

FROM alpine:3.8

WORKDIR /root/

RUN mkdir /k8slog

COPY --from=builder /auth/cmd/app/app .

CMD ["./app"]