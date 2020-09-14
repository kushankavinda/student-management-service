# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:1.12

# Add a new user "john" with user id 8877
RUN useradd -u 8877 john
# Change to non-root privilege
USER john

# Add Maintainer Info
LABEL maintainer="Kushan Edirisooriya<kushan@entrusttitle.com>"

# Copy the executable file
COPY ./main ./

# Expose port 5000 to the outside world
EXPOSE 5000

# Command to run the executable
CMD ["./main"]
