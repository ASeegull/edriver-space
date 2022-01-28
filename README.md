# eDriver Space

[![Lint](https://github.com/ASeegull/edriver-space/actions/workflows/lint.yml/badge.svg)](https://github.com/ASeegull/edriver-space/actions/workflows/lint.yml)
[![Unit tests](https://github.com/ASeegull/edriver-space/actions/workflows/unit-tests.yml/badge.svg)](https://github.com/ASeegull/edriver-space/actions/workflows/unit-tests.yml)
[![Go Report Card](https://goreportcard.com/badge/ASeegull/edriver-space)](https://goreportcard.com/report/ASeegull/edriver-space)
[![Coverage Status](https://coveralls.io/repos/github/ASeegull/edriver-space/badge.svg?branch=main)](https://coveralls.io/github/ASeegull/edriver-space?branch=main)
![Edit diagram](docs/images/e-driver.drawio.svg)

### Recommended workspace layout

Please use [layered architecture](https://www.oreilly.com/library/view/software-architecture-patterns/9781491971437/ch01.html)
The main point is that layers should be isolated from one another and changes in one layer should affect this and possibly another associated layer.

```
├─ src (or however you'll name your root directory)
│   ├─ edriver-space (root of current repository)
│         ├─ docs
│         ├─ cmd
│             ├─ app/ main.go (<-- program entry point)
│         ├─ conf
│         ├─ handlers
│         ├─ migrations
│         ├─ model
│         ├─ strorage
│         ├─ parking
│         ├ docker-compose.yaml
│   ├─ edriver-space-webapp (root of [this](https://github.com/ASeegull/edriver-space-webapp) repository)
│   ├─ edriver-space-aggregator (root of [this](https://github.com/ASeegull/edriver-space-aggregator) repository)
```
