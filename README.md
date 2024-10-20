# EndServer
#### EndServer handles websocket connections initiated by users. They are responsible for authentication/authorization and persistence of messages. Routing of messages between users connected to different EndServer containers is done using the [DistributionServer](https://github.com/Adarsh-Kmt/DistributionServer). Load balancing of user websocket connections is done by [WebsocketReverseProxy](https://github.com/Adarsh-Kmt/WebsocketReverseProxy).
---
## Features
- Uses gRPC protocol to send messages to the distribution server.
- Implemented mutual TLS between distribution server container and end server container, which enables authentication of both parties, and encryption of messages exchanged.
- Leveraged multistage docker builds to reduce the size of the image by 90.78% (719.26 MB -> 66.32 MB).
- Implemented basic authentication (username/password), and JWT tokens for authorization.
- User info stored in MySQL database.

