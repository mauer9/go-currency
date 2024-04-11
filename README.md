![Go Report](https://goreportcard.com/badge/github.com/mauer9/go-currency)
![Repository Top Language](https://img.shields.io/github/languages/top/mauer9/go-currency)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mauer9/go-currency)
![Github Repository Size](https://img.shields.io/github/repo-size/mauer9/go-currency)
![Github Open Issues](https://img.shields.io/github/issues/mauer9/go-currency)
![License](https://img.shields.io/badge/license-MIT-green)
![GitHub last commit](https://img.shields.io/github/last-commit/mauer9/go-currency)
![GitHub contributors](https://img.shields.io/github/contributors/mauer9/go-currency)
![Simply the best ;)](https://img.shields.io/badge/simply-the%20best%20%3B%29-orange)


<p align="center">
<img align="center" height="200px" src="./images/bg.png">
</p>

# Go Currency service :money_with_wings: 📈

## Task description

We need to create a Golang & gorilla/mux web service that will fetch data from a national bank's public API on demand and save the data to a local TEST database

Constraints:

- 🚧 the service must listen to the port specified in config.json
- 🚧 Connection (Connection String) to MS SQL Server must be stored in config.json
- 🚧 2 GET methods must be implemented in the service:

🚨1st GET method: API .../currency/save/{date}
the date parameter must be passed to the national bank's API

After the necessary data is obtained, it should be saved to the R_CURRENCY table.

Saving should be done asynchronously, in the goroutine. Give the response to the user without waiting for completion. If an error occurs during saving, record it in the logs.
A successful response of this method is a JSON object with the success = true field.
Unsuccessful response of this method is to come up with it yourself.

🚨2nd GET method: .../currency/{date}/{*code} takes two parameters date and code, where code is an optional parameter.

The method needs to pull data with the specified parameters from the database and return it to the user.
The successful response of this method is an array of data in JSON format
Unsuccessful response of this method is to come up with it yourself (you can't show the user the real error from the database)

## Solution notes

- :trident: clean architecture (handler->service->repository)
- :cd: Makefile included
- :card_file_box: MSSQL migrations included

## HOWTO

- run with `make run`
