# go-vs-faust
Performance under load for basic Go and Faust applications.

Comparing senders:   
py-sender (FastAPI, Faust, 100 locust users):  58% CPU usage, 0.4% memory usage   
go-sender (FastAPI, goka,  100 locust users): 6.5% CPU usage, 0.1% memory usage   

Comparing collectors:   
go-collector (basic net/http) uses 3-4 times less CPU resources than fastapi-collector under the same load (go-sender, 100 locust users)

