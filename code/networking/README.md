# Networking Practice

Roadmap: `roadmap/networking.md` · Notes: `notes/networking/` · Dictionary: `reference/networking.md`

## What practice looks like
- Small Go programs: TCP/UDP echo servers, HTTP clients/servers with timeouts, WebSocket/SSE, mTLS.
- Packet captures (`.pcap`) with `tcpdump`, analyzed in Wireshark; command output from `dig`, `curl -v`, `ss`, `openssl s_client`.
- Docker setups: nginx reverse proxy in front of Go replicas, network namespaces, `tc netem` fault injection, kind for Kubernetes networking.

## Environment setup
```bash
# Tools (Debian/Ubuntu)
sudo apt install -y tcpdump wireshark tshark dnsutils iproute2 iputils-ping \
  traceroute mtr-tiny netcat-openbsd iperf3 openssl curl nftables conntrack

# Capture loopback traffic for a local Go server
sudo tcpdump -i lo -n port 9000 -w capture.pcap

# Decrypt TLS in Wireshark: run the client with
export SSLKEYLOGFILE=$PWD/keys.log   # then set it in Wireshark: TLS → (Pre)-Master-Secret log

# Debug container with all network tools
docker run --rm -it --network host nicolaka/netshoot

# nginx reverse proxy
docker run -d --name nginx -p 8080:80 -v $PWD/nginx.conf:/etc/nginx/nginx.conf:ro nginx
```

## Folder convention
```
code/networking/
├── README.md
└── <NN>-<topic-slug>/        # e.g. 05-tcp, 13-tls-pki
    ├── README.md             # exercise instructions (from the 🛠 Practice line of NET <NN>)
    └── ...                   # main.go, nginx.conf, compose.yaml, *.pcap, findings
```
- `<NN>` = topic number in the roadmap (NET 5 → `05-tcp`). Create the folder when the topic starts.
- Keep pcaps small; write what you saw (and the filter used) in the exercise README.
