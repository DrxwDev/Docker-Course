# Docker Valuable Info

Docker is a Platform that allows to al inside a Container

- Develop
- Ship
- Run applications

## Container

- Standardized unit of software
- Packages code and all the dependencies
- Application runs quickly and reliably
- Container guarantees the application will work

## Docker Work Flow

- docker run -> Docker CLI -> Docker Daemon ->
  Create and run container -> Container output

## Creating a container

- docker all flags then the image
  - Example drn -d --name server -p 8000:80 nginx
  - Example drn -d --name page -p 8000:80 -v localDir:containerDir nginx
  - localDir -> local product directory
  - containerDir -> where do you want to be in the container
