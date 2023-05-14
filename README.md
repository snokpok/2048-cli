# 2048-cli

I was flying back home when I saw someone playing 2048 on the inflight's infotainment screen.

![A woman that "might" be playing 2048 on the airplane's infotainment system](./demonstration.jpg)

I was bored and was without Wifi so I built 2048 but you can play it on the CLI!

![The real 2048](./2048-game.jpeg)


### Installation

Prerequisites:
- Go ^1.18

1. `go build .`
2. `./2048-cli`


### An amazing tactical gameplay:
        
```
$ ./2048-cli 
5 _ _ _ 
_ _ _ _ 
_ _ _ _ 
_ _ _ _ 

> right
5 _ _ 5 
_ _ _ _ 
_ _ _ _ 
_ _ _ _ 

> down
5 _ _ _ 
_ _ _ _ 
_ _ _ _ 
5 _ _ 5 

> up
10 5 _ 5 
_ _ _ _ 
_ _ _ _ 
_ _ _ _ 

> up
10 5 5 5 
_ _ _ _ 
_ _ _ _ 
_ _ _ _ 

> right
5 10 5 10 
_ _ _ _ 
_ _ _ _ 
_ _ _ _ 

> ^C
```

### TODO
- [ ] Implement some more specific cell merging rules of the official 2048 game
- [ ] Option to expand board size (e.g. 5x5, 6x6, etc...)