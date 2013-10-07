cs425-mp2 Distributed gossiping protocol for failure detection
==============================================================

Authors
-------

Karan Shah (kshah3)
Martin Wozniewicz (wozniew1)


Compiling and Running
---------------------
You'll need [Go](http://golang.org) to build this project.

Call `make` in root directory. 
There should be a `dgroup` executable. Call 

    ./dgroup -h

to see explanations of the parameters. The first node of the group 
can just be started with `./dgroup`. This will use the default parameters.

After that, each new node needs to be given the address of a node already
in the group. You can also specify the port to listen on, useful if you 
are running several instances on the same machine.


Algorithm
---------

Nodes pick a random recipient from all known machines at random each 
"heartbeat". They then send information about all the other machines to
that node, which compares its own records with the ones being sent. 

Each node keeps track of the number of heartbeats since another node
was last heard from. Once these heartbeats cross a constant threshold, 
the machine is presumed to have failed.

At each heartbeat, each node only broadcasts to a single other node on
the network, saving bandwidth and spreading out the cost of keeping 
everyone updated. 

The packets are made up of a subject ID (IP,timestamp of creation) and the
heartbeat value for that machine.


Benchmarks
----------

At a heartbeat interval of 50ms, with four machines, average network usage
is approximately 5KB/s. Each "gossip" is between 30B and 80B, depending
on the exact contents of the network packet. With four machines, each 
heartbeat means that a node will send approximately 250B of network traffic.

This is a very conservative estimate, since the machines detect failure
within 2 seconds, well within the requirement. A longer heartbeat interval
could be found to decrease network usage, and lengthen detection time.


1. When a node joins, the bandwidth is the size of it's address. So it's a 
20B packet, sent to a single machine.

2. When a node dies, bandwidth actually decreases, since it is no longer 
broadcasting. Other nodes still broadcast its information, so the total decrease
is only about `80(N-1) + 60 bytes`.

3. When a node quits, the bandwidth used is the same as joining (it only has to 
send this to a single other node, that will handle propagating the information)

The experiment to come up with these values involved printing the size of 
each packet as it comes in, as well as thinking about how our algorithm works
at a high level.


Integration with MP1
--------------------

MP1 wasn't very useful in debugging this MP, since we needed to be able to see
log messages in real-time (to see the order in which packets were sent/arrived
at other machines). Instead of using mp1, we used the built in logging 
functions of Go, and set up parameters for testing locally instead of in a 
truly distributed environment. This allowed us to view logs for all "nodes" on a
single screen in real-time. 
