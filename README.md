# gomicroservice

# direcotory structure 
/service1
    /api
        /controllers
        /dto
        /middlewares
        /routes
    /config
    /internal
        /domain
            /models
            /repositories
        /usecases
    /pkg
    /cmd
        /service1
    /docs
    /scripts
    /tests
    /infra
        /db
        /cache
        /messaging
    Dockerfile
    go.mod
    go.sum
    main.go
# Directory Definition
    /service1
    /api : Handles HTTP communication.
        /controllers :  Contains controller logic for handling API requests.
        /dto : Data Transfer Objects (structs for incoming/outgoing requests).
        /middlewares : HTTP middlewares like authentication, logging, etc.
        /routes : Defines routes and routing logic.
    /config : Configuration files (e.g., environment variables, external service credentials).
    /internal : Internal logic specific to this service (cannot be imported by other services).
        /domain : Business domain logic and core models.
            /models : Structs for data models.
            /repositories : Interfaces and implementation for data storage.
        /usecases : Business use case handlers and service logic.
    /pkg : Shared packages within the service that are reusable.
    /cmd
        /service1 : Entry point for the service. This contains the main.go or equivalent file to start the service.
    /docs : API documentation (e.g., Swagger, Postman collections).
    /scripts : Scripts for tasks like deployment, migrations, etc.
    /tests : Unit tests, integration tests, etc.
    /infra : Infrastructure-related files.
        /db : Database connections, migrations.
        /cache : Cache connections (e.g., Redis).
        /messaging : Message queues (e.g., RabbitMQ, Kafka).
