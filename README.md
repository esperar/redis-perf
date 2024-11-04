# Redis Perf

This is a Go-based project designed to evaluate the performance, consistency, and memory management of Redis. It supports three main scenarios:
1. High-throughput data processing
2. Consistency and failover testing
3. TTL expiration and memory management

## Features

### 1. High-Throughput Data Processing
- **Description:** Measures Redis throughput by processing large volumes of `SET` and `GET` requests.
- **Usage:** Runs bulk requests to analyze transactions per second (TPS) and response latency.

### 2. Consistency & Failover Testing
- **Description:** Tests Redis data consistency during simulated failover events.
- **Usage:** Simulates master node failure, verifies if data is correctly available after failover.

### 3. TTL Expiration & Memory Management
- **Description:** Evaluates Redis key expiration and memory management, verifying `EXPIRE` behavior and `maxmemory-policy`.
- **Usage:** Assigns different TTL values to keys and checks for correct data expiry.

---

## Installation

Clone the repository and build the project:
```bash
git clone https://github.com/your-repo/redis-testing-tool.git
cd redis-testing-tool
go build -o redis-tester main.go
```

## Configuration

Set up Redis and configure as needed in `config.yaml`:
```yaml
redis:
  address: "localhost:6379"
  password: ""         # Redis password if set
  db: 0                # Database index
  maxMemoryPolicy: "allkeys-lru"  # Redis eviction policy for TTL tests
test:
  throughputRequests: 10000       # Number of requests for throughput test
  ttlExpiration: [5, 10, 20]      # TTL values in seconds for keys
  failoverSimulation: true        # Enable failover testing
```

## Usage

Each test can be executed independently or all together:
```bash
./redis-tester throughput       # Run throughput test
./redis-tester consistency       # Run failover and consistency test
./redis-tester ttl              # Run TTL expiration and memory management test
```

## Example Output

```
Throughput Test:
    Transactions/sec: 35,000
    Average latency: 1.2ms

Failover Test:
    Failover triggered. Data consistency verified.

TTL Expiration Test:
    Key TTLs expired correctly. Maxmemory-policy enforced.
```

## Requirements

- Redis server (locally or remotely accessible)
- Go 1.18 or later

## Contributing

Contributions are welcome! Feel free to submit issues and pull requests.

## License

MIT License.