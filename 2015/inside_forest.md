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