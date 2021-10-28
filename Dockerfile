FROM ubuntu:20.04
COPY ./dist/azureog /usr/bin/azureog
ENTRYPOINT ["/usr/bin/azureog"]