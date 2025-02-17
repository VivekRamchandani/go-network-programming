# Network Programming with Go
Creating Network Programs using Golang.


## TCP
TCP provides a reliable connection-oriented protocol over IP.
The [tcp](tcp) folder contains all the recepies for working with TCP Sockets using Golang.
The [simpleServer](tcp/simpleServer.go) recieves messages from [simpleClient](tcp/simpleClient.go) untill "STOP" is sent.

## UDP
UDP is a conectionless protocol where no 'session' is established.
The [udp](udp) folder contains all the recepies for working with UDP Sockets using Golang.
The [simpleServer](udp/simpleServer.go) recieves message from [simpleClient](udp/simpleClient.go).