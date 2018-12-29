FROM golang:1.11.2-stretch as builder

WORKDIR /auth

COPY . ./

# Building using -mod=vendor, which will utilize the v
RUN cd /src/cmd/app && CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o app 

FROM alpine:3.8

WORKDIR /root/

RUN mkdir /k8slog

COPY --from=builder /auth/cmd/app/app .

CMD ["./app"]