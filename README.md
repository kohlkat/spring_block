# Jack the Rippler: XRPL market optimization software

```
       __           __      __  __                 _             __         
      / /___ ______/ /__   / /_/ /_  ___     _____(_)___  ____  / /__  _____
 __  / / __ `/ ___/ //_/  / __/ __ \/ _ \   / ___/ / __ \/ __ \/ / _ \/ ___/
/ /_/ / /_/ / /__/ ,<    / /_/ / / /  __/  / /  / / /_/ / /_/ / /  __/ /    
\____/\__,_/\___/_/|_|   \__/_/ /_/\___/  /_/  /_/ .___/ .___/_/\___/_/     
                                                /_/   /_/                   

```


### What?

Spring_block is a trading algorithm that pulls data from the public Ripple API to create
a real time graph of transactions and uses Bellman-Ford to try to find arbitrage opportunities.
A script then submits a transaction that executes the opportunity.

### Why?

This tool can be used for multiple purposes :
 
    - helping create more liquidity on the market
    
    - analyze behaviours of actors on the network


 



### To run

You can run the program by doing

`./deploy.sh` on mac os
 
then go to

`localhost:8080`
 



### To build

You will need golang to build from source. Please check how to install it on your system.

`git clone https://github.com/GaspardPeduzzi/spring_block`

To build and execute

`go build && ./spring_block`

