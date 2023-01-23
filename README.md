<p align="center">
<br />
<h1 align="center">NFTir</h1>
<h5 align="center">SUNY OSwego </h3>
<h6 align="center">CSC 482 <h4>
</p>

###### Currently work in progress

## Overview

NFTir/server is a Golang-based application server, boasts a plethora of RESTful endpoints that facilitate the efficient processing and management of metadata and information pertaining to specific NFTs through the utilization of the NFTGo API stored in an AWS DynamoDB table by [NFTir poller agent](https://github.com/NFTir/agent)



## Getting Started

### Requirement

- [git](https://git-scm.com/)
- [golang](https://go.dev/)
- [docker](https://www.docker.com/)
- [aws-cli](https://aws.amazon.com/cli/)

### Clone the repo

```bash
git clone https://github.com/NFTir/server.git
cd server
```

### Set up environment variables

At the root of the directory, create a .env file using .env.example as the template and fill out the variables.

### Running the project

1.1 Build and run `agent` locally using `Make` scripts

```bash
make go-build-local
```

1.2 Run `agent` locally with hot-reload using `Make` scripts

```bash
make dev-mode
```

2. Build and run `agent` on Docker using `Make` scripts

```bash
make build-app
```

### Resources

- [NFTGo.io](https://nftgo.io/)
- [Amazon DynamoDB](https://aws.amazon.com/dynamodb/)
- [AWS-CLI](https://aws.amazon.com/cli/)
- [Golang](https://go.dev)
