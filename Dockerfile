FROM scratch
# set labels for metadata
LABEL maintainer="Oliver Guggenbuehl<me@containerize.ch>" \
    name="base" \
    description="GIN API Container" \
    summary="GIN API Container"

ENV GIN_MODE=release \
    PORT=8080 \
    TZ=Europe/Zurich

# set Env Vars
EXPOSE $PORT

WORKDIR /usr/local/bin
COPY bin/trinity_server .
CMD ["trinity_server"]