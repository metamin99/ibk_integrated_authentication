package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type User struct {
	userId   string
	userName string
}

type SmartContract struct {
}

func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "Create" {
		return t.Create(stub, args)
	}

	if function == "Read" {
		return t.Read(stub, args)
	}

	return shim.Success(nil)

}

func (t *SmartContract) Create(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	err := stub.PutState(args[0], []byte(args[1]))

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SmartContract) Read(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	value, err := stub.GetState(args[0])

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(value)
}

func main() {
	err := shim.Start(new(SmartContract))

	if err != nil {
		fmt.Printf("Error starting SmartContract: %s", err)
	}
}
