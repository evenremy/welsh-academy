FROM scratch

ADD ["api", "/api.bin"]
ADD ["sql", "/sql/"]
ADD ["etc", "/etc/"]
EXPOSE 8888/tcp
ENTRYPOINT ["/api.bin"]
