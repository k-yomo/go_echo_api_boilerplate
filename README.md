# Go API boilerplate built with echo framework
[![CircleCI](https://circleci.com/gh/k-yomo/go_echo_boilerplate/tree/master.svg?style=svg&circle-token=f3e183e8330bf74666f5916886927656c31ad777)](https://circleci.com/gh/k-yomo/go_echo_boilerplate/tree/master)
[![codecov](https://codecov.io/gh/k-yomo/go_echo_boilerplate/branch/master/graph/badge.svg?token=cGgCiXQXVc)](https://codecov.io/gh/k-yomo/go_echo_boilerplate)


## Getting Started

### Prerequisites

- Go 1.12 (skip if you use docker)
- MySQL 5.7 (skip if you use docker)
- direnv

### External Service
- Twilio (for SMS authentication)
- SendGrid (for sending email)

### Usage

1. Clone repo
```
git clone https://github.com/k-yomo/go_echo_boilerplate.git
cd go_echo_boilerplate
```

2. Create `.env` file in reference to .env.sample

3. Install dependent modules

```
go mod install
```

4. Run dev server
```
// Listening on localhost:1323 with hot reloading(localhost:5002)
realize start --server
```

### Usage with Docker

1. Clone repo
```
git clone https://github.com/k-yomo/go_echo_boilerplate.git
cd go_echo_boilerplate
```

2. Create `.env` file in reference to .env.sample

3. Run containers
```
// Listening on localhost:1323 with hot reloading(localhost:5002)
docker-compose up -d
```

## Running the tests

```
make test
```


### with coverage
```
make cover
```

## API Docs
1. Run dev server
2. Open `localhost:1323/swagger/index.html`

## Migration
- See [cmd/db/migrate](./cmd/db)


## Deployment(WIP)
