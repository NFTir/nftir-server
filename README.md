<p align="center">
<br />
<h1 align="center">NFTir</h1>
<h5 align="center">SUNY OSwego </h3>
<h6 align="center">CSC 482 <h4>
</p>

###### Currently work in progress

## Overview

NFTir/server is a golang app server hosting multiple RESTful endpoints that helps processing metadata and information of specific NFTs stored in an AWS DynamoDB table

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

1. Build and run `agent` locally using `Make` scripts
```bash
make go-build-local
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
