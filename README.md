# **GO + GRPC**

From Youtube:

https://www.youtube.com/watch?v=a6G5-LUlFO4

Could be used a a build block/ template for more complex projects.


Project uses gRPC in 4 different modes:

* Unary
* Server streaming . Client request
* Client Streaming . Server request
* Bidirectional streaming : with GO routine and channel.

Client & Server have their own directory.

Uses protocol buffers a.k.a protobuf

At the time of working on this project, the version of protoc for Windows was broken. 
I had to run protoc to generate the GO code on Linux (Ubuntu for WSL)

