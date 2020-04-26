#!/bin/sh

#
#  Copyright 2019 Pnoker. All Rights Reserved.
#

set -e

mongod --smallfiles --bind_ip_all &

while true; do
  mongo /pnoker/mongo/config/mongo-init.js && break
  sleep 5
done

wait