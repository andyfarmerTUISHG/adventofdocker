# FROM golang tells docker to use the official golang image as a base image
FROM golang

#COPY . . copies all files from the current directory into the image
COPY . .

# This defines that the build argument IS_PRODUCTION is set to false by default.
ARG IS_PRODUCTION=false
# if its production, add a compilation flag

#RUN go build -o main main.go builds our application and names the output file main.
# RUN if [ "$IS_PRODUCTION" = "true" ]; then go build -o main main.go -ldflags "-s -w"; else go build -o main main.go; fi
RUN if [ "$IS_PRODUCTION" = "true" ]; then CGO_ENABLED=0 go build -o main main.go; else go build -o main main.go; fi


# Note:
    # The EXPOSE instruction tells docker to expose a port to the host. you will often see an EXPOSE instruction. The expose instruction actually doesn’t expose the port to the host,
    #
EXPOSE 8080
#CMD ["./main"] tells docker to run the main executable when the container starts.
CMD ["./main"]