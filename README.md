# BFTWithoutSignatures_Client

### Contents
- [About](#about)
- [Installation](#installation)
- [Execution](#execution)
- [References](#references)


## About
A Golang with ZeroMQ implementation of the client for the algorithm:
<div style="font-size: 15px">
From Consensus to Atomic Broadcast: Time-Free Byzantine-Resistant Protocols without Signatures
</div>
<div style="font-size: 13px">
    By Miguel Correia, Nuno Ferreira Neves, Paulo Verissimo
</div>


## Installation
### Golang
If you have not already installed **Golang** follow the instructions [here](https://golang.org/doc/install).
### Clone Repository
```bash
cd ~/go/src/
git clone https://github.com/v-petrou/BFTWithoutSignatures_Client.git
```


## Execution
### Manually
Open a different terminal for each client. In each terminal run:
```bash
go install BFTWithoutSignatures_Client
BFTWithoutSignatures_Client <ID> <N> <Remote>
```
### Script
Adjust the script (BFTWithoutSignatures_Client/scripts/run.sh) and run:
```bash
bash ~/go/src/BFTWithoutSignatures_Client/scripts/run.sh
```
When you are done and want to kill the processes run:
```bash
bash ~/go/src/BFTWithoutSignatures_Client/scripts/kill.sh
```


## References
- [From Consensus to Atomic Broadcast: Time-Free Byzantine-Resistant Protocols without Signatures](https://www.researchgate.net/publication/220459271_From_Consensus_to_Atomic_Broadcast_Time-Free_Byzantine-Resistant_Protocols_without_Signatures)