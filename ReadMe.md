# Distributed Rate Limiter

## Overview
This project implements a distributed rate-limiting system using gRPC and Kubernetes operators. The system consists of three key components:

1. **RateService**: A gRPC service that manages rate limits using Redis.
2. **DummyService**: A FastAPI-based test service to validate rate limiting.
3. **RateLimiterOperator**: A Kubernetes operator that dynamically configures rate limits using custom resources.

## Architecture

- **RateService**:
  - Provides APIs to get and set request-per-minute (RPM) limits.
  - Limits requests based on IP over a 60-second window.
  - Stores rate limits in Redis.
  
- **DummyService**:
  - A test service making gRPC calls to RateService to validate rate limiting.
  
- **RateLimiterOperator**:
  - Monitors custom resource definitions (CRDs) for rate-limiting policies.
  - Updates the RateService database when configurations change.

## Repository Structure
```
.
├── Design.txt                # Design document
├── kubeconfig/               # Kubernetes configurations
│   ├── dummy-service/        # Deployment & Service configs for dummy service
│   ├── rateservice/          # Deployment config for RateService
│   ├── redis/                # Deployment & Service configs for Redis
├── dummyservice/             # FastAPI dummy service
│   ├── main.py               # API implementation
│   ├── requirements.txt      # Python dependencies
│   ├── Dockerfile            # Containerization script
├── ratelimiterapi/           # gRPC-based RateService
│   ├── main.go               # Main entry point
│   ├── internal/             # Internal logic for rate limiting
│   ├── handlers/             # API handlers
│   ├── proto/                # gRPC definitions
│   ├── Dockerfile            # Containerization script
├── ratelimiteroperator/      # Kubernetes operator
│   ├── api/                  # Custom resource definitions
│   ├── cmd/                  # Operator command-line interface
│   ├── internal/             # Controller logic
│   ├── config/               # Kubernetes CRDs and role bindings
│   ├── test/                 # End-to-end tests
└── test_script.py            # Script to test rate limiter
```

## Setup & Deployment

### Prerequisites
- Docker
- Kubernetes
- `kubectl`
- `helm`
- Redis
- Go (for RateService & Operator)
- Python (for DummyService)

### Steps


1. **Build and deploy RateService**:


2. **Build and deploy DummyService**:


3. **Build and deploy RateLimiterOperator**:

4. **Verify the behavior using the test script**:
   ```sh
   python test_script.py
   ```

## Future Improvements
- Cleanup, Logging and Testing

## License
This project is licensed under the MIT License.

