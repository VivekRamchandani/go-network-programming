# Network Programming with Go
Creating Network Programs using Golang.


## TCP
TCP provides a reliable connection-oriented protocol over IP.<br>
The [tcp](tcp) folder contains all the recepies for working with TCP Sockets using Golang.

**Recipie 1: Simple Client Server using TCP**

The [simpleServer](tcp/simpleServer.go) recieves messages from [simpleClient](tcp/simpleClient.go) untill "STOP" is sent.

**Recipie 2: Simple file Sender and Reciever**

The [fileSender](tcp/fileSender.go) is a file server. The [fileReciever](tcp/fileReciever.go) requests file from the file server.

## UDP
UDP is a conectionless protocol where no 'session' is established.<br>
The [udp](udp) folder contains all the recepies for working with UDP Sockets using Golang.

**Recipie 1: Simple Server and Client using UDP**

The [simpleServer](udp/simpleServer.go) recieves message from [simpleClient](udp/simpleClient.go).