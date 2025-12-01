# Layers of a Image

Images are Read-only File system

- Cannot modify an image once it is built.
- Image is static and read-only.
- Container adds a temporary, writable layer.

Every instruction you have in your Dockerfile
will be a layer.

So a docker image can have as many layers as yours
instruction grow.

## Layer 1

This will be the minimal operating system

- Base image ex: ubuntu, node, go

## Layer 2

Will be commands on top of the base layer

- RUN apt-get update && apt-get install -y python3

## Layer 3

Copy Command

- This layer adds the application source code into the app directory

## Layer 4

This layer set the defaults command to run

CMD ["python3", "/app/main.py"]
