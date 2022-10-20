FROM golang:1.19 as builder

WORKDIR /src/
COPY *.go go.mod go.sum ./
RUN go mod download
RUN go build -o keysync-test ./

# Copy from the builder to keep the final image reproducible and small.  If we
# don't do this, we end up with non-deterministic build artifacts.
FROM scratch
COPY --from=builder /src/keysync-test /
EXPOSE 8443

CMD ["/keysync-test"]
