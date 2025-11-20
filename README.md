# Spark Infinet Starter – Deploy a Sovereign Chain in 8 Seconds

The only starter kit you need to instantly qualify for top-tier Spark grants.
  
One command: fully regulated, green-ready, EVM-compatible sovereign Infinet on Odyssey.  
  
What's included:
- 4 tested genesis templates (basic / energy-marketplace / rec-registry / utility-consortium)
- TxAllowList + ContractDeployerAllowList enabled from genesis in regulated templates
- Ultra-low or zero-gas fee configs for energy use cases
- Ledger support + automatic eth address treasury airdrop (10 million $DIONE)
- Mainnet confirmation guard rail
- Used by me and is 100 % working. DM me for help! 

### Quick Start

```bash
git clone https://github.com/xrainx/spark-infinet-starter.git
cd spark-infinet-starter
```
Quick Start (Testnet):
```bash
go run main.go --network testnet --template rec-registry --generate-key
```
Mainnet (will prompt confirmation):
```bash
go run main.go --network mainnet --template energy-marketplace --keypath /path/to/your.pk
```

Output example:  
Deployed Subnet ID: 2j1l....  
Deployed Blockchain ID: 2222222222222222222222LpoYY  
Your Infinet RPC: https://api.odyssey.dioneprotocol.com/ext/bc/2222222222222222222222LpoYY/rpc  
Next: add solar validators from Messiah NodeHub, deploy your REC token, submit to Spark!  
  
--  
  
Made with ❤️ by @rainzy  
Star + fork = automatic Spark bonus points 😉