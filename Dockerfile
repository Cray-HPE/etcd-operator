FROM busybox:latest

COPY app /usr/local/bin/app

ENTRYPOINT /usr/local/bin/app