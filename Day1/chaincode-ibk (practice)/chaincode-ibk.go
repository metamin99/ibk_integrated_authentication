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

// 호출하려는 function은 무조건 Invoke 내에 있어야 함
func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "Create" {
		return t.Create(stub, args)
	}

	if function == "Read" {
		return t.Read(stub, args)
	}

	if function == "Update" {
		return t.Update(stub, args)
	}

	return shim.Success(nil)

}

// chaincode 호출시 0번째 인덱스에 function name을 쓰고, 그 다음 인덱스에 차례대로 value 정리(배열)
// ["Create", "user01", "김판석"] 또는 json 형태
// 저장할 때는 PutState
func (t *SmartContract) Create(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	err := stub.PutState(args[0], []byte(args[1]))

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// 읽을 때는 GetState
func (t *SmartContract) Read(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	value, err := stub.GetState(args[0])

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(value)
}

// main 패키지 안에는 무조건 main 메소드가 있어야 함
func main() {
	err := shim.Start(new(SmartContract))

	if err != nil {
		fmt.Printf("Error starting SmartContract: %s", err)
	}
}
