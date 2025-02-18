<div id="top"></div>

## Siosabin-go &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;![MySQL](https://img.shields.io/badge/-MySQL-00ADD8.svg?logo=mysql&style=for-the-badge)![Docker](https://img.shields.io/badge/-Docker-1488C6.svg?logo=docker&style=for-the-badge)![Python](https://img.shields.io/badge/-Python-F2C63C.svg?logo=python&style=for-the-badge)![Go](https://img.shields.io/badge/-Go-00ADD8.svg?logo=go&style=for-the-badge)

Siosabin-go is a refactor of the backend used in Siosabin, rewritten in Go.<br>
This project uses Docker to set up three containers: MySQL, Python, and Go.<br>
The original repository can be found [here](https://github.com/assam5649/Siosabin).

## Getting Started

### Installing

```
git clone https://github.com/assam5649/Siosabin-go.git
```

### Starting

Install Docker and compose at the root of the project.

```
docker-compose up
```

## API

Gin-based API

- /auth/<br>
    - GET /auth<br>
    Perform a health check for the auth endpoint.<br><br>
    - POST /auth/register<br>
    Register a user.<br><br>
    - POST /auth/login<br>
    Log in a user.<br><br>

- /main/<br>
    - GET /main<br>
    Perform a health check for the main endpoint.<br><br>
    - POST /main/data<br>
    Send data recorded by the device to the server.<br><br>
    - GET /main/users/<DeviceID><br>
    Retrieve the most recent record for the specified DeviceID.<br><br>
    - GET /main/location<br>
    Retrieve location information based on DeviceID.<br><br>
    - GET /main/salinity<br>
    Retrieve salinity information based on location.<br><br>