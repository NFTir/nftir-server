<p align="center">
<br />
<h1 align="center">NFTir - Agent</h1>
</p>

## Overview

NFTir agent is a golang server which polls the NFTGo server every 6 hours to get a dataset of the top 25 NFTs ranking based on their trading volumes.
The agent then breaks down the dataset into single items and push them into an Amazon DynamoDB table which will be used by the [NFTir RESTful server](https://github.com/logann131/NFTir/tree/master/server) to process the data.

## Getting Started

### Requirement

- [git](https://git-scm.com/)
- [golang](https://go.dev/)
- [docker](https://www.docker.com/)
- [aws-cli](https://aws.amazon.com/cli/)

### Clone the repo

```bash
git clone https://github.com/logann131/NFTir.git
cd NFTir/agent
```

### Set up environment variables

At the root of the directory, create a .env file using .env.example as the template and fill out the variables.

### Running the project

```bash
go run .
```

or

```bash
docker compose up --build
```

### Resources

- [NFTGo.io](https://nftgo.io/)
- [Amazon DynamoDB](https://aws.amazon.com/dynamodb/)
- [AWS-CLI](https://aws.amazon.com/cli/)
- [Golang](https://go.dev)
