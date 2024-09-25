#!/bin/bash
PRIVATE_CONFIG=ignore nohup ./build/bin/geth \
--datadir node1 \
--nodiscover \
--verbosity 5 \
--networkid 33333 \
--raft \
--raftport 50407 \
--rpc \
--rpcaddr 0.0.0.0 \
--rpcport 22007 \
--rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,raft \
--emitcheckpoints \
--port 21007 \
>> node1.log 2>&1 &
