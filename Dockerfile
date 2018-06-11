FROM alpine:3.7
LABEL maintainer="Max Proske <mproske@sfu.ca>"

COPY ./cloud-native-go /app/cloud-native-go
RUN chmod +x /app/cloud-native-go

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT /app/cloud-native-go