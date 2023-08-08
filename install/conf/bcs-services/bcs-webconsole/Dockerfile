FROM alpine:3.18.3
RUN apk --update --no-cache add ca-certificates bash vim curl
COPY bcs-webconsole /bcs-webconsole
ENTRYPOINT ["/bcs-webconsole"]
CMD []
