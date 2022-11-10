FROM scratch

ADD ["api", "/app/api.bin"]
ADD ["sql", "/app/sql/"]
ADD ["etc", "/app/etc/"]
EXPOSE 8888/tcp
ENTRYPOINT ["/app/api.bin"]
