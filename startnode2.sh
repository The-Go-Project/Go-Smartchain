#!/bin/bash
PRIVATE_CONFIG=ignore nohup ./build/bin/geth \
--datadir node2 \
--nodiscover \
--verbosity 5 \
--networkid 33333 \
--raft \
--raftport 50408 \
--rpc \
--rpcaddr 0.0.0.0 \
--rpcport 22008 \
--rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,raft \
--emitcheckpoints \
--port 21008 \
>> node2.log 2>&1 &
