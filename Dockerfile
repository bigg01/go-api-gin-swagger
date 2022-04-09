FROM scratch
WORKDIR /usr/local/bin
COPY bin/main .
CMD ["main"]