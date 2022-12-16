# ScanMC
Golang learning Project Portscanner to find Liveoverflows minecraft server.

Yout give it a File which looks like:

127.0.0.1
127.0.0.1
127.0.0.1
127.0.0.1
127.0.0.1
...

And a Port, and it will scan that Port for a Minecraft Server on every address and safe the results in a mongodb.
I recommand using something like masscan to gater the hosts in the fisrt place.
Then reorder the output from something like this to this:

"open tcp 25565 127.0.0.1 1670874259" -> "127.0.0.1" 

per line.

