# DevOps Fizzbuzz

- [DevOps Fizzbuzz](#devops-fizzbuzz)
  - [Notes/context/pre-requisites](#notescontextpre-requisites)
  - [Task 1: Architectural/lifecycle improvements](#task-1-architecturallifecycle-improvements)
  - [Task 2: Observability improvements](#task-2-observability-improvements)
  - [Task 3: Configurability improvements](#task-3-configurability-improvements)
  - [Task 4: Containerisation](#task-4-containerisation)
  - [Task 5: Orchestration for local use](#task-5-orchestration-for-local-use)
  - [Task 6: Orchestration for deployment](#task-6-orchestration-for-deployment)
  - [Task 7: Automation implementation](#task-7-automation-implementation)
  - [Task 8: Documentation](#task-8-documentation)
- [FAQs](#faqs)


This repository contains two code files, `main.go` and `index.js`. They contain code that should be doing the exact same thing:

1. A server to listen for incoming connections
2. Some HTTP endpoint handlers to simulate failures (`/stubbed-process-X`)
3. An endpoint to save data that returns a UUID `/save`
   
   > Test this with `curl --data '{"hello":"world"}' -X POST 'http://localhost:3000/save`
4. An endpoint to load data using a UUID `/load/:uuid`
   
   > Test this with `curl 'http://localhost:3000/load/${UUID_FROM_SAVE}`
5. A healthcheck routine that performs a `curl` to another service to check if it's alive

## Notes/context/pre-requisites

1. You will need the following tools installed:
   1. Runtime
      1. [Go](https://go.dev/doc/install) if you're attempting this in Go
      2. [Node](https://nodejs.org/en/download/) if you're attempting this in Node
   2. [Docker](https://docs.docker.com/get-docker/)
   3. [Kubectl](https://kubernetes.io/docs/tasks/tools/)
   4. [Helm](https://helm.sh/docs/intro/install/)
   5. [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
2. The code contains a service that is meant for deployment into a Kubernetes cluster
3. You may ignore error handling unless it's an intentional design decision
4. For all tasks, functionality in terms of endpoint behaviour must remain the same
5. Assessment of this task will be a combination of the quality of the submission and a conversation over design decisions made during the implementation
6. If you are completely unfamiliar with Go and it would be faster to attempt this in another language, you may rewrite it into any language of your choice but the functionality must be the same

## Task 1: Architectural/lifecycle improvements

To the best of your current knowledge, refactor the code so that it aligns with best practices in software design and code lifecycle.

## Task 2: Observability improvements

To the best of your current knowledge, instrument the code so that it becomes observable. You may assume use of any monitoring tools you are familiar with.

## Task 3: Configurability improvements

To the best of your current knowledge, refactor the code so that the service becomes configurable for a cloud-native deployment.

## Task 4: Containerisation

Implement a Docker image for this service using any best practices in containerisation you know of.

## Task 5: Orchestration for local use

Create a Docker Compose manifest such that 3 instances of this service is spun up with each service performing the healthcheck on the next service. For example, `A` should check if `B` is alive, `B` should check if `C` is alive, `C` should check if `A` is alive.

## Task 6: Orchestration for deployment

Create a Helm chart that deploys the previous workload orchestration into Kubernetes. You can spin up a Kubernetes cluster locally using [the `kind` CLI tool](https://kind.sigs.k8s.io/):

```
kind create cluster;
```

## Task 7: Automation implementation

Developers may not always be familiar with build processes. Implement a way (eg. Makefile, shell scripts, etc) for developers to easily build and run this service locally.

## Task 8: Documentation

Replace the content of this `README.md` with relevant documentation on how to build/run/test/deploy/release this service.

# FAQs

**How long should I spend on these tasks?**

These tasks are based on day-to-day system administration tasks we do at watchTowr and are meant to be relatively simple (not complex) for a mid-senior level engineer. We expect this will take about 6 hours at maximum but there is no timeframe that you should rush this out within.

**Am I allowed to add more files?**

Yes

**Am I allowed to add more functionality/endpoints?**

Yes, but avoid deleting existing functionality.

**Can I rewrite the application to _lang?**

Yes, if you are not familiar with Go/JS, feel free to rewrite the functionality in another language. We are looking for outcomes and depth of understanding over semantics.

**Can I use any tooling I am familiar with?**

Asides from Docker and Kubernetes, yes. We are generally looking for understanding of concepts and experience handling real world issues instead of hard technical knowledge which can easily be picked up.

**What happens after my submission?**

If it looks fine, we will ping you for a follow-up call within a calendar week to:

1. Discuss design decisions made in your submission
2. Possibly modify your submission
3. A whiteboard exercise on infrastructure/architecture topics
