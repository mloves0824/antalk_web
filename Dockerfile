FROM golang:1.10

RUN apt-get update
RUN apt-get install -y autoconf

ENV GOPATH /gopath
ENV ANTALK_WEB  ${GOPATH}/src/github.com/mloves0824/antalk_web
ENV PATH   ${GOPATH}/bin:${PATH}:${ANTALK_WEB}/bin
COPY . ${ANTALK_WEB}

RUN make -C ${ANTALK_WEB} clean
RUN make -C ${ANTALK_WEB} all

WORKDIR /antalk_web
