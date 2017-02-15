
Probable dependencies:
json: to serialize the transaction dictionary into a JSON formatted string;
sha3: to hash the serialized transaction; and
cryptoconditions: to create conditions and fulfillments

####################################################

Steps:
1. Create key for Alice
2. Create key for Bob
3. Prepare Tx for Create Op
4. FulFill Tx for Create Op
5. Sent Tx for Create Op
6. Prepare Tx for Transfer Op
7. Fulfill Tx for Transfer Op
8. Sent Tx for Transfer Op

####################################################
Step 1:
Alice
CryptoKeypair(
    private_key='HDN1ajL1k1MyNvF7ETwT7K1Jdqw2yNAx4eN3TmDXtW7J',
    public_key='AZhTRnY2QhX7rNqjPsfKDK4UJQhk9Hm6yecmpzQwRsrM'
)

####################################################
Step 2:
Bob
CryptoKeypair(
    private_key='6aoSvdZePgsCedHmFaidWXnGVpBgyrCBqb8xobsiF67',
    public_key='Cs4byJu7ZqzTFfCnTwiQbCqDxoBgpeW3joSZ8MGQCmWg'
)

####################################################
Step 3:
CREATE Prepared Tx:
{
    'version': '0.9',
    'id': '0fecf357f1f2c83f27ec08542ce5d92764bd4a2f85078b0d53d67456366abe3b',
    'operation': 'CREATE',

    'asset': {
        'data': {
            'bicycle': {
                'serial_number': 'abcd1234',
                'manufacturer': 'bkfab'
            }
        }
    },

    'inputs': [{
        'owners_before': ['AZhTRnY2QhX7rNqjPsfKDK4UJQhk9Hm6yecmpzQwRsrM'],
        'fulfills': None,
        'fulfillment': {
            'signature': None,
            'type_id': 4,
            'public_key': 'AZhTRnY2QhX7rNqjPsfKDK4UJQhk9Hm6yecmpzQwRsrM',
            'bitmask': 32,
            'type': 'fulfillment'
        }
    }]

    'outputs': [{
        'condition': {
            'details': {
                'signature': None,
                'type_id': 4,
                'public_key': 'AZhTRnY2QhX7rNqjPsfKDK4UJQhk9Hm6yecmpzQwRsrM',
                'bitmask': 32,
                'type': 'fulfillment'
            },
            'uri': 'cc:4:20:jhmC9jomZDEjXk7z6MVMNZCd95abFe95xvJck7jn0nY:96'
        },
        'public_keys': ['AZhTRnY2QhX7rNqjPsfKDK4UJQhk9Hm6yecmpzQwRsrM'],
        'amount': 1
    }],
    'metadata': {
        'planet': 'earth'
    },
}
####################################################
Step 4:
CREATE Fulfilled Tx:
 {
    'version': '0.9',
    'id': '0fecf357f1f2c83f27ec08542ce5d92764bd4a2f85078b0d53d67456366abe3b',
     'operation': 'CREATE',
     'asset': {
         'data': {
             'bicycle': {
                 'serial_number': 'abcd1234',
                 'manufacturer': 'bkfab'
             }
        }
    },
     'outputs': [{
         'condition': {
             'details': {
                 'signature': None,
                 'type_id': 4,
                 'public_key': 'AZhTRnY2QhX7rNqjPsfKDK4UJQhk9Hm6yecmpzQwRsrM',
                 'bitmask': 32,
                 'type': 'fulfillment'
             },
             'uri': 'cc:4:20:jhmC9jomZDEjXk7z6MVMNZCd95abFe95xvJck7jn0nY:96'
         },
         'public_keys': ['AZhTRnY2QhX7rNqjPsfKDK4UJQhk9Hm6yecmpzQwRsrM'],
         'amount': 1
     }],
     'metadata': {
         'planet': 'earth'
     },

     'inputs': [{
         'owners_before': ['AZhTRnY2QhX7rNqjPsfKDK4UJQhk9Hm6yecmpzQwRsrM'],
         'fulfills': None,
         'fulfillment': 'cf:4:jhmC9jomZDEjXk7z6MVMNZCd95abFe95xvJck7jn0nYbay0wy6ICWWnAZmem7TpgwX7worvSXPfNgw8iFIoNJtp1s5VUgPlDkE0YlW6ZhuwWtUr-XF-vvevp2CUePMoG'
     }]
 }
####################################################
Step 5:
CREATE Sent Tx:
 {
     'outputs': [{
         'condition': {
             'details': {
                 'signature': None,
                 'type_id': 4,
                 'public_key': 'AZhTRnY2QhX7rNqjPsfKDK4UJQhk9Hm6yecmpzQwRsrM',
                 'bitmask': 32,
                 'type': 'fulfillment'
             },
             'uri': 'cc:4:20:jhmC9jomZDEjXk7z6MVMNZCd95abFe95xvJck7jn0nY:96'
         },
         'public_keys': ['AZhTRnY2QhX7rNqjPsfKDK4UJQhk9Hm6yecmpzQwRsrM'],
         'amount': 1
     }],
     'version': '0.9',
     'asset': {
         'data': {
             'bicycle': {
                 'serial_number': 'abcd1234',
                 'manufacturer': 'bkfab'
             }
         }
     },
     'metadata': {
         'planet': 'earth'
     },
     'id': '0fecf357f1f2c83f27ec08542ce5d92764bd4a2f85078b0d53d67456366abe3b',
     'operation': 'CREATE',
     'inputs': [{
         'owners_before': ['AZhTRnY2QhX7rNqjPsfKDK4UJQhk9Hm6yecmpzQwRsrM'],
         'fulfills': None,
         'fulfillment': 'cf:4:jhmC9jomZDEjXk7z6MVMNZCd95abFe95xvJck7jn0nYbay0wy6ICWWnAZmem7TpgwX7worvSXPfNgw8iFIoNJtp1s5VUgPlDkE0YlW6ZhuwWtUr-XF-vvevp2CUePMoG'
     }]
 }

