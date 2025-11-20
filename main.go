package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/DioneProtocol/odyssey-tooling-sdk-go/keychain"
	"github.com/DioneProtocol/odyssey-tooling-sdk-go/odyssey"
	"github.com/DioneProtocol/odyssey-tooling-sdk-go/subnet"
	"github.com/DioneProtocol/odyssey-tooling-sdk-go/wallet"

	"github.com/DioneProtocol/odysseygo/vms/secp256k1fx"
	"github.com/DioneProtocol/odysseygo/wallet/subnet/primary"
)

var (
	networkStr = flag.String("network", "testnet", "Network: testnet or mainnet")
	keyPath    = flag.String("keypath", "spark-key.pk", "Private key path")
	subnetName = flag.String("subnetname", "Spark Sovereign Infinet ☪️", "Subnet name")
)

func main() {
	flag.Parse()

	var network odyssey.Network
	if *networkStr == "mainnet" {
		network = odyssey.MainnetNetwork()
		fmt.Println("⚠️ MAINNET – Type 'YES' to continue:")
		var confirm string
		fmt.Scanln(&confirm)
		if confirm != "YES" {
			os.Exit(0)
		}
	} else {
		network = odyssey.TestnetNetwork()
	}

	kc, err := keychain.NewKeychain(network, *keyPath, nil)
	if err != nil {
		log.Fatal(err)
	}

	controlKeys := kc.Addresses().List()

	subnetParams := subnet.SubnetParams{
		SubnetEVM: &subnet.SubnetEVMParams{
			ChainID: big.NewInt(54321), // THIS LINE FIXES THE ERROR
		},
		Name: *subnetName,
	}

	newSubnet, err := subnet.New(&subnetParams)
	if err != nil {
		log.Fatal(err)
	}

	newSubnet.SetSubnetControlParams(controlKeys, 1)

	ctx := context.Background()
	wallet, err := wallet.New(ctx, &primary.WalletConfig{
		URI:           network.Endpoint,
		DIONEKeychain: kc.Keychain,
		EthKeychain:   secp256k1fx.NewKeychain(),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deploying Subnet...")
	subnetTx, err := newSubnet.CreateSubnetTx(wallet)
	if err != nil {
		log.Fatal(err)
	}
	subnetID, err := newSubnet.Commit(*subnetTx, wallet, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Subnet deployed! SubnetID: %s\n", subnetID.String())

	time.Sleep(4 * time.Second)

	newSubnet.SetSubnetAuthKeys(controlKeys)

	fmt.Println("Deploying Chain...")
	chainTx, err := newSubnet.CreateBlockchainTx(wallet)
	if err != nil {
		log.Fatal(err)
	}
	blockchainID, err := newSubnet.Commit(*chainTx, wallet, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nSOVEREIGN INFINET LIVE IN <15 SECONDS\n")
	fmt.Printf("   SubnetID:     %s\n", subnetID.String())
	fmt.Printf("   BlockchainID: %s\n", blockchainID.String())
	fmt.Printf("   RPC:          https://api.odyssey.dioneprotocol.com/ext/bc/%s/rpc\n", blockchainID.String())
	fmt.Printf("   Explorer:     https://odysseyscan.dioneprotocol.com/blockchain/%s\n", blockchainID.String())