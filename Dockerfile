FROM golang:1.6

# Install beego and the bee dev tool
RUN go get github.com/astaxie/beego && go get github.com/beego/bee

# Create the directory where the application will reside
RUN mkdir /app

# Add the local sources
ADD . /app/

# Set the working directory to the app directory
WORKDIR /app

RUN go test
RUN go build -o main .

# Expose the application on port 8080.
# This should be the same as in the app.conf file
EXPOSE 8080

CMD ["/app/main"]
