package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//SmartContract 정의
type SmartContract struct {
}

//User Struct 정의
type User struct {
	ObjectType     string `json:"docType"`
	UserId         string `json:"userId"`
	UserName       string `json:"userName"`
	UserPassword   string `json:"userPassword"`
	LastLoginTime  string `json:"lastLoginTime"`
	LimitLoginTime string `json:"limitLoginTime"`
	RegDate        string `json:"regDate"`
}

//logger 설정
var logger = shim.NewLogger("SmartContract")

//초기화 함수(Instantiate/Upgrade)
func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### IBKChaincode Init ###########")
	return shim.Success(nil)
}

//Invoke 함수(function collect)
func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### IBKChaincode Invoke ###########")
	function, args := stub.GetFunctionAndParameters()

	//User 생성
	if function == "CreateUser" {
		return t.CreateUser(stub, args)
	}

	//User 수정
	if function == "UpdateUser" {
		return t.UpdateUser(stub, args)
	}

	//User 삭제
	if function == "DeleteUser" {
		return t.DeleteUser(stub, args)
	}

	//User 조회 By ID
	if function == "QueryUserByUserId" {
		return t.QueryUserByUserId(stub, args)
	}

	//User 조회 전체
	if function == "QueryAllUsers" {
		return t.QueryAllUsers(stub, args)
	}

	logger.Errorf("Unknown action: %s, check the first argument, must be one of 'delete', 'query', 'echo', 'testTransient' or 'move'", args[0])
	return shim.Error(fmt.Sprintf("Unknown action: %s, check the first argument, must be one of 'delete', 'query', 'echo', 'testTransient' or 'move'", args[0]))
}

//User 생성
func (t *SmartContract) CreateUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// validate value 설정
	methodName := "[CreateUser]"
	expectedNumber := 3
	expectedPayload := "[userId, userName, userPassword]"

	// validate 함수
	if len(args) != expectedNumber {
		fmt.Printf(methodName + " Incorrect number of arguments. Expecting " + strconv.Itoa(expectedNumber) + " " + expectedPayload)
		return shim.Error(methodName + " Incorrect number of arguments. Expecting " + strconv.Itoa(expectedNumber) + " " + expectedPayload)
	}

	// user struct 생성
	key := strings.TrimSpace(args[0])
	docTypeKey := "user"
	user := User{
		ObjectType:     docTypeKey,
		UserId:         strings.TrimSpace(args[0]),
		UserName:       strings.TrimSpace(args[1]),
		UserPassword:   strings.TrimSpace(args[2]),
		LastLoginTime:  "",
		LimitLoginTime: "",
		RegDate:        time.Now().Format("20060102150405"),
	}

	// user struct -> jsonString 변환
	userJsonAsBytes, err := json.Marshal(user)

	// user struct -> jsonString 변환
	err = stub.PutState(key, userJsonAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

//User 수정
func (t *SmartContract) UpdateUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// validate value 설정
	methodName := "[UpdateUser]"
	expectedNumber := 3
	expectedPayload := "[userId, lastLoginTime, limitLoginTime]"

	// validate 함수
	if len(args) != expectedNumber {
		fmt.Printf(methodName + " Incorrect number of arguments. Expecting " + strconv.Itoa(expectedNumber) + " " + expectedPayload)
		return shim.Error(methodName + " Incorrect number of arguments. Expecting " + strconv.Itoa(expectedNumber) + " " + expectedPayload)
	}

	// 사용자 찾기
	key := strings.TrimSpace(args[0])
	stateBytes, err := stub.GetState(key)
	if err != nil {
		return shim.Error(err.Error())
	}

	// 사용자가 존재하지 않을 때 에러처리
	if len(stateBytes) == 0 {
		return shim.Error("Value is not exist ")
	}

	// 현재 로그인 정보 업데이트
	user := User{}
	json.Unmarshal(stateBytes, &user)
	user.LastLoginTime = strings.TrimSpace(args[1])
	user.LimitLoginTime = strings.TrimSpace(args[2])

	userJsonAsBytes, err := json.Marshal(user)
	if err != nil {
		return shim.Error(err.Error())
	}

	// 로그인 정보 블록체인에 저장
	err = stub.PutState(key, userJsonAsBytes)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

//User 삭제
func (t *SmartContract) DeleteUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// validate value 설정
	methodName := "[DeleteUser]"
	expectedNumber := 1
	expectedPayload := "[userId]"

	// validate 함수
	if len(args) != expectedNumber {
		fmt.Printf(methodName + " Incorrect number of arguments. Expecting " + strconv.Itoa(expectedNumber) + " " + expectedPayload)
		return shim.Error(methodName + " Incorrect number of arguments. Expecting " + strconv.Itoa(expectedNumber) + " " + expectedPayload)
	}

	// 사용자 찾기
	key := strings.TrimSpace(args[0])
	err := stub.DelState(key)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

//User 조회 by userId
func (t *SmartContract) QueryUserByUserId(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// validate value 설정
	methodName := "[QueryUserByUserId]"
	expectedNumber := 1
	expectedPayload := "[userId]"

	// validate 함수
	if len(args) != expectedNumber {
		fmt.Printf(methodName + " Incorrect number of arguments. Expecting " + strconv.Itoa(expectedNumber) + " " + expectedPayload)
		return shim.Error(methodName + " Incorrect number of arguments. Expecting " + strconv.Itoa(expectedNumber) + " " + expectedPayload)
	}

	// 현재 사용자 찾기
	key := strings.TrimSpace(args[0])

	//stateDB를 조회 - []byte 형태로 return
	stateBytes, err := stub.GetState(key)

	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(stateBytes)
}

//User 조회(전체)
func (t *SmartContract) QueryAllUsers(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	methodName := "[QueryAllUsers]"
	fmt.Printf(methodName)

	keyUsedCouchDB := fmt.Sprintf("{\"selector\":{\"docType\":\"user\"}}")

	queryResults, err := getQueryResultForQueryString(stub, keyUsedCouchDB)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func main() {
	err := shim.Start(new(SmartContract))

	if err != nil {
		fmt.Printf("Error starting SmartContract: %s", err)
	}
}

//Util 함수(Iterator)
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

//Util 함수(Iterator)
func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return &buffer, nil
}

//Util 함수(등록일자 sorting)
func sortUserByRegDate(users []User) []User {
	sort.Slice(users, func(i, j int) bool {
		return users[i].RegDate < users[j].RegDate
	})

	return users
}
