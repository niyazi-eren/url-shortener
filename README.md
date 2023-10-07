## <p align="center"> URL Shortener </p>

<p style="text-align: center"> implementation in Go of https://codingchallenges.fyi/challenges/challenge-url-shortener/ </p>

<p style="text-align: justify">This is a web application that allows users to (shorten) long URLs, making them easier to share and manage.</p>

### Stack
---
<b>Backend</b>: Go, Chi (router), Redis lite (own implementation: [Redis Lite](https://github.com/niyazi-eren/coding-challenges/tree/master/redis_server))

<b>Frontend</b>: Svelte, TypeScript, Tailwind CSS

<b>CI/CD / Deployment</b>: GitHub Actions, Terraform, AWS EC2

### Pre-requisite
- Docker
- Docker Compose

### Installation
---
Clone the repository: ```git clone https://github.com/niyazi-eren/url-shortener```

```cd url-shortener```

```docker-compose up -d```

Teardown with: ```docker-compose down```

### Usage

Access the application at http://localhost:5000.
