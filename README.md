# Wetesa

An example CRUD api. Uses Go as the language and PostgreSQL as a datastore. Has a minimum of dependencies.

Fork me, please! I know I will! :-)

## Requirements

- [Docker](https://www.docker.com/)

## Usage

### Quick start

- Install [Docker](https://www.docker.com/)
- Install Make
    - There are several different ways to get Make installed. Google for your opperating system.
- Clone the Wetesa repo to your computer
- In your command prompt change your working directory to where you cloned Wetesa
- This example uses 2 docker containers. "datastore" for PostreSQL and "api" for the api. The follwing command should bring them both up.

      make all-up

- http://localhost:8080/ and http://localhost:8080/health Should both give responses.
- make all-up will launch air on the api code. Air rebuilds the binary when ever it detects code changes.
- To end the running services press ctrl-c. Then

      make all-down

### Detailed usage

Additional make commands are provided for running the containers independently. This could be helpful depending on what one is trying to accomplish. For example:

in command prompt 1:

    make ds-up

in command prompt 2:

    make api-bash

Will launch the datastore service and leave you on the command line in the api service. Once inside the api service you can run go commands on the code as needed.

See ./Makefile and ./api/Makefile for additional commands.

## Notes

Much of the original code was generated using [Go Blueprint](https://github.com/Melkeydev/go-blueprint). Thank you Go Blueprint team!
