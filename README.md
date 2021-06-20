# Roulette Service

Roulette service provides a REST API for a roulette game. 

The service consists of 2 main parts;

* Tables - Tables are the roulette tables and is the foundation of any game. A bet can only be placed on a table, and a game can only be played on a table. 
* Bets - Bets belong to a table and are an individual bet for the current round.

## Prerequisites

* Install Go [v1.16](https://golang.org/dl/)

* Install [Docker](https://www.docker.com/get-started)

## Usage

To run the application locally

```shell
make serve
```

Go to the Open API documentation for API usage

http://localhost:8080/v1/docs

## Test and Coverage

```shell
make test
make cover
```
