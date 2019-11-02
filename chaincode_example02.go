package main

import (
    "fmt"
    "time"
    "bytes"
    "strings"
    "strconv"
    "encoding/json"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/hyperledger/fabric/protos/peer"
)

// SimpleAsset implements a simple chaincode to manage an asset
type SmartContract struct {
}
type Patient struct {
 TID string `json:"tid"`
 PID string `json:"pid"`
 Subject string `json:"subject"`
 BID string `json:"bid"`
 Hash string `json:"hash"`
}

// Init is called during chaincode instantiation to initialize any
// data. Note that chaincode upgrade also calls this function to reset
// or to migrate data.
func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
    // Get the args from the transaction proposal
    args := stub.GetStringArgs()
    var check bool
    if len(args) != 5 {
            return shim.Error("Incorrect arguments, Expecting 5")
    }

    // Set up any variables or assets here by calling stub.PutState()

    // We store the key and the value on the ledger
    if len(args[4])!= 46 {
    	return shim.Error(fmt.Sprintf("hash should be 46 characters long"))
    }
    check = strings.HasPrefix(args[4], "Qm")
    if check != true {
    	return shim.Error(fmt.Sprintf("hash should start with Qm"))
    }
    //aadhar,errp := strconv.Atoi(args[0])
    //if errp != nil {
    	//return shim.Error(fmt.Sprintf("Error starting SmartContract chaincode: %s", errp))
    //}
    var data = Patient{TID: args[0], PID: args[1], Subject: args[2], BID: args[3], Hash: args[4]}
    PatientBytes, _ := json.Marshal(data)
    err := stub.PutState(args[1], PatientBytes)
    if err != nil {
            return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
    }
    return shim.Success(nil)
}

// Invoke is called per transaction on the chaincode. Each transaction is
// either a 'get' or a 'set' on the asset created by Init function. The Set
// method may create a new asset by specifying a new key-value pair.
func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
    // Extract the function and args from the transaction proposal
    fn, args := stub.GetFunctionAndParameters()

    var result string
    var err error
    if fn == "set" {
            result, err = set(stub, args)

    } else if fn == "get"{ // assume 'get' even if fn is nil
            result, err = get(stub, args)
    } else if fn == "getHistory"{
    	return getHistory(stub, args)
    } else if fn == "getall"{
    	return getall(stub)
    }
    if err != nil {
            return shim.Error(err.Error())
    }
    return shim.Success([]byte(result))

}

// Set stores the asset (both key and value) on the ledger. If the key exists,
// it will override the value with the new one
func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
    var check bool
    if len(args) != 5 {
            return "", fmt.Errorf("Incorrect arguments, Expecting 5")
    }
    if len(args[4])!=46 {
    	return "", fmt.Errorf("hash should be 46 characters long")
    }
    check = strings.HasPrefix(args[4], "Qm")
    if check != true {
    	return "", fmt.Errorf("hash should start with Qm")
    }
    //aadhar, errp := strconv.Atoi(args[0])
    //if errp != nil {
    	//return "",fmt.Errorf("Error starting SmartContract chaincode: %s", errp)
    //}
    var data = Patient{TID: args[0], PID: args[1], Subject: args[2], BID: args[3], Hash: args[4]}
    PatientBytes, _ := json.Marshal(data)
    err := stub.PutState(args[1], PatientBytes)
    if err != nil {
            return "", fmt.Errorf("Failed to set asset: %s", args[0])
    }
    return args[1], nil
}

// Get returns the value of the specified asset key
func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
    if len(args) != 1 {
            return "", fmt.Errorf("Incorrect arguments. Expecting a key")
    }

    PatientBytes, err := stub.GetState(args[0])
    if err != nil {
            return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
    }
    if PatientBytes == nil {
            return "", fmt.Errorf("Asset not found: %s", args[0])
    }
    return string(PatientBytes), nil
}
func getHistory(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments, Expecting 1")
	}
	PatientId := args[0]
	fmt.Printf("- start getHistory: %s\n", PatientId)
	resultsIterator, err := stub.GetHistoryForKey(PatientId)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	//buffer is a json array containing historic vallues for the patient
	var buffer bytes.Buffer 
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		//Add a comma before array members, suppress it for the first array member
	if bArrayMemberAlreadyWritten == true {
		buffer.WriteString(",")
	}
	buffer.WriteString("{\"TxId\":")
	buffer.WriteString("\"")
	buffer.WriteString(response.TxId)
	buffer.WriteString("\"")
	buffer.WriteString(", \"Value\":")
	//if it was a delete operation on given key, then we need to set the
	//corresponding value null.Else, we'll write the response.Value
	//as-is (as the Value itself a JSON patient)
	if response.IsDelete {
		buffer.WriteString("null")
	} else {
		buffer.WriteString(string(response.Value))
	}
	buffer.WriteString(",\"Timestamp\":")
	buffer.WriteString("\"")
	buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64 (response.Timestamp.Nanos)).String())
	buffer.WriteString("\"")
	buffer.WriteString(", \"IsDelete\":")
    buffer.WriteString("\"")
    buffer.WriteString(strconv.FormatBool(response.IsDelete))
    buffer.WriteString("\"")
    buffer.WriteString("}")
    bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	fmt.Printf("- getHistory returning:\n%s\n", buffer.String())
	return shim.Success(buffer.Bytes())
}
func  getall(stub shim.ChaincodeStubInterface) peer.Response {
    startkey := "abc1"
    endkey := "abc99"

    PatientBytes, err := stub.GetStateByRange(startkey,endkey)
    if err != nil {
		return shim.Error(err.Error())
	}
    defer PatientBytes.Close()
    var buffer bytes.Buffer
    for PatientBytes.HasNext() {
      queryResponse, err := PatientBytes.Next()
 
      if err != nil {
            return shim.Error(err.Error())
       }
      buffer.WriteString(string(queryResponse.Value))
    }
    fmt.Printf("-all query:\n%s\n",buffer.String())
    return shim.Success(buffer.Bytes())
}



// main function starts up the chaincode in the container during instantiate
func main() {
    err := shim.Start(new(SmartContract));
    if  err != nil {
            fmt.Printf("Error starting SmartContract chaincode: %s", err)
    }
}



