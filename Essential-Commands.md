# Docker Essential Commands

- docker run
- docker ps
- docker ps -a
- docker images ls
- docker container ls
- docker volume
- docker container logs
- docker container inspect -> returns json info
- docker build -> Command to build the image

## Flags

-it Interactive shell
-tty Allocates a pseudo-TTy, essentially giving us a proper terminal prompt
-d Detached
-e Environment Variable
-p Port Binding
--name Name a container
--rm remove container after it exits
--network connect container to a network
-f -> dlo -f container ID -> show logs in real time
-docker build -t -> Name the image and give a tag
