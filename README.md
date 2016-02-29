# twitter-breaking-news [![Circle CI](https://circleci.com/gh/danesparza/twitter-breaking-news.svg?style=svg)](https://circleci.com/gh/danesparza/twitter-breaking-news)
Simple breaking news microservice written in Go

*To build, make sure you have the latest version of [Go](http://golang.org/) installed.  If you've never used Go before, it's a quick install and [there are installers for multiple platforms](http://golang.org/doc/install), including Windows, Linux and OSX.*

### Quick Start

Run the following commands get latest and build.

```bash
go get github.com/danesparza/twitter-breaking-news
go build
```

A docker image will be shipped soon

### Starting and testing the service
To start the service, just run `twitter-breaking-news`.  

If you need help, just run `twitter-breaking-news --help`.

There are a few command line parameters available:

Parameter       | Description
----------      | -----------
port            | The port the service listens on.  
allowedOrigins  | comma seperated list of [CORS](http://en.wikipedia.org/wiki/Cross-origin_resource_sharing) origins to allow.  In order to access the service directly from a javascript application, you'll need to specify the origin you'll be running the javascript site on.  For example: http://www.myjavascriptapplication.com
consumerKey     | The consumer key (API key).  Get this from your [Twitter app dashboard](https://apps.twitter.com/).  
consumerSecret  | The consumer secret (API secret).  Get this from your [Twitter app dashboard](https://apps.twitter.com/).  
authToken       | The auth token (Access Token).  Get this from your [Twitter app dashboard](https://apps.twitter.com/).  
authSecret      | The auth secret (Access Token Secret).  Get this from your [Twitter app dashboard](https://apps.twitter.com/).  

Once the service is up and running, you can connect to it using
`http://yourhostname:3000/news/screenname` where `screenname` is the twitter screen name you're using to get breaking news.  

Example: `http://yourdomain.com:3000/news/cnnbrk`

To test your service quickly, you can use the [Postman Google Chrome Extension](https://chrome.google.com/webstore/detail/postman-rest-client/fdmmgilgnpjigdojojpjoooidkmcomcm?hl=en) to call the service and see the JSON return format.

News information will be returned as a JSON payload.
