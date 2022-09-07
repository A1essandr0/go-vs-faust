# go-vs-faust
Performance under load for basic Go and Faust applications

Running on local machine got me something like this:
py-sender (FastAPI, Faust, 100 locust users):  58% CPU usage, 0.4% memory usage
go-sender (FastAPI, goka,  100 locust users): 6.5% CPU usage, 0.1% memory usage
