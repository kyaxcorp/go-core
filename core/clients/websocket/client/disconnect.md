There are different types of disconnections....

1. When the server closes gracefully the connection by sending the close packet
2. When server dies, and doesn't send anything... the connection may still be alive, but ping/pong mechanism will later handle that!
3. When internet is bad, and packets are not received well or in time
4. 