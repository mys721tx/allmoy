# syntax=docker/dockerfile:1

FROM golang:1.22 as builder

WORKDIR /app

COPY . .

RUN make

FROM gcr.io/distroless/static:nonroot

WORKDIR /app

COPY --from=builder --chown=nonroot:nonroot /app/allmoy /app/allmoy
COPY --chown=nonroot:nonroot providers_example.yaml /app/providers.yaml

ENTRYPOINT ["/app/allmoy"]
