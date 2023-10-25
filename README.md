README

This project implements a RESTful service that calculates the number of days left until January 1, 2025. Besides, it has a middleware which authenticates the request with a secret token.

To run the application, follow the steps below:

## Set up and Run the Application

1. Clone the repository and navigate to the project directory:

```
https://github.com/Samandar587/iman-task.git
cd iman-task
```

2. Run the application:

```
go run .
```

3. Send the request with the secret token, which is provided in the source code, to the following address:

```
localhost:8080/days-left
```