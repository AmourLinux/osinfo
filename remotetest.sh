#!/bin/bash
set -ex

RemoteAddr="192.168.80.77"

#ssh -i ~/Desktop/ssh/cx-k8s.pem root@${RemoteAddr} "mkdir /root/osinfo"

scp -r -i ~/Desktop/ssh/cx-k8s.pem * root@${RemoteAddr}:/root/osinfo

ssh -i ~/Desktop/ssh/cx-k8s.pem root@${RemoteAddr} "cd /root/osinfo && go test -v ."
