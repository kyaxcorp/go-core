package config

import loggerConfig "github.com/kyaxcorp/go-core/core/logger/config"

/*
We need to have a unique name of the instance
Based on the name, others can make requests to it through broker...
It will send push notifications through FCM
If we want to have multiple connections with FCM, we just add in configuration another instance
of it, and launch it!

*/

/*
So we will have listeners like:
- HTTP (1)
- BROKER CHANNEL (multiple ones)

Each of them can be enabled/disabled

So that's where we receive the full structure of the message and to who to send!
----------------------------------
Second is that, these listeners should write to the DB!
If they fail, they should write into a tmp local db like sqlite and then retry
The retry should be made only for a specific period of time... because if the database went offline for 2 hours
it's meaningless to send the push notifications!
-----------------------------------
Third -> there is the 2-nd thread which sends the notifications! It will always read from the DB
It will receive triggers with specific ID's from the one that made the writes to the DB! so that it will read from DB
That maybe will improve read performance!
------------------------------------
We can add here different options like:
- delete the row from db after success sending
-
------------------------------------
files can also be sent, they are simply converted to bytes and added to a specific table with files?!?!


*/

type Config struct {
	// API Key -> it's FCM authentication
	ApiKey string
	// We should have here broker instances to which we will be listening/subscribing
	ListenOnBrokerInstances []string
	ListenHttpsAddress      string
	// Auto generation of certificate?! or add standard configuration!?

	// TODO: listen on a port?! to receive through HTTP API?!
	// TODO: it should be HTTPS by default!!!!
	// TODO: receive triggers through Broker! -> we should subscribe to a channel/pipe/hub

	// TODO: we should have 2 threads
	// TODO: 1 will send messages
	// TODO: 2 will receive and write messages to the DB to add to the queue and after that will just notify the 1

	// It's better that if someone wants to send a notification, it should send through the Broker or the HTTP Server
	// If the connection is bad or something like that, and the request from the client hasn't being received or lost
	// We should create a persistent local storage for things like that... where data will be stored in JSON format
	// maybe in an encrypted version, and afterwards will send!
	// It can be even a DB File or something like that...

	// TODO: we should read the queue from the DB

	Logger loggerConfig.Config
}
