#!/bin/bash

/usr/sbin/tcpdump -i eth0 -s 0 -W 50 -C 50 -w /tmp/tcpdump.pcap
