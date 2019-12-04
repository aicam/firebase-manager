FROM golang:1.12
COPY ./main /bin/main
COPY ./libs/google-services.json /home/
#RUN adduser -H -h /app appuser
#USER appuser
