# go-mqtt-server
go server get data from mqtt broker and save it to mongo db


Require:

MongoDB server version: 4.2.6
go version go1.14.2 linux/amd64


Prepare a mqtt-broker and mongdb to test.


mqtt-server file is a execute binary file for linux/amd64:	
		mqtt-server -broker=[broker-link] -topic=[topic-name]
Example:
		./mqtt-server -broker=tcp://0.0.0.0:1883 -topic=sensors
		

		
		

others CMD:
	
	go test : check Go project  (Run it at first time to get all dependencies)
	go build: build execute file
	go run  : run without execute file  (Example:  go run main.go -broker=tcp://0.0.0.0:1883 -topic=sensors)
	

