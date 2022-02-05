# Developer Resources

* [Code Contribution Guidelines](https://github.com/MotoAcidic/eunod/tree/master/docs/code_contribution_guidelines.md)

* [JSON-RPC Reference](https://github.com/MotoAcidic/eunod/tree/master/docs/json_rpc_api.md)
  * [RPC Examples](https://github.com/MotoAcidic/eunod/tree/master/docs/json_rpc_api.md#ExampleCode)

* The MotoAcidic Bitcoin-related Go Packages:
  * [eunorpcclient](https://github.com/MotoAcidic/eunod/tree/master/rpcclient) - Implements a
    robust and easy to use Websocket-enabled Bitcoin JSON-RPC client
  * [eunojson](https://github.com/MotoAcidic/eunod/tree/master/eunojson) - Provides an extensive API
    for the underlying JSON-RPC command and return values
  * [wire](https://github.com/MotoAcidic/eunod/tree/master/wire) - Implements the
    Bitcoin wire protocol
  * [peer](https://github.com/MotoAcidic/eunod/tree/master/peer) -
    Provides a common base for creating and managing Bitcoin network peers.
  * [blockchain](https://github.com/MotoAcidic/eunod/tree/master/blockchain) -
    Implements Bitcoin block handling and chain selection rules
  * [blockchain/fullblocktests](https://github.com/MotoAcidic/eunod/tree/master/blockchain/fullblocktests) -
    Provides a set of block tests for testing the consensus validation rules
  * [txscript](https://github.com/MotoAcidic/eunod/tree/master/txscript) -
    Implements the Bitcoin transaction scripting language
  * [eunoec](https://github.com/MotoAcidic/eunod/tree/master/eunoec) - Implements
    support for the elliptic curve cryptographic functions needed for the
    Bitcoin scripts
  * [database](https://github.com/MotoAcidic/eunod/tree/master/database) -
    Provides a database interface for the Bitcoin block chain
  * [mempool](https://github.com/MotoAcidic/eunod/tree/master/mempool) -
    Package mempool provides a policy-enforced pool of unmined bitcoin
    transactions.
  * [eunoutil](https://github.com/MotoAcidic/eunod/eunoutil) - Provides Bitcoin-specific
    convenience functions and types
  * [chainhash](https://github.com/MotoAcidic/eunod/tree/master/chaincfg/chainhash) -
    Provides a generic hash type and associated functions that allows the
    specific hash algorithm to be abstracted.
  * [connmgr](https://github.com/MotoAcidic/eunod/tree/master/connmgr) -
    Package connmgr implements a generic Bitcoin network connection manager.
