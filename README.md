# Jack the Rippler: XRPL market optimization software





### What?

Spring_block is a trading algorithm that pulls data from the public Ripple API to create
a real time graph of transactions and uses Bellman-Ford to try to find arbitrage opportunities.
A script then submits a transaction that executes the actual opportunity.



### To run

You can run the program by doing

`./spring_block` on mac os

or executing 

`spring_block.exe` on windows
 



### To build

You will need golang to build from source. Please check how to install it on your system.

`git clone https://github.com/GaspardPeduzzi/spring_block`

To build and execute

`go build && ./spring_block`

