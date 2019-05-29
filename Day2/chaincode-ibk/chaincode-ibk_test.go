package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("초기화 중 오류가 발생했습니다")
		fmt.Println("\n======================error message=======================")
		fmt.Println(string(res.Message))
		fmt.Println("======================================================")
		t.FailNow()
	}
}

func checkState(t *testing.T, stub *shim.MockStub, key string, expectedValue string) {
	bytes, _ := stub.GetState(key)
	// if bytes == nil {
	// 	fmt.Println("현재 StateDB에 저장된 값이 없습니다")
	// 	t.FailNow()
	// }
	if string(bytes) != expectedValue {
		fmt.Println("\n======================error message======================")
		fmt.Println("[key]")
		fmt.Println(string(key))
		fmt.Println("\n[current state]")
		fmt.Println(string(bytes))
		fmt.Println("\n[expected current state]")
		fmt.Println(expectedValue)
		fmt.Println("======================================================")
		t.FailNow()
	}
}

func checkQuery(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)

	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	} else {
		fmt.Println("======================================================")
		fmt.Println(string(res.Payload))
		fmt.Println("======================================================")
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
}

func TestInit(t *testing.T) {
	SmartContract := new(SmartContract)
	stub := shim.NewMockStub("SmartContract", SmartContract)
	checkInit(t, stub, [][]byte{[]byte("init")})
}

// ===========================================================
//   Test_CreateUser: CreateUser Test
// ===========================================================
func Test_CreateUser(t *testing.T) {
	SmartContract := new(SmartContract)
	stub := shim.NewMockStub("SmartContract", SmartContract)

	checkInit(t, stub, [][]byte{[]byte("init")})
	now := time.Now().Format("20060102150405")
	checkInvoke(t, stub, [][]byte{[]byte("CreateUser"),
		[]byte("user01"),
		[]byte("홍길동"),
		[]byte("1q2w3e4r@")})
	checkState(t, stub, "user01", "{\"docType\":\"user\",\"userId\":\"user01\",\"userName\":\"홍길동\",\"userPassword\":\"1q2w3e4r@\",\"lastLoginTime\":\"\",\"limitLoginTime\":\"\",\"regDate\":\""+now+"\"}")
}

// ===========================================================
//   Test_UpdateUser: UpdateUser Test
// ===========================================================
func Test_UpdateUser(t *testing.T) {
	SmartContract := new(SmartContract)
	stub := shim.NewMockStub("SmartContract", SmartContract)

	checkInit(t, stub, [][]byte{[]byte("init")})
	now := time.Now().Format("20060102150405")
	checkInvoke(t, stub, [][]byte{
		[]byte("CreateUser"),
		[]byte("user01"),
		[]byte("홍길동"),
		[]byte("1q2w3e4r@")})

	checkInvoke(t, stub, [][]byte{
		[]byte("UpdateUser"),
		[]byte("user01"),
		[]byte("20190101121212"),
		[]byte("20190101121212")})

	checkState(t, stub, "user01", "{\"docType\":\"user\",\"userId\":\"user01\",\"userName\":\"홍길동\",\"userPassword\":\"1q2w3e4r@\",\"lastLoginTime\":\"20190101121212\",\"limitLoginTime\":\"20190101121212\",\"regDate\":\""+now+"\"}")
}

// ===========================================================
//   Test_DeleteUser: DeleteUser Test
// ===========================================================
func Test_DeleteUser(t *testing.T) {
	SmartContract := new(SmartContract)
	stub := shim.NewMockStub("SmartContract", SmartContract)

	checkInit(t, stub, [][]byte{[]byte("init")})
	checkInvoke(t, stub, [][]byte{
		[]byte("CreateUser"),
		[]byte("user01"),
		[]byte("홍길동"),
		[]byte("1q2w3e4r@")})

	checkInvoke(t, stub, [][]byte{
		[]byte("DeleteUser"),
		[]byte("user01")})

	checkState(t, stub, "user01", "")
}
