# FROM golang tells docker to use the official golang image as a base image
FROM golang
#COPY . . copies all files from the current directory into the image
COPY . .
#RUN go build -o main main.go builds our application and names the output file main.
RUN go build -o main main.go

# Note:
    # The EXPOSE instruction tells docker to expose a port to the host. you will often see an EXPOSE instruction. The expose instruction actually doesn’t expose the port to the host,
    #
EXPOSE 8080
#CMD ["./main"] tells docker to run the main executable when the container starts.
CMD ["./main"]