1. Create cluster nodes
2. Create connections as a proxy to each cluster node
3. forward/receive messages between nodes
4. each message sent through him should have a payload id which identifies the message
5. the payload id should be generated automatically by the client
6. Add WithCancel for Start/Stop