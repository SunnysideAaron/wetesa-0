# Wetesa-0

An example CRUD API. Uses Go as the language and PostgreSQL as a datastore.
Dependencies are kept to a minimum. pgx is the only dependency and only because
the standard library does not include a sql driver. The decisions going into making
this example are documented in
[docs\Decision Records](https://github.com/SunnysideAaron/wetesa-0/tree/main/docs/Decision%20Records).
For the TLDR see
[TSDR-000](https://github.com/SunnysideAaron/wetesa-0/blob/main/docs/Decision%20Records/Tech%20Stack/TSDR-000%20What%20and%20TLDR.md)
and [ADR-000](https://github.com/SunnysideAaron/wetesa-0/blob/main/docs/Decision%20Records/Architecture/ADR-000%20What%20and%20TLDR.md).

Since Go 1.22 (2024-FEB) many recommend using the standard library instead of a 
framework. Most frameworks in Go were developed before Go 1.22 added better
routing.

Found myself unable to find good, complete, and working examples of how to use
the standard library to build an API. Specifically, around routing. Built the
example I wanted! Leaned heavily on the information from
[How I write HTTP services in Go after 13 years](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/)
by Mat Ryer

Wetesa-0 is not a framework! It is a fully working example of how to use the
standard library to build an api. Is this how I would build an api? Possibly. If
the project was small enough or if I was very concerned about having too many
dependencies. [TSDR-008 Possible Future Dependencies.md](https://github.com/SunnysideAaron/wetesa-0/blob/main/docs/Decision%20Records/Tech%20Stack/TSDR-008%20Possible%20Future%20Dependencies.md) 
covers some ideas that might make sense to add / change depending on the project.

Ways to use Wetesa-0:
- Example code for api routing using the standard library.
- A base line for evaluating packages. How they would change code? What specific
  benefits do they bring.
- The decision records as a starting point for any new project. Any api / project has
to answer the same questions. Going through them and finding your own answers is
a good way to start a new api / project.

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

- http://localhost:8080/healthz, http://localhost:8080/healthdbz, and http://localhost:8080/api/v0.1/clients Should all give responses.
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

This project borrows heavily from many sources. While it has traveled a bit from
them we would like to thank them. Please check them out.

- [How I write HTTP services in Go after 13 years](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/) by Mat Ryer
- [Go Blueprint Code](https://github.com/Melkeydev/go-blueprint)
- [Go Blueprint Web](https://go-blueprint.dev/)

