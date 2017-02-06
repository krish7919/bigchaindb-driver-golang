package main

import (
	"fmt"
	"math/big"

	"github.com/tv42/base58"
	"golang.org/x/crypto/ed25519"
)

type OpType string

const (
	version           = "0.9"
	bdbServer         = "localhost:59984"
	OpCreate   OpType = "CREATE"
	OpTransfer OpType = "TRANSFER"
	OpGenesis  OpType = "GENESIS"
)

type Transaction struct {
	//version number of the transaction model
	Version string `json:"version"`

	//tx corresponding to the version
	// TODO Have a union-like struct here?
	*TxPayloadVerZeroNine
}

type TxPayloadVerZeroNine struct {
	//The id of the transaction, and also the database primary key.
	Id string `json:"id"`

	//Type of the transaction
	Operation OpType `json:"operation"`

	//Description of the asset being transacted.
	AssetObject *Asset `json:"asset"`

	Inputs []Input `json:"inputs"`

	Outputs []Output `json:"output"`

	// TODO support for metadata
	Metadata map[string]string `json:"metadata"`
}

//Input contains a pointer to an unspent output and a crypto fulfillment that
//satisfies the conditions of that output. A fulfillment is usually a signature
//proving the ownership of the asset.
//An input spends a previous output, by providing one or more fulfillments that
//fulfill the conditions of the previous output.
type Input struct {
	//List of public keys of the previous owners of the asset.
	OwnersBefore []string `json:"owners_before"`

	//A payload that satisfies the condition of a previous output to prove that
	//the creator(s) of this transaction have control over the listed asset.
	Fulfillment *OutputCondition `json:"fulfillment"`

	//Reference to the output that is being spent.
	Fulfills *Output `json:"fulfills"`
}

//Output contains crypto-conditions that need to be fulfilled by a transfer
//transaction in order to transfer ownership to new owners.
type Output struct {
	//Integral amount of the asset represented by this condition.
	Amount int `json:"amount"`

	//List of public keys associated with the conditions on an output.
	PublicKeys []string `json:"public_keys"`

	Condition *OutputCondition `json:"condition"` // TODO ?
}

// Describes the condition that needs to be met to spend the output.
// TODO
type OutputCondition struct {
	Details string `json:""` //TODO

	Uri string `json:""`
}

type Asset struct {
	//ID of the transaction that created the asset. Is this uuid?
	Id string `json:"id"`
	//User provided metadata associated with the asset. May also be null.
	//TODO
	// For the time being, we will assume it to be json string.
	//(As in the bicycle example)
	Data map[string]interface{} `json:"data"`
	// TODO Divisible  bool                   `json:"divisible"`
	// TODO Refillable bool                   `json:"refillable"`
	// TODO Updatable  bool                   `json:"updatable"`
}

/*func NewTransaction(
	version string,
	tx *TxPayloadVerZeroNine) *Transaction {
	return &Transaction{
		Version: version,
		tx,
	}
}*/

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

func NewInput(
	prevOwners []string,
	outputCondition *OutputCondition,
	output *Output) *Input {
	return &Input{
		OwnersBefore: prevOwners,
		Fulfillment:  outputCondition,
		Fulfills:     output,
	}
}

func NewOutput(
	amount int,
	publicKeys []string,
	outputCondition *OutputCondition) *Output {
	return &Output{
		Amount:     amount,
		PublicKeys: publicKeys,
		Condition:  outputCondition,
	}
}

/* TODO
func NewOutputCondition(
    details ??,
    uri string) *OutputCondition {
    return &OutputCondition{
        Details:
        Uri: uri,
    }
}
*/

func NewAsset(
	assetId string,
	data map[string]interface{}) *Asset {
	return &Asset{
		Id:   assetId,
		Data: data}
}

type Keypair struct {
	PublicKey  []byte
	PrivateKey []byte
}

// Returns a pair of (private_key, public_key) encoded in base58.
func generateKeypair() *Keypair {
	/* DEPRECATED: crypto lib uses /dev/random or /dev/urandom internally
	fh, err := os.Open("/dev/urandom")
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	*/
	// Create random public and private keys
	publicKeyBytes, privateKeyBytes, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}

	// Convert random bytes to BigInt so that we can easily convert it to base58
	i := new(big.Int)
	var publicKeyBase58, privateKeyBase58 []byte
	publicKeyBase58 = base58.EncodeBig(publicKeyBase58, i.SetBytes(publicKeyBytes))
	privateKeyBase58 = base58.EncodeBig(privateKeyBase58, i.SetBytes(privateKeyBytes))

	return &Keypair{
		PrivateKey: privateKeyBase58,
		PublicKey:  publicKeyBase58,
	}
}

func generateKeypairForAlice() *Keypair {
	return &Keypair{
		PrivateKey: []byte("HDN1ajL1k1MyNvF7ETwT7K1Jdqw2yNAx4eN3TmDXtW7J"),
		PublicKey:  []byte("AZhTRnY2QhX7rNqjPsfKDK4UJQhk9Hm6yecmpzQwRsrM"),
	}
}

func generateKeypairForBob() *Keypair {
	return &Keypair{
		PrivateKey: []byte("6aoSvdZePgsCedHmFaidWXnGVpBgyrCBqb8xobsiF67"),
		PublicKey:  []byte("Cs4byJu7ZqzTFfCnTwiQbCqDxoBgpeW3joSZ8MGQCmWg"),
	}
}

func main() {
	// Step 1. Create a keypair for alice
	alice := generateKeypairForAlice()
	// Step 2. Create a keypair for bob
	bob := generateKeypairForBob()
	fmt.Printf("Alice: %s\nBob: %s\nGeneric: %s\n", alice, bob, generateKeypair())
	// Step 3. Create Operation
	// Step 3.1 Prepare tx for Create Op
	// Step 3.1.1 Create the asset
	// Step 3.1.2 Create the inputs
	// Step 3.1.3 Create the outputs
	// Step 3.1.4 Create the metadata
	// Step 3.1.5 Create the full tx body

	// Step 3.2. Fulfill tx for Create Op
	// Sent 3.3. Sent Tx for Create Op
	// Step 4. Transfer Operation
	// Step 4.1 Prepare tx for Transfer Op
	// Step 4.2. Fulfill tx for Transfer Op
	// Sent 4.3. Sent Tx for TransferOp

}
