package main

import (
	"encoding/json"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/jtremback/crypto-conditions"
	"golang.org/x/crypto/ed25519"
)

const (
	bdbServer = "localhost:59984"
	VERSION   = "0.9"
)

// Returns a pair of (private_key, public_key) encoded in base58.
func generateKeypair() *Keypair {
	//	DEPRECATED: crypto lib uses /dev/random or /dev/urandom internally
	//	fh, err := os.Open("/dev/urandom")
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer fh.Close()

	// Create random public and private keys
	publicKeyBytes, privateKeyBytes, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}

	publicKeyBase58 := base58.Encode(publicKeyBytes)
	privateKeyBase58 := base58.Encode(privateKeyBytes)

	return &Keypair{
		PrivateKey: privateKeyBase58,
		PublicKey:  publicKeyBase58,
	}
}

func generateKeypairForAlice() *Keypair {
	return &Keypair{
		PrivateKey: "HDN1ajL1k1MyNvF7ETwT7K1Jdqw2yNAx4eN3TmDXtW7J",
		PublicKey:  "AZhTRnY2QhX7rNqjPsfKDK4UJQhk9Hm6yecmpzQwRsrM",
	}
}

func generateKeypairForBob() *Keypair {
	return &Keypair{
		PrivateKey: "6aoSvdZePgsCedHmFaidWXnGVpBgyrCBqb8xobsiF67",
		PublicKey:  "Cs4byJu7ZqzTFfCnTwiQbCqDxoBgpeW3joSZ8MGQCmWg",
	}
}

func NewAsset(data string) *Asset {
	return &Asset{
		Data: data}
}

func NewInput(
	ownersBefore []string,
	fulfillment *OutputCondition,
	fulfills *Output) *Input {
	return &Input{
		OwnersBefore: ownersBefore,
		Fulfillment:  fulfillment,
		Fulfills:     fulfills,
	}
}

func NewOutput(
	amount int32,
	publicKeys []string,
	outputCondition *OutputCondition) *Output {
	return &Output{
		Amount:     amount,
		PublicKeys: publicKeys,
		Condition:  outputCondition,
	}
}

func NewOutputCondition(
	details string,
	uri string) *OutputCondition {
	return &OutputCondition{
		Details: details,
		Uri:     uri,
	}
}

func main() {
	// Step 1. Create a keypair for creator
	creator := generateKeypair()
	// Step 2. Create a keypair for receiver
	receiver := generateKeypair()

	fmt.Printf("DEBUG: Creator: %s\n", creator)
	fmt.Printf("DEBUG: Receiver: %s\n", receiver)

	// Step 3. Create Operation
	// Step 3.1 Prepare tx for Create Op
	// Step 3.1.1 Create the asset

	assetData := map[string]interface{}{
		"bicycle": map[string]interface{}{
			"serial_number": "abcd1234",
			"manufacturer":  "bkfab",
		},
	}
	assetBytes, err := json.Marshal(assetData)
	if err != nil {
		panic(err)
	}
	asset := NewAsset(string(assetBytes))
	assetBytes, err = json.Marshal(asset)
	if err != nil {
		panic(err)
	}
	fmt.Printf("DEBUG: Asset: %s\n", assetBytes)

	// Step 3.1.2 Create the inputs
	prevOwners := []string{creator.PublicKey}
	input := NewInput(prevOwners, nil, nil)
	fmt.Printf("DEBUG: Input: %s\n", input)

	// Step 3.1.3 Create the outputs
	ed25519, err := CryptoConditions.ParseEd25519Fulfillment(
		[]byte(creator.GetPublicKey()))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Printf("DEBUG: ed25519: %s\n", ed25519)

	//details:=
	//uri:=
	//outputCondition := NewOutputCondition()

	//amount := 1
	//publicKeys := []string{creator.PublicKey}
	//output := NewOutput(amount, publicKeys, out?)

	// Step 3.1.4 Create the metadata
	// Step 3.1.5 Create the full tx body

	// Step 3.2. Fulfill tx for Create Op
	// Sent 3.3. Sent Tx for Create Op
	// Step 4. Transfer Operation
	// Step 4.1 Prepare tx for Transfer Op
	// Step 4.2. Fulfill tx for Transfer Op
	// Sent 4.3. Sent Tx for TransferOp

}

/*

func NewTxPayloadVerZeroNine(
	id string,
	op OpType,
	asset *Asset,
	inputs []Input,
	outputs []Output,
	metadata map[string]string) *TxPayloadVerZeroNine {
	return &TxPayloadVerZeroNine{
		Id:          id,
		Operation:   op,
		AssetObject: asset,
		Inputs:      inputs,
		Outputs:     outputs,
		Metadata:    metadata,
	}
}

func NewTransaction(
	version string,
	tx *TxPayloadVerZeroNine) *Transaction {
	return &Transaction{
		Version:              version,
		TxPayloadVerZeroNine: tx,
	}
}


func createTransaction(
	signers []ed25519.PublicKey,
	recipients []ed25519.PublicKey,
	asset *Asset,
	metadata map[string]string,
	inputs []Input,
	outputs []Output) *Transaction {

	// TODO sanity

	// As per docs in py driver!
	if recipients == nil {
		recipients = signers
	}
	id := "id" //TODO
	txZeroNine := NewTxPayloadVerZeroNine(
		id,
		OpCreate,
		asset,
		inputs,
		outputs,
		metadata)

	tx := NewTransaction(VERSION, txZeroNine)
	return tx
}
*/
