# Dockerfile

## What is a Dockerfile

- Blueprint or recipe
- Text file
- Series of commands or instructions
- Create a New layered Docker image
- Each instructions creates a new layer in the image
- Automation, repeatability, and transparency

## Key instructions

- FROM
  - Most important instruction
  - First non-comment line in Dockerfile
  - Sets base image for your build
- RUN
  - Executes command inside container
  - During build process
  - Installing packages, running scripts and config envs
- COPY
  - Copy files and directories
  - From host machine
  - Into image's filesystem
- ADD
  - Similar to COPY instruction
  - Has extra features
  - Handle tar files or copy files from a URL
- WORKDIR
  - Sets working directory
  - RUN, CMD, ENTRYPOINT, COPY, and ADD instructions
  - Same as the cd command
- EXPOSE
  - Documentation instruction
  - Listen on the specified network ports at runtime
  - Does not actually publish the port
- CMD
  - Provides default command
  - Runs when container starts
  - Only one CMD instruction in a Dockerfile
- ENTRYPOINT
  - Will run as an executable
  - Behave like a binary
  - Provide default arguments to the ENTRYPOINT

## Best practices

- Use .dockerignore file
  - Which files and directories to ignore
  - Prevents sending unnecessary files
  - Improve build performance and security
- Layer caching
  - Each instruction creates a cacheable layer
  - Use cache for as many layer as possible
  - Place less frequently changed instructions at the top
- Combine RUN instructions
- Small Base Images
  - Start with alpine-based images
  - Reduce your images size
  - Smaller images are faster to transfer and more secure
- Run as a non-root user
  - Create a new user for Good security practice
  - Create and switch to a non-root user
  - If application doesn't require root privileges
