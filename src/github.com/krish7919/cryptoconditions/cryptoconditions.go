package cryptoconditions

//package main

import (
	"bytes"
	"encoding/asn1"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ed25519"
)

type ConditionType int

const (
	PREIMAGE_SHA256 ConditionType = iota
	RSA_SHA256
	PREFIX_SHA256
	THRESHOLD_SHA256
	ED25519
)

const (
	FeatureBitmask_SHA256    int = 0x01
	FeatureBitmask_PREIMAGE  int = 0x02
	FeatureBitmask_PREFIX    int = 0x04
	FeatureBitmask_THRESHOLD int = 0x08
	FeatureBitmask_RSA_PSS   int = 0x10
	FeatureBitmask_ED25519   int = 0x20
)

// Start of ASN.1 Encoding for CryptoConditions

type ConditionASN1 []struct {
	Type           ConditionType `asn1:"type"`
	FeatureBitmask asn1.RawValue `asn1:"featureBitmask"`
	// Fingerprint is the public key
	Fingerprint          asn1.RawValue `asn1:"fingerprint"`
	MaxFulfillmentLength int64         `asn1:"maxFulfillmentLength"`
}

type FulfillmentASN1 []struct {
	Type    ConditionASN1 `asn1:"type"`
	Payload asn1.RawValue `asn1:"payload"`
}

type PrefixSha256FulfillmentPayloadASN1 []struct {
	Prefix         asn1.RawValue   `asn1:"prefix"`
	Subfulfillment FulfillmentASN1 `asn1:"subfulfillment"`
}

type ThresholdSha256FulfillmentPayloadASN1 []struct {
	threshold       int32                         `asn1:"threshold"`
	Subfulfillments []ThresholdSubfulfillmentASN1 `asn1:"subfulfillments"`
}

type ThresholdSubfulfillmentASN1 []struct {
	Weight      int32           `asn1:"weight,default:1"`
	Condition   ConditionASN1   `asn1:"condition,optional"`
	Fulfillment FulfillmentASN1 `asn1:"fulfillment,optional"`
}

type RsaSha256FulfillmentPayloadASN1 []struct {
	Modulus   asn1.RawValue `asn1:"modulus"`   // (SIZE(128..512)), TODO
	Signature asn1.RawValue `asn1:"signature"` // (SIZE(128..512)) TODO
}

type Ed25519FulfillmentPayloadASN1 []struct {
	PublicKey asn1.RawValue `asn1:"publicKey"` //(SIZE(32)), TODO
	Signature asn1.RawValue `asn1:"signature"` // (SIZE(64)) TODO
}

type Sha256FingerprintASN1 asn1.RawValue  // (SIZE(32)) -- digest TODO
type Ed25519FingerprintASN1 asn1.RawValue // (SIZE(32)) -- publicKey

type PrefixSha256FingerprintContentsASN1 []struct {
	Prefix    asn1.RawValue `asn1:"prefix"`
	Condition ConditionASN1 `asn1:"condition"`
}

type ThresholdSha256FingerprintContentsASN1 []struct {
	Threshold     int32                       `asn1:"threshold"`
	Subconditions []ThresholdSubconditionASN1 `asn1:"subconditions"`
}

type ThresholdSubconditionASN1 []struct {
	Weight    int32         `asn1:"weight"`
	Condition ConditionASN1 `asn1:"condition"`
}

type RsaSha256FingerprintContentsASN1 int64

// End of ASN.1 Encoding for CryptoConditions

// Start of ASCII Encoding for CryptoConditions

type ConditionASCII []struct {
	Type                 ConditionType `json:"type"`
	FeatureBitmask       []byte        `json:"bitmask"`
	Fingerprint          []byte        `json:"public_key"`
	MaxFulfillmentLength int64         `json:"maxFulfillmentLength"`
}

type FulfillmentASCII []struct {
	Type    ConditionASCII `json:"type"`
	Payload []byte         `json:"payload"`
}

type PrefixSha256FulfillmentPayloadASCII []struct {
	Prefix         []byte           `json:"prefix"`
	Subfulfillment FulfillmentASCII `json:"subfulfillment"`
}

type ThresholdSha256FulfillmentPayloadASCII []struct {
	Threshold       int32                          `json:"threshold"`
	Subfulfillments []ThresholdSubfulfillmentASCII `json:"subfulfillments"`
}

type ThresholdSubfulfillmentASCII []struct {
	Weight      int32            `json:"weight,default:1"`
	Condition   ConditionASCII   `json:"condition,optional"`
	Fulfillment FulfillmentASCII `json:"fulfillment,optional"`
}

type RsaSha256FulfillmentPayloadASCII []struct {
	Modulus   []byte `json:"modulus"`   // (SIZE(128..512)), TODO
	Signature []byte `json:"signature"` // (SIZE(128..512)) TODO
}

type Ed25519FulfillmentPayloadASCII []struct {
	PublicKey []byte `json:"publicKey"` //(SIZE(32)), TODO
	Signature []byte `json:"signature"` // (SIZE(64)) TODO
}

type Sha256FingerprintASCII []byte  // (SIZE(32)) -- digest TODO
type Ed25519FingerprintASCII []byte // (SIZE(32)) -- publicKey

type PrefixSha256FingerprintContentsASCII []struct {
	Prefix    []byte         `json:"prefix"`
	Condition ConditionASCII `json:"condition"`
}

type ThresholdSha256FingerprintContentsASCII []struct {
	Threshold     int32                        `json:"threshold"`
	Subconditions []ThresholdSubconditionASCII `json:"subconditions"`
}

type ThresholdSubconditionASCII []struct {
	Weight    int32          `json:"weight"`
	Condition ConditionASCII `json:"condition"`
}

type RsaSha256FingerprintContentsASCII int64

func (cc *ConditionASCII) String() string {
	var retVal bytes.Buffer
	for _, c := range cc {
		retVal.WriteString(
			fmt.Sprintf("cc:%s:%s:%s:%s",
				cc.Type, cc.FeatureBitMask, cc.Fingerprint,
				cc.MaxFulfillmentLength))
	}
	return retVal.String()
}

func (cf *FulfillmentASCII) String() string {
	return fmt.Sprintf("cf:%s:%s", cf.FeatureType, cf.Payload)
}

// End of ASCII Encoding for CryptoConditions

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

//// End of cryptoCons

// Returns a pair of (private_key, public_key) encoded in base58.
func GenerateKeypair() (string, string) {
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

	publicKeyBase58 := base58.Encode(publicKeyBytes)
	privateKeyBase58 := base58.Encode(privateKeyBytes)

	return privateKeyBase58, publicKeyBase58
}

/*

 22     public static byte[] getSha256Hash(byte[] input) {
 23         try {
 24             MessageDigest digest = MessageDigest.getInstance("SHA-256");
 25             return digest.digest(input);
 26         } catch (NoSuchAlgorithmException e) {
 27             throw new IllegalArgumentException(e);
 28         }
 29     }
 30
 31 }

    public void writeFulfillment(Fulfillment fulfillment) throws IOException {
 36         writeConditionType(fulfillment.getType());
 37         writePayload(fulfillment.getPayload().payload);
 38
 39     }
 40
 41     public void writeConditionType(ConditionType type)
 42             throws IOException {
 43         write16BitUInt(type.getTypeCode());
 44     }
 45
 46     protected void writePayload(byte[] payload)
 47             throws IOException {
 48         writeOctetString(payload);
 49     }

*/

// TODO func main() {}

/*
	`json:"version"`
	`protobuf:"varint,1,opt,name=amount" json:"amount,omitempty"`*/
