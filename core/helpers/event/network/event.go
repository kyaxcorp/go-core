package network

// Events between processes through network
/*
There be a way of sending the data to:
- an EndPoint through HTTP
	- This will require the EndPoint Address
	- Credentials to authenticate to the Listener?!
- through a websocket connection
	- The server becomes the dispatcher itself...
	- the server is a 1 service which is running in background automatically, and different functions
	can send dispatches through this listener!
	- the listener should connect to the server and authenticate... authentication is through REQUEST GET PARAM or
	Header
	-

Connections should be encrypted with Certificates!



It would be much much easier if the programmer should not write where and how etc etc...
It will just define: cluster.user.created and the system auto connects, detects etc, the client will receive a JSON?!
or even a decoded interface{} with data...

Events are like pipes, we defined different pipes names and we communicate through them!
It's easier to work with pipes knowing only the names! but there are sometimes and drawbacks using only names...
We should see how we handle that also...


To build such a network of microservices talking to each other, we should create an algo how they will find each other
in some environment, it's always simpler if there is a broker between services....
So, we can create a script and call it broker or somehow like that, and add it to command line...
It will handle automatically all requests and it will also forward them!
In the configuration file, we should also add the settings to connect to broker!

There will be a special broker handling different things on different routes!
*/

/*
18.07.2021
Any client should always try to connect to the broker..
This setting can be turned off... but by default in the setting this option will be on!
It will automatically connect through the broker client

WE should see how the dispatchers and listeners will work...
we want that:
1. the delivering and receiving the message be very fast
2. do not consume a lot of traffic by doing broadcasting to multiple clients
3. maybe each event can be a separate connection to the broker
4. if the programmer calls the Listen function it automatically creates the connection in a separate goroutine
and saving the event in the stack somewhere...because later we can remove it!
5. How we will know that someone has received a message?!
The principle is that no one knows and doesn't guarantee receivable of a message
6. The dispatcher also if calls the dispatch command, it will create a new connection and it will be saved in the stack
The connection will be maintained until the program shuts off or we can set some kind of a timeout for a connection to live...
7. Events will be identified by names!
8.

Best way is to:
Have 1 to 1 Connection even if the program has multiple listeners or dispatchers
On the broker side we should create some kind of filter or splitter.
It should identify the dispatchers and the subscribers, and in this case, it will know to whom it should send the message
But there are also drawbacks...
We should write different events of subscription...
If a program terminates, the broker will know that has terminated and will kill the stack with subscribers dispatchers from there
If the broker dies, the client should take care to re-subscribe to the broker...

And if the internet is lost on the client side... there will be LOTS of reconnections going on... that's insane
and not very practical... the socket stack can go insane....

So:
- 1 to 1 connection
- On the client side there will be go routine which will take care of the connection
	- it will also have a writer or simple a function that sends the messages to the server
	- maybe we should create a buffer?!... but that depends...
	- it will auto reconnect
	- it will have a stack will all known Subscribers and Dispatchers?
	- Dispatchers are not obligatory to memorize
	- We should save the listeners/subscribers and keep them until the termination of the program!
	- If the app reconnects, it should send again all subscribers to the broker containing:
		- event name
- On the server side there will also be a handler... which will:
	- memorize all subscribers having:
		- event name:
		- connection id
	- when receiving a message it should create instantly a goroutine which will perform the processing and forwarding the
	messages
	-
-------------------------------------
TODO: MAYBE WE SHOULD USE THE SAME OBJECT event that is in 1 program... but overlaying with simple functions and ws things

in this case the events can work much slower than usual... and everything will be chained with payload id or event id
a priority will be formed... and the broker will send the data to the first listener.. and if it's async, it will not
wait for a response... but if it's sync it will wait for one, and then send the message to the next one in the priority...
The response that has being received from the client should be sent forward to next one...
THe idea is that the data it's being processed and passed along by the priority order
There is a problem with connection... if a client disconnects in the moment of processing we should decide what to do:
	- an idea is to wait for him to reconnect (how we will identify and know that is him that was missing?!)
     we need to generate some kind of unique id on the client side which will identify on reconnect!
	- don't wait and go further... if it's an error of sending, also go further...
For the server we can use the broker... but in the hub/pipe we can create a special handler...which will transmit
the events where's necessary

*/

/*
if using pipes, they are just transmitters of data... communication between microservices, neurons
none of them are storing anything... so, to have a centralized event manager, we need to create one...it means a
manager which will be connected to the same piping system
*/

func Listen(eventName string, listenOptions *ListenOptions) {
	go _listen(eventName, listenOptions)
}
