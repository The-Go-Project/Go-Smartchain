# <img src="https://github.com/The-Go-Project/Go-Smartchain/blob/master/go-logo.png" width="200" height="35"/>


![Build Check](https://github.com/jpmorganchase/quorum/workflows/Build%20Check/badge.svg?branch=master)
[![Docker Pulls](https://img.shields.io/docker/pulls/quorumengineering/quorum)](https://hub.docker.com/r/quorumengineering/quorum)
[![Discord](https://img.shields.io/discord/697535391594446898)](https://discord.com/channels/697535391594446898/747810572937986240)


Go-Smartchain is an Ethereum-based distributed ledger protocol with transaction/contract privacy and new consensus mechanisms.

Go-Smartchain is a fork of [go-ethereum](https://github.com/ethereum/go-ethereum) and is updated in line with go-ethereum releases.

Key enhancements over go-ethereum:

* [__Privacy__](https://consensys.net/docs/goquorum//en/latest/concepts/privacy/) - Go-Smartchain supports private transactions and private contracts through public/private state separation, and utilises peer-to-peer encrypted message exchanges (see [Tessera](https://github.com/consensys/tessera)) for directed transfer of private data to network participants
* [__Alternative Consensus Mechanisms__](https://consensys.net/docs/goquorum//en/latest/concepts/consensus/overview/) - with no need for POW/POS in a permissioned network, Go-Smartchain instead offers multiple consensus mechanisms that are more appropriate for consortium chains:
    * [__QBFT__](https://consensys.net/docs/goquorum/en/latest/configure-and-manage/configure/consensus-protocols/qbft/) - Improved version of IBFT that is interoperable with Hyperledger Besu
    * [__Istanbul BFT__](https://consensys.net/docs/goquorum/en/latest/configure-and-manage/configure/consensus-protocols/ibft/) - a PBFT-inspired consensus algorithm with transaction finality, by AMIS.
    * [__Clique POA Consensus__](https://github.com/ethereum/EIPs/issues/225) - a default POA consensus algorithm bundled with Go Ethereum.
    * [__Raft-based Consensus__](https://consensys.net/docs/goquorum/en/latest/configure-and-manage/configure/consensus-protocols/raft/) - a consensus model for faster blocktimes, transaction finality, and on-demand block creation
* [__Peer Permissioning__](https://consensys.net/docs/goquorum/en/latest/concepts/permissions-overview/) - node/peer permissioning, ensuring only known parties can join the network
* [__Account Management__](https://consensys.net/docs/goquorum/en/latest/concepts/account-management/) - Go-Smartchain introduced account plugins, which allows GoQuorum or clef to be extended with alternative methods of managing accounts including external vaults.
* [__Pluggable Architecture__](https://consensys.net/docs/goquorum/en/latest/concepts/plugins/) -  allows adding additional features as plugins to the core `geth`, providing extensibility, flexibility, and distinct isolation of GoQuorum features.
* __Higher Performance__ - Go-Smartchain offers significantly higher performance throughput than public geth


## Third Party Tools/Libraries

The following Go-Smartchain related libraries/applications have been created by Third Parties and as such are not specifically endorsed by J.P. Morgan.  A big thanks to the developers for improving the tooling around Go-Smartchain!

* [Go-Explorer Blockchain Explorer](https://github.com/The-Go-Project/Go-Explorer) - a Blockchain Explorer for Go-Smartchain which supports viewing private transactions
* [ERC20 REST service](https://github.com/web3labs/erc20-rest-service) - a Go-Smartchain supported RESTful service for creating and managing ERC-20 tokens



## Reporting Security Bugs
Security is part of our commitment to our users. At Go-Smartchain we have a close relationship with the security community, we understand the realm, and encourage security researchers to become part of our mission of building secure reliable software. This section explains how to submit security bugs, and what to expect in return.

## License

The go-ethereum library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html), also
included in our repository in the `COPYING.LESSER` file.

The go-ethereum binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also included
in our repository in the `COPYING` file.

Any project planning to use the `crypto/secp256k1` sub-module must use the specific [secp256k1 standalone library](https://github.com/ConsenSys/goquorum-crypto-secp256k1) licensed under 3-clause BSD.
