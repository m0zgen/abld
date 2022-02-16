#!/bin/bash
# Sys env / paths / etc
# -------------------------------------------------------------------------------------------\
PATH=$PATH:/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin
SCRIPT_PATH=$(cd `dirname "${BASH_SOURCE[0]}"` && pwd)

#

chown -R blockyusr:blockyusr /opt/blocky/

systemctl stop blocky

sleep 3

setcap cap_net_bind_service=ep /opt/blocky/blocky

systemctl start blocky