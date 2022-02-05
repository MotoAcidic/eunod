module github.com/MotoAcidic/eunod/eunoutil/psbt

go 1.17

require (
	github.com/MotoAcidic/eunod v0.22.0-beta.0.20220111032746-97732e52810c
	github.com/MotoAcidic/eunod/eunoec/v2 v2.1.0
	github.com/MotoAcidic/eunod/eunoutil v1.0.0
	github.com/davecgh/go-spew v1.1.1
)

require (
	github.com/MotoAcidic/eunolog v0.0.0-20170628155309-84c8d2346e9f // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
)

replace github.com/MotoAcidic/eunod/eunoec/v2 => ../../eunoec

replace github.com/MotoAcidic/eunod/eunoutil => ../

replace github.com/MotoAcidic/eunod => ../..
