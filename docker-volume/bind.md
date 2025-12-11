# Bind Mounts

- Very different
- Maps specific file or directory
- From host machine to container
- Creates a two-way mirror
- Fantastic for development
- Not great for production persistence

> Example:
> docker run -v /home/user/my_app:/app my-image
