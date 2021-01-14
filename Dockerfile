FROM ubuntu:focal

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

RUN apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get install -qy \
  git \
  bc \
  bison \
  flex \
  libssl-dev \
  make \
  libc6-dev \
  libncurses5-dev \
  kmod \
  crossbuild-essential-armhf \
  crossbuild-essential-arm64

VOLUME [ "/build" ]
