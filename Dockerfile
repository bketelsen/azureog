FROM scratch
COPY azureog /usr/bin/azureog
ENTRYPOINT ["/usr/bin/azureog"]