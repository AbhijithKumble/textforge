# Install and Start Learning

## Prerequisites

Install:

- [Go 1.22 or newer](https://go.dev/dl/)
- Git

Check that they are available:

~~~bash
go version
git --version
~~~

## Clone the project

~~~bash
git clone https://github.com/AbhijithKumble/textforge.git
cd textforge
~~~

Download the Go dependencies:

~~~bash
go mod download
~~~

## Start learning

List the available lessons:

~~~bash
go run . learn
~~~

Read a specific lesson:

~~~bash
go run . learn regex-01-literals
~~~

Start the interactive practice session:

~~~bash
go run . practice
~~~

TextForge presents the exercises in order and saves each correct answer. Type
q or quit at any exercise to leave; your completed exercises remain saved.

## Useful commands

~~~bash
go run . progress
go run . practice regex-02-character-classes
go run . learn regex-03-quantifiers
go build -o textforge .
~~~

After building, run the executable from the project directory:

~~~bash
./textforge learn
./textforge practice
~~~

## Reset progress

Progress is saved in .textforge/progress.json. To start over:

~~~bash
rm .textforge/progress.json
~~~

You can use a separate progress file while experimenting:

~~~bash
go run . practice --progress /tmp/textforge-progress.json
~~~

## Run tests

~~~bash
go test ./...
~~~
