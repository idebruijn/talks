Welkom everybody...

Today we want to talk to you about Forest and how an why we created this package for testing REST based Webservices.
Forest is a package with a few simple helpers which help you creating functional tests that call a REST based Webservice.

So why did we start making this?

# Tools we used
I have been testing other services at bol.com and other companies and here I used different tools like FitNesse, SOAPUI, Jmeter. But all these tests didn't provide me the thinks I wanted. FitNesse was to slow, not easy to refactor and error handeling wasn't any good. SOAPUI was fast but creating and maintaining in the GUI wasn't what I liked and you couldn't really word together on test because the tests are in one huge file. Same things apply for Jmeter. 

#speed
I saw Ernest using 'go test' for his unit tests and they where so fast I wanted that to but I didn't know well how to program or Go. During a hackathon at bol.com I started writing my first tests.

The Go standard library has all you need to write these test, but you needed to write but the required amount of code, especially the error handling, makes tests harder to read and longer to write. And for me it was hard at first. 

Example: verbose

It should be easier to Setup Http Requests, send them and inspect Http Responses and verifying the response payload and a few lines. So Ernest started to write some helpers which made my tests more lean, easier to write and maintain.

Example: lean

Verifying the response is always the most important part of the test because it wil vallidate if it is correct or not.
The forest package offers several functions to inspect JSON and XML documents.
All functions operate on the *http.Response and take the *testing.T for reporting errors.

Example: ExpectStatus

But I wanted to reuse the *http.Response in my tests for other checks and that is also possible with Forest. Now I can use is to check status codes, error codes and other assertions in the response.

Because now I don't have a gui when running the test, I only use the terminal. But when you run lots of tests finding which one went wrong wasn't easy enough. So we added colorized logging when a test fails and when a rest failed we also Dump the response in the logging giving you more information for you analysis.

Example: error + dump

I did mis some human readable logging when I was reading my tests in the terminal and for services testing a like the Given-When-Then style. Ernest wasn't a fan because it ment more code. So this is not part of Forest but I did wanted to show you. This is not a BDD framework but just some nice colorized logging in Given-When-Then style. Now I can say we like it because it makes it easier to read the tests, exspecially when a test has many REST calls in it.

Example: terminal

I do like writing these tests now! We have created a large collection of fast, maintainable and readable tests and because of the speed of compilation and execution, we now run these API tests as frequently as we run the unit tests of our services. Which is really important now that we are going more and more to continues delivery.