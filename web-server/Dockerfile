# Base image for building the go project
FROM golang:1.14-alpine AS build

LABEL name="builder"
LABEL Production="false"
# Updates the repository and installs git
#RUN apk update && apk upgrade && apk add --no-cache git

# Switches to /tmp/app as the working directory, similar to 'cd'
WORKDIR /tmp/app

## If you have a go.mod and go.sum file in your project, uncomment lines 13, 14, 15
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY main.go ./
COPY queries ./queries
RUN ls .
# Builds the current project to a binary file called server
# The location of the binary file is /tmp/app/out/api
RUN GOOS=linux go build -o ./out/webserver .

#########################################################
# The project has been successfully built and we will use a
# lightweight alpine image to run the server 
FROM alpine:latest AS web-stage

ARG SERVER_PORT=8000
ARG DB_SERVER
ARG NIC_NAME

LABEL name="webserver"
LABEL Production="true"

# Adds CA Certificates to the image
RUN apk add ca-certificates

# Copies the binary file from the BUILD container to /app folder
COPY --from=build /tmp/app/out/webserver /app/webserver

# Switches working directory to /app
WORKDIR "/app"

ENV SERVER_PORT_ENV=${SERVER_PORT}
ENV DB_SERVER_ENV=${DB_SERVER}
ENV NIC_NAME_ENV=${NIC_NAME}

COPY start-server.sh .

RUN chmod +x start-server.sh
#ADD img.tar.gz /app/

# Exposes the port from the container
EXPOSE $SERVER_PORT

# Runs the binary once the container starts
#CMD ["sh", "-c", "/app/server", "--port", "${SERVER_PORT}", "--addr", "$(cat ip.conf)"]
#CMD ["sh", "-c", "cat /etc/hosts && cat /app/start_server.sh && /app/start_server.sh"]
CMD ["/app/start-server.sh", "$NIC_NAME"]