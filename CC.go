package main

import "github.com/hyperledger/fabric/core/chaincode/shim"

type Chaincode struct {
}

type Asset struct {
}

type UserInfo struct {
	Maker    string
	AllAsset []Asset
}

func (t *Chaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	return nil, nil
}

func (t *Chaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	return nil, nil
}

func (t *Chaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	Avalbytes, _ := stub.GetState(args[0])

	return Avalbytes, nil
}

func main() {
	shim.Start(new(Chaincode))
}
