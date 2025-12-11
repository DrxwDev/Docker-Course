# Docker Network

## Docker networking

- Isolation
- Outbound connectivity
- Inbound port mapping

## Three default networks

- Bridge
  - Default network
  - Virtual network
  - Private network switch
  - Communicate with other containers
  - Communicate with outside world
  - Network Address Translation
- Host
  - Simpler but less isolated
  - Bypass Docker's network isolation
  - Uses host's own network
  - Does not get own private IP
  - No NAT overhead
  - Lack of isolation
- None
  - Simplest of all networks
  - Has no networking interfaces
  - Has loopback interface
  - Cannot communicate with other containers
  - Security-sensitive workloads
  - No network access required

## How containers communicate

- Automatic DNS
- IP-based communication

## Custom Network

- Automatic DNS resolution
- Better isolation
