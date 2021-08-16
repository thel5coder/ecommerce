# Ecommerce Microservices

## Project Structure

- domain

In this domain directory, most of the interfaces are divided into several directories based on their respective responsibilities.

- commands

This directory contains interfaces for save, delete, modify operations in the database for each table or model.

- configs

The config directory contains interfaces and implementations of third-party libraries or links to infrastructure such as database connection interfaces, interfaces to temporary databases, validators, and etc.

- files

The files directory is used to store static files related to static assets.

- handlers

This directory contains interfaces for functions to accept requests and return responses.

- models

The models directory contains structures or properties that represent a table in the database and also contains a function implementation of the imodels interface which has two functions that are used to scan rows and rows.

- queries

This directory contains interfaces for select operations in the database for each table or model.

- requests

This directory contains the request structure used to receive the request body from the http request process

- usecase

The use case directory contains interfaces of functions that manage logic and business logic

- view\_models

This directory contains the structs that will be sent as the response body.

- migrations

This directory contains the sql commands that migrations will run.

- repository

This directory contains database operations functions such as select, insert, update, delete. However, these functions are divided into two directories, namely commands and queries.

- commands

The command directory is an implementation of the interface in the command directory inside the domain directory which contains the operation functions of saving, modifying, and deleting data in the database.

- queries

Directory queries is an implementation of the interface that exists in the queries directory inside the domain directory which contains the select and select count operations data on a table in the database.

- server
This server directory will actually be divided into two http and grpc depending on the protocol used.
  
    - boot 
  The boot directory contains a struct file that contains elements such as config, app instances, and more. It also contains a routers directory and a routers file which functions to handle and register all routing addresses or urls.

      - routers 
      This directory contains routers files which are split by resources/modules/features which implement the irouters interface which has a function to register the url or routing address of a particular module or feature.

    - handlers
    This directory is an implementation of the handlers interface on the handlers directory in the domain directory. These functions are responsible for accepting requests and returning responses.

    - middlewares
    This directory contains a collection of middleware functions

- usecases

This directory is an implementation of the interface in the usecases directory in the domain directory, which contains a collection of logical functions and business logic of a feature or business flow.

## How To Run

1. On your system you must have docker and docker-compose installed.
2. You have to create a docker network with the name ecommerce-net.

```$ docker network create ecommerce-net```

1. Prepare the postgresql and redis database, if you don&#39;t have one, you can use the docker I have prepared inside infra directory, just run the docker-compose.yml file.

But if you already have postgresql and redis installed on your machine, you can skip this process, to run the docker compose for postgresql and redis you can use the command below.

```$ docker-compose up -d --build```

1. After your postgresql is running create 3 databases with the names product-service, transaction-service, and user service.
2. After you create 3 databases import data from files that are already available in the root directory of the project.
3. To run all services, all you have to do is run the docker-compose that has been provided.

```$ docker-compose up -d --build```

1. To turn off services, you just have to turn off all services with the docker-compose down command.

```$ docker-compose down```

1. You can use Postman to try each service. Use this link [https://documenter.getpostman.com/view/1502395/Tzz8qwbN](https://documenter.getpostman.com/view/1502395/Tzz8qwbN) to open postman in the browser and then you can import from the browser to the postman desktop.
2. You can login using email=[superadmin@gmail.com](mailto:superadmin@gmail.com) and password=123456789 as admin.