# <img alt="short-url" src="https://github.com/gyaan/short-urls/blob/develop/assets/short_url.png" width="220" />
Short urls create tiny url of a long urls and keep track of clicks. short urls required when someone want to send urls in messages or emails or keep it short to remember it. short url reduced the text of mobile messages. 

## Implementation Details
Short urls are generated using a algorithm you can find more details about here. After generating the short url, short url can be used to open the actual url its actually find out the actual url, increments the clicks and then redirect to actual url.  


## Tech stack
This project build on top of golang as primary programming language, mongodb as database and reactjs for front end application. below is list of the versions required to use this application.
 - Go >= v1.13
 - MongoDB >= v4.2.1
 - Reactjs >= 16.12.0

## Installation and building application
You can run apis (backend application) and frontend application separately. backend can be run as golang project and front end as ReactJs application but this required golang, mongodb and node setup in you local machine.
This project can be build using docker you can run below command to run project without installing any of tech in you local machine (off course you need docker installation in you local machine)

```docker-compose up -file= short-urls/build/package/docker-compose.yml```

## What you can learn form this project?
If you are looking to implement following things, this project can be very help to you
 - Project structure for golang application.
 - Mongo official golang driver.
 - Building multiple application with docker-compose.
 - Golang Chi package for routing.
 - Golang Viper package for managing config file.