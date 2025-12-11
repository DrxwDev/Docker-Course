# tmpfs Mounts

- Don't persist data at all
- Stored in host's machine memory only
  - **docker run --tmpfs /app/cache my-image**
  - Or the more modern --mount flag
- Temporary, sensitive, or high-performance data
- Session caches, swap files, or storing temporary secrets
- Blazingly fast because it's in RAM
- Ultimate ephemeral storage
