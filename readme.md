This is experimental showing Load Balancer Level 4 & Load Balancer Level 7

# Layer 4 set up
Installation
```
brew install haproxy
```

How does it work
1. start server 1 `cd server1 && go run main.go` and you can check in the browser `localhost:5555`
2. open another terminal and start server 2 `cd server2 && go run main.go` and you can check in the browser `localhost:4444`
3. start load balancer `haproxy -f tcp.cfg`
4. and try to play with this url `localhost:8888`
