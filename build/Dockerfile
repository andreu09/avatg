FROM alpine:3.18.0

CMD ["/bin/sh", "-c", "./avatg"]

EXPOSE 80

RUN adduser -D -g '' avatg

WORKDIR /opt/avatg/bin

RUN install -g avatg -o avatg -d log

COPY --chown=avatg:avatg avatg /opt/avatg/bin/

USER avatg
