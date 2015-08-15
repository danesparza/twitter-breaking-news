# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# To configure the app, set these environment variables or use the command line flags
ENV TWITTER_ALLOWED_ORIGINS *
ENV TWITTER_CONSUMER_KEY YOUR_CONSUMER_KEY
ENV TWITTER_CONSUMER_SECRET YOUR_CONSUMER_SECRET
ENV TWITTER_AUTH_TOKEN YOUR_AUTH_TOKEN
ENV TWITTER_AUTH_SECRET YOUR_AUTH_SECRET

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/danesparza/twitter-breaking-news

# Build and install the app inside the container.
RUN go get github.com/danesparza/twitter-breaking-news/...

# Run the app by default when the container starts.
ENTRYPOINT /go/bin/twitter-breaking-news

# Document that the app listens on port 3000.
EXPOSE 3000