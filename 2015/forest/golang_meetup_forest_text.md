Welkom everybody...

Today we want to talk to you about Forest and how an why we created this package for testing REST based Webservices.
Forest is a package with a few simple helpers which help you creating functional tests that call a REST based Webservice.

So why did we start making this?

# Tools we used
I have been testing other services at bol.com and other companies and here I used different tools like FitNesse, SOAPUI, Jmeter. But all these tests didn't provide me the thinks I wanted. FitNesse was to slow, not easy to refactor and error handeling wasn't any good. SOAPUI was fast but creating and maintaining in the GUI wasn't what I liked and you couldn't really word together on test because the tests are in one huge file. Same things apply for Jmeter. 

# First test written in Go
I saw Ernest using 'go test' for his unit tests and they where so fast I wanted that to but I didn't know well how to program or Go. During a hackathon at bol.com I started writing my first tests.
The Go standard library has all you need to write these test, but you needed to write but the required amount of code, especially the error handling, makes tests harder to read and longer to write. And for me it was hard at first. 

# boilerplate

# Lean test with Forest
So whe thought it should be easier to Setup Http Requests, send them and inspect Http Responses and verifying the response payload and a few lines. So Ernest started to write some helpers which made my tests more lean, easier to write and maintain.

Example: lean

# Real testcase DO A DEMO!!
So here is a real testcase which creates a topic on our service BoQs and expects a http status back in the response.

# Verifying response
Verifying the response is always the most important part of the test because it will validate if it is correct or not.
The forest package offers several functions to inspect JSON and XML documents.
All functions operate on the *http.Response and take the *testing.T for reporting errors.

Example: JSONpath

# Analyse failing tests
When a test fails you want all the information so you can make a fast analysis why a test failed. For test this usually starts with the http response.
Forest make the response reuseable so we can DUMP the message when a test fails. But I can reuse the *http.Response in my tests for other checks. Now I can use is to check status codes, error codes and other assertions in the response.
Because now I don't have a GUI when running the tests, I only use the terminal. But when you run lots of tests finding which one went wrong wasn't easy enough. So we added colorized logging when a test fails.

# Given-When-Then DO A DEMO!!
I did that my tests where not human readable when I was reading my tests in the terminal and for services testing a like the Given-When-Then style. Ernest wasn't a fan because it ment more code, but I do want to use it in my tests. This is not a BDD framework but just some nice colorized logging in Given-When-Then style. After using it I can say we like it because it makes it easier to read the tests, specially when a test has many REST calls in it. So lets have a DEMO!

Example terminal: bdd + colorized error + dump


I like writing these tests now! We have created a large collection of fast, maintainable and readable tests and because of the speed of compilation and execution, we now run these API tests as frequently as we run the unit tests of our services. Which is really important now that we are going more and more to continues delivery.