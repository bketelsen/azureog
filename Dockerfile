FROM ubuntu:20.04
COPY azureog /usr/bin/azureog
ENTRYPOINT ["/usr/bin/azureog"]