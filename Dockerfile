FROM golang:1.16.5-alpine as builder

WORKDIR /app

COPY . /app/

RUN CGO_ENABLED=0 go build -o /bin/ori

FROM scratch AS ori
ENV YML_FILE="/config.yml"

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /bin/ori /bin/ori
COPY --from=builder /etc/ssl /etc/ssl
COPY config.yml /config.yml

USER nobody

ENTRYPOINT [ "/bin/ori" ]

CMD []
