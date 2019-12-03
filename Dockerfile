FROM golang:1.12
COPY ./main /bin/main
COPY ./libs/google-services.json /home/
EXPOSE 4300
CMD ["/bin/main","/home/google-services.json"]