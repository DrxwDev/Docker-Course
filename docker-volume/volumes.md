# Docker Data Volumes

## Volumes

- Storage units managed by Docker
- Best practice for persisting data
- Store in dedicated directory
- var/lib/docker/volumes/
- Docker Daemon is in control
- Portable and secure

## Anonymous volumes

- Created automatically on the fly
- Creates volume with random name
- Useful for temporary, disposable data
- Using the -v flag in docker run command
- Specify path inside the container
- Or VOLUME instruction in Dockerfile

Example:

- docker run -v /var/www/html nginx

## Named volumes

- Should use 99% of the time
- Explicitly create and name the volume
  - **docker volume create my_volume_name**
- Or, use it in a run command and Docker will create it if it doesn't exist
  - **docker run -v my_volume_name:container_directory docker_image**

## Flags

- -v or --volume
- Syntax -v source:destination:option
- Command docker run -v my_volume:/data my-image

- --mount
- Command docker run --mount type=volume,source=my_volume,target=/data my-image

## Mount

- Explicitness
- No Magic Behavior
- Easier for Complex Scenarios
