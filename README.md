#Authorizer

Authorizer is an application designed and developed to authorize credit-card transactions, based on  a number of business rules, such as:
  
  - Credit-card available limit
  - Status of credit-card (active of inactive)
  - Duplicate transactions from same merchant
  - Multiple transactions in a short period of time
  
  
###Running

To run this application you should have go 1.13+ installed or docker

####With docker

Build the application

````make docker-build````

Run the application

```make docker-run < scenarios/{testFile}.txt ```

####Locally with golang

Install dependencies

````make deps````

Build the application

```make build```

Run the application

```./authorize < scenarios/{testFile}.txt```

The scenarios folder has a number of possible scenarios to be processed against the application

###Tests

The application has unit tests that are responsible for testing every layer of the application, such as processor, usecase and rules, the business rules have 100% of coverage of tests,meaning that all the scenarios are covered by unit tests.

####Running tests
To run the tests you must have golang 1.13+ installed on your computer

Running tests

`````make test-cov`````

The tests will be ran and two files will be generated, one called cover.txt and one called cover.html, you can open the cover.html on your browser to check the coverage of the code by the unit tests.