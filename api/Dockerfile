FROM scratch

WORKDIR /app
ADD ["api", "./api.bin"]
ADD ["sql", "./sql/"]
ADD ["etc", "./etc/"]
EXPOSE 8888/tcp

CMD ["./api.bin"]
