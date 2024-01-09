# Blocky - Backend Listener Daemon (BLD)

* [OpenBLD.net](https://openbld.net) DNS project uses BLD as backend DNS server.
* BLD based on [Blocky](https://github.com/0xERR0R/blocky).

## How OpenBLD.net works

- [How it works](https://openbld.net/docs/overwiew/how-it-works/)
- [Get Started](https://openbld.net/docs/category/get-started/)
- [Main Goals](https://openbld.net/docs/overwiew/main-goals/)
- [Philosophy](https://openbld.net/docs/intro/#philosophy)

## Features

- **Performance** - Improves speed and performance in your network

    * Customizable caching of DNS answers for queries -> improves DNS resolution speed and reduces amount of external DNS
    queries
    * Prefetching and caching of often used queries
    * Using multiple external resolver simultaneously
    * Low memory footprint

- **Various Protocols** - Supports modern DNS protocols

    * DNS over UDP and TCP
    * DNS over HTTPS (aka DoH)
    * DNS over TLS (aka DoT)
    * Supports modern DNS extensions: DNSSEC, eDNS, ...

- **Integration** - various integration

  * [Prometheus](https://prometheus.io/) metrics
  * Prepared [Grafana](https://grafana.com/) dashboards (Prometheus and database)
  * Logging of DNS queries per day / per client in CSV format or MySQL/MariaDB/PostgreSQL database - easy to analyze
  * CLI tool

- **Simple installation/configuration** - blocky was designed for simple installation

    * Stateless (no database, no temporary files)
    * Docker image with Multi-arch support
    * Single binary
    * Supports x86-64 and ARM architectures -> runs fine on Raspberry PI

## Quick start

- You can jump to [Installation](https://0xerr0r.github.io/blocky/latest/installation/) chapter in the documentation.
- Full documentation and configuration examples on [https://0xERR0R.github.io/blocky/](https://0xERR0R.github.io/blocky/)

## Blocking domain / web-sources issues

You can inform me about of accidents blocked / unblocked domains through [DNS-HOLE](https://github.com/m0zgen/dns-hole.git) repo.

## Links

* https://openbld.net/
* https://github.com/m0zgen/blocky-installer.git
* https://github.com/m0zgen/prometheus-stack-installer.git
* https://github.com/m0zgen/dns-hole.git
* https://github.com/m0zgen/dns-tester.git
* https://github.com/m0zgen/bench-dns.git
* https://github.com/m0zgen/bld-agregator
* https://github.com/m0zgen/ip2drop
* https://github.com/m0zgen/cactusd
* https://github.com/m0zgen/zdns
* https://lab.sys-adm.in
