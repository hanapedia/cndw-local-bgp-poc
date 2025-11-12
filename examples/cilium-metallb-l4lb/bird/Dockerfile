FROM ubuntu:jammy

RUN apt update && \
    apt install -y iproute2 iputils-ping traceroute tcpdump iptables bird2 curl

# We need this path at runtime
RUN mkdir -p /run/bird

ENTRYPOINT ["/usr/sbin/bird", "-c", "/etc/bird.conf", "-d"]
