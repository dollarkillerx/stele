FROM alpine:latest
WORKDIR /app
COPY ./cmd/stele/stele /app
COPY ./cert/ /app/cert
#COPY ./configs/config.yml /app/configs/config.yml

CMD ["/app/stele"]