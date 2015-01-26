### Lager ###

# Build
```
go build src/lager/lager.go
```

# Use
```
ping 10.0.0.3 | ./lager --prefix=ROUTER_PING
```

# Example
```
[tw-mbp-klaplante lager]$ ping 10.0.0.3 | ./lager --prefix=ROUTER_PING
Jan 25 23:41:16 PST ROUTER_PING: PING 10.0.0.3 (10.0.0.3): 56 data bytes
Jan 25 23:41:16 PST ROUTER_PING: 64 bytes from 10.0.0.3: icmp_seq=0 ttl=64 time=0.043 ms
Jan 25 23:41:17 PST ROUTER_PING: 64 bytes from 10.0.0.3: icmp_seq=1 ttl=64 time=0.133 ms
Jan 25 23:41:18 PST ROUTER_PING: 64 bytes from 10.0.0.3: icmp_seq=2 ttl=64 time=0.137 ms
Jan 25 23:41:19 PST ROUTER_PING: 64 bytes from 10.0.0.3: icmp_seq=3 ttl=64 time=0.105 ms
```

# Help
```
lager -h
```
