

FROM golang:1.15.11-alpine3.13

# work port
ARG VERSION=0.0.0
ARG EXECNAME=webapi

# create directories
RUN mkdir -p /app
ADD . /app
WORKDIR /app

# build
RUN go build -o ./build/app.exe odonto/cmd/${EXECNAME}
# create script
RUN echo "#!/bin/bash" > run.sh
RUN echo "/app/build/app.exe -c configs/context/app-dev.yaml -v ${VERSION}" >> run.sh

# command
CMD [ "/bin/sh", "/app/run.sh" ]

