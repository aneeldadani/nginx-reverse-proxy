#!/bin/bash
aws_profile=s3-pcap
backup_path=/tmp
bucket_name=my-pcap-bucket

aws s3 sync $backup_path s3://$bucket_name --exclude *.log --include *.pcap* --profile=$aws_profile
