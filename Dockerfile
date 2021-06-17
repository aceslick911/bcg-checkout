# ==== build - Build main artifacts
FROM golang:1.15.3-alpine AS build

  ENV CGO_ENABLED=1
  ENV GO111MODULE=on
  RUN apk add --no-cache git git gcc g++ make sqlite jq sqlitebiter

  # Set the Current Working Directory inside the container
  WORKDIR /src

  # We want to populate the module cache based on the go.{mod,sum} files.
  COPY src/go.mod .
  COPY src/go.sum .
  COPY src/Makefile .

  # Print copied files
  #RUN pwd && find .

  RUN make dep

  RUN mkdir -p /src/cmd/api
  COPY src/ .

  # Print copied files
  #RUN pwd && find .

  # Build
  RUN make build

  # Print copied files
  #RUN pwd && find .

# ==== docs - Documentation build
FROM build as docs

  # Print copied files
  #RUN pwd && find .

  WORKDIR /src
  RUN make docs
  
# ==== docs-output - Output target
  FROM scratch as docs-output
    COPY --from=docs /src/docs /docs


# ==== app - Build app into alpine
FROM alpine:3.12 as app
  RUN apk add ca-certificates

  WORKDIR /app

  # Print copied files
  #RUN pwd && find .

  COPY --from=build /src/bin/restapi /app/restapi
  COPY --from=build /src/data /app/data
  COPY --from=docs /src/docs /app/data

  RUN chmod +x restapi

  # This container exposes port 3000 to the outside world
  EXPOSE 3000

  # Run the binary program produced by `go install`
  ENTRYPOINT ./restapi
