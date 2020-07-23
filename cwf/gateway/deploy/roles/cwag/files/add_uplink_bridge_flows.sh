#!/usr/bin/env bash
#
# Copyright (c) 2018-present, Facebook, Inc.
# All rights reserved.
#
# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree. An additional grant
# of patent rights can be found in the PATENTS file in the same directory.

ovs-ofctl del-flows uplink_br0

# gw0 configuration, TODO might want to avoid NORMAL for perf reasons
ip_addr=$(/sbin/ifconfig gw0 | grep -m 1 inet | awk '{ print $2}')
ovs-ofctl add-flow uplink_br0 "table=0, priority=999, arp, nw_src=$ip_addr, actions=NORMAL"
ovs-ofctl add-flow uplink_br0 "table=0, priority=999, arp, nw_dst=$ip_addr, actions=NORMAL"
ovs-ofctl add-flow uplink_br0 "table=0, priority=999, ip, nw_src=$ip_addr, actions=NORMAL"
ovs-ofctl add-flow uplink_br0 "table=0, priority=999, ip, nw_dst=$ip_addr, actions=NORMAL"

# Some setups might not have 2 nics. In that case just use $1
if [ -n "$1" ] && [ -n "$2" ]
then
  ovs-vsctl --may-exist add-port uplink_br0 "$1"
  ovs-vsctl --may-exist add-port uplink_br0 "$2"
  ovs-ofctl -O OpenFlow14 add-group uplink_br0 "group_id=42, type=select, selection_method=dp_hash, bucket=actions=output:$1, bucket=actions=output:$2"
  ovs-ofctl add-flow uplink_br0 "table=0, priority=10, in_port=$1, actions=output:uplink_patch"
  ovs-ofctl add-flow uplink_br0 "table=0, priority=10, in_port=$2, actions=output:uplink_patch"
  ovs-ofctl add-flow uplink_br0 "table=0, priority=10, in_port=uplink_patch, actions=group:42"
elif [ -n "$1" ]
then
  ovs-vsctl --may-exist add-port uplink_br0 "$1"
  ovs-ofctl add-flow uplink_br0 "table=0, priority=10, in_port=$1, actions=output:uplink_patch"
  ovs-ofctl add-flow uplink_br0 "table=0, priority=10, in_port=uplink_patch, actions=output:$1"
else
  ovs-ofctl add-flow uplink_br0 "table=0, priority=10, actions=LOCAL"
fi