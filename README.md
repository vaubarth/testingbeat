# Testingbeat

Testingbeat allows reading test result files (e.g. Junit xml files) and send them to elasticsearch.

## Using Testingbeat

tbd


## Building Testingbeat

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/vaubarth/testingbeat`

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Init Project
To get running with Testingbeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Testingbeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/vaubarth/testingbeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Testingbeat run the command below. This will generate a binary
in the same directory with the name testingbeat.

```
make
```


### Run

To run Testingbeat with debugging output enabled, run:

```
./testingbeat -c testingbeat.yml -e -d "*"
```


### Test

To test Testingbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```


### Cleanup

To clean  Testingbeat source code, run the following command:

```
make fmt
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Testingbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/vaubarth/testingbeat
git clone https://github.com/vaubarth/testingbeat ${GOPATH}/src/github.com/vaubarth/testingbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make release
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.
