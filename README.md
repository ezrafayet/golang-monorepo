# golang-monorepo

Golang-monorepo is a basic monorepo for self-education, featuring Go services and shared libraries. It doesn't use a monorepo "framework" (NX or Bazel). It was proven to be good enough for small to medium projects.

# Project Structure Overview

```
.
├── apps               # Each service is housed in its own directory here.
│   ├── service-a      # A standalone service
│   └── service-b      # Another standalone service
├── libs               # This is where shared libraries are kept.
│   └── golang
│       └── lib-a      # An example of a shared library
├── docker-compose.yml # Defines the services for local development
└── Makefile           # Facilitates build and run commands
```

## Potential Enhancements

- Introducing a containers directory could help organize Dockerfiles more efficiently, especially for dependencies like databases or message brokers (e.g., Kafka, RabbitMQ).
- A pipeline directory for Continuous Integration and Continuous Deployment (CI/CD) would streamline development workflows.
- Adding a doc folder for comprehensive documentation could significantly aid in project maintainability and onboard new contributors more smoothly.

# Operational Mechanics

Each service manages its dependencies via its own go.mod file, employing the replace directive for referencing shared libraries. This setup necessitates building the project outside of Docker and subsequently transferring the binaries into the Docker environment for execution.

## Considered Alternatives

- An alternative approach involves vendoring the libraries, including shared ones, and performing the build process within the Docker environment.
- Adopting a monorepo management framework such as NX or Bazel could introduce more structured dependency management and build optimization.
- Leveraging Go's workspace feature might offer some benefits.

# Running the services

From the root, run:

```bash
make dev
```

From the browser navigate to:
```
- localhost:5050/api/service-a

and

- localhost:5050/api/service-b
```

# What is missing for production
