FROM scratch

ADD ["api", "/api.bin"]
COPY ["etc", "/etc/"]
EXPOSE 8888/tcp
ENTRYPOINT ["./api.bin"]
