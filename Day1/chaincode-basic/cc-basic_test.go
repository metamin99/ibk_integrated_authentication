package main

import (
	"fmt"
	"testing"

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
	if bytes == nil {
		fmt.Println("현재 StateDB에 저장된 값이 없습니다")
		t.FailNow()
	}
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
//   TestInvoke_Case1: CreateUser Test
// ===========================================================
func TestInvoke_Case1(t *testing.T) {
	SmartContract := new(SmartContract)
	stub := shim.NewMockStub("SmartContract", SmartContract)

	checkInit(t, stub, [][]byte{[]byte("init")})
	checkInvoke(t, stub, [][]byte{[]byte("Create"),
		[]byte("user01"),
		[]byte("홍길동")})
	checkState(t, stub, "user01", "홍길")
}

// ===========================================================
//   TestInvoke_Case2: CreateUser Test
// ===========================================================
func TestInvoke_Case2(t *testing.T) {
	SmartContract := new(SmartContract)
	stub := shim.NewMockStub("SmartContract", SmartContract)

	checkInit(t, stub, [][]byte{[]byte("init")})
	checkInvoke(t, stub, [][]byte{[]byte("Create"),
		[]byte("user01"),
		[]byte("홍길동")})
	checkInvoke(t, stub, [][]byte{[]byte("Create"),
		[]byte("user02"),
		[]byte("유관순")})
	checkState(t, stub, "user01", "홍길동")
	checkState(t, stub, "user02", "유관순")
}