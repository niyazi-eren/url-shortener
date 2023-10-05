## <p align="center"> URL Shortener </p>

<p style="text-align: center"> implementation in Go of https://codingchallenges.fyi/challenges/challenge-url-shortener/ </p>

<p style="text-align: justify">This is a web application that allows users to (shorten) long URLs, making them easier to share and manage.
In its current state though I use ec2 public DNS which makes it more of a url extender.</p>

### Stack
---
<b>Backend</b>: Go, Chi (router), Redis lite (own implementation: [Redis Lite](https://github.com/niyazi-eren/coding-challenges/tree/master/redis_server))

<b>Frontend</b>: Svelte, TypeScript, Tailwind CSS

<b>CI/CD / Deployment</b>: GitHub Actions, Terraform, AWS EC2
