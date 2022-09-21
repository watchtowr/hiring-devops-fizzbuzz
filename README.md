# DevOps Fizzbuzz

Observe that the `main.go` file contains a server as well as a healthcheck routine.

## Notes/context/pre-requisites

1. You will need the following tools installed:
   1. [Go](https://go.dev/doc/install)
   2. [Docker](https://docs.docker.com/get-docker/)
   3. [Kubectl](https://kubernetes.io/docs/tasks/tools/)
   4. [Helm](https://helm.sh/docs/intro/install/)
   5. [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
2. The code contains a service that is meant for deployment into a Kubernetes cluster
3. You may ignore error handling unless it's deemed architecturally significant (detailed error-handling can be very verbose in Go)
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
