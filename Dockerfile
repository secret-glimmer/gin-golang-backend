FROM golang:1.21.4

# Set destination for COPY
WORKDIR /opt/app

# Prepare swag cli
RUN export PATH=$(go env GOPATH)/bin:$PATH

# Copy the source code.
COPY . .

# Download Go modules
RUN go mod download

RUN go build -o app ./cmd/...


ENTRYPOINT [ "/opt/app/app" ]