# Golang-monorepo

Golang-monorepo is a basic monorepo for self-education, featuring Go services and shared libraries. It doesn't use a monorepo "framework" (NX or Bazel). It was proven to be good enough for small to medium projects.

# Service Execution

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

# Project Architecture

The project contains a proxy and two services, sharing a single library.

```
                            +------------+
                        +-->|  Service A |
                        |   +------------+
        +---------+     |        :8000 
        |  Proxy  +-----+  
        +---------+     | 
           :5000        |  
                        |   +------------+
                        +-->|  Service B |
                            +------------+
                                 :8010
```

# Project Structure

```
.
├── apps <----------- This is where services are housed in their own directory.
│   ├── service-a
│   ├── service-b
│   └── proxy-service
├── libs <----------- This is where shared libraries are kept.
│   └── golang
│       └── lib-a
├── docker-compose.yml
└── Makefile
```

# Potential Additions

- Introducing a `containers` directory for databases or message brokers.
- Introducing a `tests` folder for end-to-end tests.
- Introducing a `pipeline` directory for CI/CD.
- Introducing a `doc` folder for documentation.
- Introducing a `scripts` folder for utility scripts.
- Introducing a `contracts` folder for shared interfaces or gRcp definitions.

# Operational Mechanics

Each service manages its dependencies via its own go.mod file, employing the replace directive for referencing shared libraries. This setup necessitates building the project outside of Docker and subsequently transferring the binaries into the Docker environment for execution.

# Alternative approaches

- An alternative approach involves vendoring the libraries, including shared ones, and performing the build process within the Docker environment.
- Adopting a monorepo management framework such as NX or Bazel could introduce more structured dependency management and build optimization.
- Leveraging Go's workspace feature might offer some benefits.
