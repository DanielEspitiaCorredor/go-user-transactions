# go-user-transactions
Project used to process debit and credit transactions. This project is an adaptation in Golang to the project https://github.com/DanielEspitiaCorredor/user-transactions



# Components



The go-user-transactions project is a REST API built using the Gin framework and organized into several modules.

## Module Overview
**Core Module:** This includes the main.go file and the route definitions found in internal/routes, where all project routes are specified.

**Middleware Module:** This module features two middlewares: one for managing CORS and another for authentication via API key. The API key must be defined in the .env file.

**ODM Module:** This module is responsible for managing the connection to MongoDB and adapting Go objects for compatibility with the database.

**Transaction Logic Module:** The largest module in the project, it handles the processing of files and sends the final report via email.

## Libraries and Assets
For data manipulation, the project utilizes the go-gota library, while the email functionality is handled with gomail.

The assets folder contains the CSV file used in the project, along with the email template for communication.

To facilitate email sending, the AccountBalance struct integrates relevant information into the HTML template and manages the delivery process.


# Prerequisites

- Docker
- Docker compose


# Get Started


For configure this project, please set the env file in the project path

![alt text](/assets/readme/setenv.png)


This project use a docker-compose file to launch two containers. The first is the gin application and the other is a mongo database.

```
docker compose up -d --build
```


Once the containers are running, you can execute the endpoint to generate the report. This service processes a CSV file and sends the email. To execute this code, use the following command:

```
make generate-report EMAIL=<example@example.com>
```

In the `EMAIL` param, set the email that receive the report 


# Generate data

To generate a new file with test data, please download the python version and continue with the instructions [docs](https://github.com/DanielEspitiaCorredor/user-transactions?tab=readme-ov-file#generate-data) 