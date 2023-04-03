## Overview

This repo contains a tcp server implementing algorithm of DDOS attacks mitigation using proof-of-work.
The main idea of the algorithm is to require clients that wish to connect to a server 
perform some work on their computer and show this to the server.
Service is granted to client only after successful proof of work validation.

### Implementation

Clients communicate with server using [challenge-response protocol](https://en.wikipedia.org/wiki/Proof_of_work). 
Generally it consists of the following steps:
- client establishes tcp connection with the server
- client requests a challenge
- server chooses a challenge and sends it back to the client
- client solves the challenge and sends the proof to the server
- server grants service

### PoW algorithm
The choice of the PoW challenge should be based on the following factors:
- The only way to solve the challenge is to check every element in the set of possible solutions.
Simply put - bruteforce.
- Verifying solution should be many times less complex than solving the challenge in terms of time and memory
- It must be possible to fine-tune the difficulty level of the challenge. 
Increasing the difficulty level must lead to increasing the average time of its solution finding, 
but it must not affect the time of verification.
- When checking, it must be possible to verify that the task was generated by the server.
In other words, the challenge must be signed by the server
- The challenge must be stateless. It means that the server must be able to check the solution 
without remembering the challenge itself.
- It must be impossible to prepare a large number of solved tasks in advance 
and use them simultaneously to perform a DDOS attack.

One algorithm that meets all of the above criteria is the [Client Puzzle challenge](https://en.wikipedia.org/wiki/Client_Puzzle_Protocol).
It was introduced by Ari Juels and John Brainard in 1999 
in their [paper](http://www.arijuels.com/wp-content/uploads/2013/09/JB99.pdf).


### Project structure

- `cmd/client` - entry point of the client program
- `cmd/server` - entry point of the server program
- `internal/app/server` - implementation of the server. It includes tcp connections handling and messages routing
- `internal/app/client` - implementation of the client. After launch the client connects to the server, 
goes through all challenge-response protocol steps once and finally prints quote to the console and finishes. 
- `internal/pkg/contract` - common library containing messages of client-server communication contract
It also contains methods for serialization/deserialization of contract messages
- `internal/pkg/pow` - common library containing implementation of Client Puzzle PoW challenge

### How to run
```
$ docker-compose up --detach --build
$ docker-compose run client
```
