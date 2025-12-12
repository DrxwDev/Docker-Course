# Advanced features

## Depends on

- Control start-up order
- Container to be running
- Does not validate service
- Condition syntax

## Build config

- build.
- Provide detailed instructions
- Use build key

BUILD Example:

```yaml
build:
  context:
    dockerfile: file_path
    args:
      - BUILD_ENV=production
```

## Profiles

- Define services
- Start in specific situations
- Development only tools
- Debugging utilities
- Alternative services

Profiles Example:

```yaml
services:
  web:
  # ... main web config
  redis:
  # ... redis config
  redis-commander:
    image: rediscommander/redis-commander:latest
    ports:
      - "8081:8081"
    profiles:
      - dev
```

- Activate a docker profile:
  - docker-compose --profile profile_name up
  - docker-compose up --profile profile_name
