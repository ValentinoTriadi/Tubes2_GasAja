FROM golang:1.22.2-bookworm AS development

# Create app directory
WORKDIR /app/backend

COPY . .

# Download and install the dependencies:
RUN go mod vendor
RUN go mod tidy --compat=1.17

# Build the Go app
# RUN go build -o scraper .

EXPOSE 8000

CMD ["go", "run", "."]