####################################################
Step 6:
TXFR Prepared Tx:
 {
     'outputs': [{
         'condition': {
             'details': {
                 'signature': None,
                 'type_id': 4,
                 'public_key': 'Cs4byJu7ZqzTFfCnTwiQbCqDxoBgpeW3joSZ8MGQCmWg',
                 'bitmask': 32,
                 'type': 'fulfillment'
             },
             'uri': 'cc:4:20:sEOrplnyGZgNP1t268eXBNnpJCTx_Cwkw2G7JuC3cfE:96'
         },
         'public_keys': ['Cs4byJu7ZqzTFfCnTwiQbCqDxoBgpeW3joSZ8MGQCmWg'],
         'amount': 1
     }],
     'version': '0.9',
     'asset': {
         'id': '0fecf357f1f2c83f27ec08542ce5d92764bd4a2f85078b0d53d67456366abe3b'
     },
     'metadata': None,
     'id': 'a946900369707d94650991b57ecf790cc3a2fb7a1d01a2294bef3067e61e901d',
     'operation': 'TRANSFER',
     'inputs': [{
         'owners_before': ['AZhTRnY2QhX7rNqjPsfKDK4UJQhk9Hm6yecmpzQwRsrM'],
         'fulfills': {
             'output': 0,
             'txid': '0fecf357f1f2c83f27ec08542ce5d92764bd4a2f85078b0d53d67456366abe3b'
         },
         'fulfillment': {
             'signature': None,
             'type_id': 4,
             'public_key': 'AZhTRnY2QhX7rNqjPsfKDK4UJQhk9Hm6yecmpzQwRsrM',
             'bitmask': 32,
             'type': 'fulfillment'
         }
     }]
 }

####################################################
Step 7:
TXFR Fulfilled Tx:
 {
     'outputs': [{
         'condition': {
             'details': {
                 'signature': None,
                 'type_id': 4,
                 'public_key': 'Cs4byJu7ZqzTFfCnTwiQbCqDxoBgpeW3joSZ8MGQCmWg',
                 'bitmask': 32,
                 'type': 'fulfillment'
             },
             'uri': 'cc:4:20:sEOrplnyGZgNP1t268eXBNnpJCTx_Cwkw2G7JuC3cfE:96'
         },
         'public_keys': ['Cs4byJu7ZqzTFfCnTwiQbCqDxoBgpeW3joSZ8MGQCmWg'],
         'amount': 1
     }],
     'version': '0.9',
     'asset': {
         'id': '0fecf357f1f2c83f27ec08542ce5d92764bd4a2f85078b0d53d67456366abe3b'
     },
     'metadata': None,
     'id': 'a946900369707d94650991b57ecf790cc3a2fb7a1d01a2294bef3067e61e901d',
     'operation': 'TRANSFER',
     'inputs': [{
         'owners_before': ['AZhTRnY2QhX7rNqjPsfKDK4UJQhk9Hm6yecmpzQwRsrM'],
         'fulfills': {
             'output': 0,
             'txid': '0fecf357f1f2c83f27ec08542ce5d92764bd4a2f85078b0d53d67456366abe3b'
         },
         'fulfillment': 'cf:4:jhmC9jomZDEjXk7z6MVMNZCd95abFe95xvJck7jn0nbaeI8RYxGmJiesp9DncKD7Pb5SzuCog8IctHzTCq6dyhFfmGWJ_4oR18Ogcd83Jc7QTlLKGbZm-jZyNV4kRh0N'
     }]
 }

####################################################
Step 8:
TXFR Sent Tx:
 {
     'outputs': [{
         'condition': {
             'details': {
                 'signature': None,
                 'type_id': 4,
                 'public_key': 'Cs4byJu7ZqzTFfCnTwiQbCqDxoBgpeW3joSZ8MGQCmWg',
                 'bitmask': 32,
                 'type': 'fulfillment'
             },
             'uri': 'cc:4:20:sEOrplnyGZgNP1t268eXBNnpJCTx_Cwkw2G7JuC3cfE:96'
         },
         'public_keys': ['Cs4byJu7ZqzTFfCnTwiQbCqDxoBgpeW3joSZ8MGQCmWg'],
         'amount': 1
     }],
     'version': '0.9',
     'asset': {
         'id': '0fecf357f1f2c83f27ec08542ce5d92764bd4a2f85078b0d53d67456366abe3b'
     },
     'metadata': None,
     'id': 'a946900369707d94650991b57ecf790cc3a2fb7a1d01a2294bef3067e61e901d',
     'operation': 'TRANSFER',
     'inputs': [{
         'owners_before': ['AZhTRnY2QhX7rNqjPsfKDK4UJQhk9Hm6yecmpzQwRsrM'],
         'fulfills': {
             'output': 0,
             'txid': '0fecf357f1f2c83f27ec08542ce5d92764bd4a2f85078b0d53d67456366abe3b'
         },
         'fulfillment': 'cf:4:jhmC9jomZDEjXk7z6MVMNZCd95abFe95xvJck7jn0nbaeI8RYxGmJiesp9DncKD7Pb5SzuCog8IctHzTCq6dyhFfmGWJ_4oR18Ogcd83Jc7QTlLKGbZm-jZyNV4kRh0N'
     }]
 }

####################################################



Rough goals:
- ease of use
- form tx struct, get json and/or protobuf
- driverhandles failover between nodes in a federation
- federation discovery using nslookup?
- maintain state?
- support the event API with a long-lived connection?
- https and http. spdy maybe?
- fast
- initial version handshake?

