#!/bin/bash
aws_profile=s3-pcap
backup_path=/var/log/suricata
bucket_name=my-pcap-bucket

aws s3 sync $backup_path s3://$bucket_name --exclude *.log --exclude *.json --include log.pcap.* --profile=$aws_profile