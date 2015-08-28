In my part of the talk I will give a short overview of the api
and dive a bit deeper into the callenge of implementing this package.


## Prepare a HTTP request

testing a webservice is all about sending http requests and receiving http responses.

all of our services document their API using a Swagger interface

for each webservice call, it is specified what the URL is, what method and which parameters.

so building a http request is where you typically start.


forest has a simple RequestConfig struct that can be used to build a HTTP request

is has all kinds of convenience methods for this task.



## Send the HTTP request


for actually sendig a http request we use small wrapper on the standardt http client

this wrapper has convenience methods to send requests and inspecting all possible errors

this wrapper can be reused in all tests so we declare it as a package var.

based on the configuration we build the actual request and send using the api client.



## Inspect HTTP response

after checking the status code, it is time to verify the content of the response.

depending on your api implementation, the content type can be textual or binary.

for the common cases that the api returns XML and JSON, we have some helper functions to check those.


## testing a testing package

it turned out to be a nice challenge to test this testing package.

in order to verify the calls to the standard testing package, I needed to create a mock implementation that will capture the arguments for later inspection.

for this i created an interface that declares all functions this package will need from testing.T.


## mock it

now i can define a mock version implementating exactly that interface in which actual messages are captured.


## testing against a testing service

because forest purpose is to these API actual requests and responses, i decided that as part of testing forest itself, i needed to run a local http server.

for that server, i can defined exactly for each request what the response will be.

and with that server, i can create a forest client to talk to the api  (see tsAPI)


## expectation test

so this example shows a simple test that ExpectStatus will correctly verifiy that the response status code is 404.

it looks like a test you would write for your own api.


## verifying output

using the same interface as we use to create a mock.
the logging version is implemented to capture whatever is send to Log,Error or Fatal


## TestMain

the go standard library provide a means for a global setup and teardown phase.

to support using forest in this context, we need a testing.T implementation.


