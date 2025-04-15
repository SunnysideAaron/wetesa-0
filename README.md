# Wetesa-0

An example CRUD api. Uses Go as the language and PostgreSQL as a datastore.
Dependencies are kept to a minimum. pgx is the only dependency and only because
the standard library does not include a sql driver. The decisions going into making
this example are documented in docs\Decision Records.

Since Go 1.22 (2024-FEB) many recommend to not use a framework. Using the standard
library instead. Most frameworks in Go were developed before Go 1.22 added better
routing.

I found myself unable to find good and complete working examples of how to use
the standard library to build an api. Specifically around routing. Built the
example I wanted! Leaned heavily on the information from
[How I write HTTP services in Go after 13 years](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/)
by Mat Ryer

Wetesa-0 is not a framework! It's just a fully working example of how to use the
standard library to build an api. Is this how I would build an api? Nope. There
are too many great packages out there to simplify the code. See 
TSDR-009 Non-dependencies.md for more information. I will be using the 
routing from this example and using the rest of the code as a reference when
evaluating what value packages bring. I personally believe the decision records
are also valuable. Any api has to answer the same questions. Going through them
and finding your own answers is a good way to start your own api.

## Requirements

- [Docker](https://www.docker.com/)

## Usage

### Quick start

- Install [Docker](https://www.docker.com/)
- Install Make
    - There are several different ways to get Make installed. Google for your
      operating system.
- Clone the Wetesa-0 repo to your computer
- In your command prompt change your working directory to where you cloned Wetesa-0
- This example uses 2 docker containers. "datastore" for PostreSQL and "api" for
  the api. The following command should bring them both up.

      make all-up

- http://localhost:8080/ and http://localhost:8080/health Should both give responses.
- make all-up will launch air on the api code. Air rebuilds the binary when ever
  it detects code changes.
- To end the running services press ctrl-c. Then

      make all-down

### Detailed usage

Additional make commands are provided for running the containers independently. 
This could be helpful depending on what one is trying to accomplish. For example:

in command prompt 1:

    make ds-up

in command prompt 2:

    make api-bash

Will launch the datastore service and leave you on the command line in the api service.
Once inside the api service you can run go commands on the code as needed. For example:

    make watch

will get you where make all-up got you.

Note there are two Make files. See ./Makefile and ./api/Makefile for additional commands.

## Credits

This project borrows heavily from many sources. While it has traveled a bit from them we would like to thank them. Please check them out.

- [How I write HTTP services in Go after 13 years](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/) by Mat Ryer
- [Go Blueprint Code](https://github.com/Melkeydev/go-blueprint)
- [Go Blueprint Web](https://go-blueprint.dev/)

