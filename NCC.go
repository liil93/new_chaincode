package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type CC struct {
}

type User struct {
	ScrtKey  string
	Contact  string
	AstList  string
	TrctList string
}
type Asset struct {
	UserKey   string
	ScrtKey   string
	Type      string
	Locate    string
	StartDate string
	EndDate   string
	Except    string
}
type Transaction struct {
	ProducerKey string
	ConsumerKey string
	Type        string
	StartTime   string
	EndTime     string
	Cost        string
}

type Update struct {
	Time     int
	Previous string
	Current  string
}

var _Locate map[string]string
var _Update Update

// ==================================================================================================
// main
// ==================================================================================================
func main() {
	err := shim.Start(new(CC))
	if err != nil {
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("                                 ＜ Main ＞")
		fmt.Println("                         Error starting chaincode")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
	}
}

// ==================================================================================================
// Init
// ==================================================================================================
func (t *CC) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 0 {
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("                               ＜ Init ＞")
		fmt.Println("              Incorrect number of arguments. Expecting 0")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		return nil, errors.New("[Init] Incorrect number of arguments. Expecting 0")
	}
	_Locate = make(map[string]string)
	fmt.Println()
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("                               ＜ Init ＞")
	fmt.Println("                              Init Success")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	return nil, nil
}

// ==================================================================================================
// Invoke - Our entry point for Invocations
// ==================================================================================================
func (t *CC) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "UserRegist" {
		return t.UserRegist(stub, args)
	} else if function == "UserChangeContact" {
		return t.UserChangeContact(stub, args)
	} else if function == "AssetRegist" {
		return t.AssetRegist(stub, args)
	} else if function == "AssetChange" {
		return t.AssetChange(stub, args)
	} else if function == "AssetDelete" {
		return t.AssetDelete(stub, args)
	} else if function == "TransactionRegist" {
		return t.TransactionRegist(stub, args)
	} else if function == "Test" {
		return t.Test(stub, args)
	}

	fmt.Println()
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("                               ＜ Invoke ＞")
	fmt.Println("               Invoke did not find func: " + function)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	return nil, errors.New("[Invoke] Invoke did not find func: " + function)
}

// ==================================================================================================
// Query - Our entry point for Queries
// ==================================================================================================
func (t *CC) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "UserRead" {
		return t.UserRead(stub, args)
	} else if function == "AssetRead" {
		return t.AssetRead(stub, args)
	} else if function == "TransactionRead" {
		return t.TransactionRead(stub, args)
	} else if function == "LocateSearch" {
		return t.LocateSearch(stub, args)
	} else if function == "GetUpdate" {
		return t.GetUpdate(stub, args)
	}
	fmt.Println()
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("                               ＜ Query ＞")
	fmt.Println("               Query did not find func: " + function)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	return []byte("Query did not find func: " + function), errors.New("[Query] Query did not find func: " + function)
}

// ==================================================================================================
// UserRegist
// ==================================================================================================
func (t *CC) UserRegist(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 3 {
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("                              ＜ UserRegist ＞")
		fmt.Println("              Incorrect number of arguments. Expecting 3")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		return nil, errors.New("[UserRegist] Incorrect number of arguments. Expecting 3")
	}
	for _, v := range args[0] {
		if v == 95 || v == 35 {
			fmt.Println()
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println("                              ＜ UserRegist ＞")
			fmt.Println("                             Incorrect userkey")
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println()
			return nil, errors.New("[UserRegist] Incorrect userkey")
		}
	}
	conf, _ := stub.GetState(args[0])
	if conf != nil {
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("                              ＜ UserRegist ＞")
		fmt.Println("                           Already exist userkey")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		return nil, errors.New("[UserRegist] Already exist userkey")
	}
	user := User{args[1], args[2], "", ""}
	userByte, _ := json.Marshal(user)
	stub.PutState(args[0], userByte)
	fmt.Println()
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("                              ＜ UserRegist ＞")
	fmt.Println("                              Regist success")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	return nil, nil
}

// ==================================================================================================
// UserChangeContact
// ==================================================================================================
func (t *CC) UserChangeContact(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 3 {
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("                          ＜ UserChangeContact ＞")
		fmt.Println("              Incorrect number of arguments. Expecting 3")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		return nil, errors.New("[UserChangeContact] Incorrect number of arguments. Expecting 3")
	}
	for _, v := range args[0] {
		if v == 95 || v == 35 {
			fmt.Println()
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println("                          ＜ UserChangeContact ＞")
			fmt.Println("                             Incorrect userkey")
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println()
			return nil, errors.New("[UserChangeContact] Incorrect userkey")
		}
	}
	conf, _ := stub.GetState(args[0])
	if conf == nil {
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("                          ＜ UserChangeContact ＞")
		fmt.Println("                             Not exist userkey")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		return nil, errors.New("[UserChangeContact] Not exist userkey")
	}
	user := User{}
	json.Unmarshal(conf, &user)
	if user.ScrtKey != args[1] {
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("                          ＜ UserChangeContact ＞")
		fmt.Println("                           Incorrect secretkey")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		return nil, errors.New("[UserChangeContact] Incorrect secretkey")
	}
	user.Contact = args[2]
	userByte, _ := json.Marshal(user)
	stub.PutState(args[0], userByte)
	fmt.Println()
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("                          ＜ UserChangeContact ＞")
	fmt.Println("                          Contact change success")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	return nil, nil
}

// ==================================================================================================
// UserRead
// ==================================================================================================
func (t *CC) UserRead(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("                             ＜ UserRead ＞")
		fmt.Println("              Incorrect number of arguments. Expecting 1")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		return []byte("Incorrect number of arguments. Expecting 1"), errors.New("[UserRead] Incorrect number of arguments. Expecting 1")
	}
	for _, v := range args[0] {
		if v == 95 || v == 35 {
			fmt.Println()
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println("                             ＜ UserRead ＞")
			fmt.Println("                           Incorrect userkey")
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println()
			return []byte("Incorrect userkey"), errors.New("[UserRead] Incorrect userkey")
		}
	}
	conf, _ := stub.GetState(args[0])
	if conf == nil {
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("                             ＜ UserRead ＞")
		fmt.Println("                           Not exist userkey")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		return []byte("Not exist userkey"), errors.New("[UserRead] Not exist userkey")
	}
	user := User{}
	json.Unmarshal(conf, &user)
	user.ScrtKey = "unknown"
	conf, _ = json.Marshal(user)
	fmt.Println()
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("                             ＜ UserRead ＞")
	fmt.Println("                            Reading success")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	return conf, nil
}

// ==================================================================================================
// AssetRegist
// ==================================================================================================
func (t *CC) AssetRegist(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 4 {
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("                           ＜ AssetRegist ＞")
		fmt.Println("              Incorrect number of arguments. Expecting 4")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		return nil, errors.New("[AssetRegist] Incorrect number of arguments. Expecting 4")
	}
	for _, v := range args[0] {
		if v == 95 || v == 35 {
			fmt.Println()
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println("                           ＜ AssetRegist ＞")
			fmt.Println("                           Incorrect userkey")
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println()
			return nil, errors.New("[AssetRegist] Incorrect userkey")
		}
	}
	conf, _ := stub.GetState(args[0])
	if conf == nil {
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("                           ＜ AssetRegist ＞")
		fmt.Println("                           Not exist userkey")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		return nil, errors.New("[AssetRegist] Not exist userkey")
	}
	user := User{}
	json.Unmarshal(conf, &user)
	if user.ScrtKey != args[1] {
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("                           ＜ AssetRegist ＞")
		fmt.Println("                          Incorrect secretkey")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		return nil, errors.New("[AssetRegist] Incorrect secretkey")
	}
	asset := Asset{}
	asset.UserKey = args[0]
	asset.ScrtKey = args[1]
	asset.Type = args[2]
	asset.Locate = args[3]
	assetByte, _ := json.Marshal(asset)
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := rand.Int() % 10000; ; i++ {
		confkey, _ := stub.GetState(strconv.Itoa(i) + "#" + args[2])
		if confkey == nil {
			key := strconv.Itoa(i) + "#" + args[2]
			stub.PutState(key, assetByte)
			user.AstList = user.AstList + key + "/"
			userByte, _ := json.Marshal(user)
			stub.PutState(args[0], userByte)
			_Locate[args[3]] = _Locate[args[3]] + key + "/"

			timeNow := time.Now().Minute()
			if timeNow < 10 {
				timeNow = 0
			} else if timeNow < 20 {
				timeNow = 10
			} else if timeNow < 30 {
				timeNow = 20
			} else if timeNow < 40 {
				timeNow = 30
			} else if timeNow < 50 {
				timeNow = 40
			} else if timeNow < 60 {
				timeNow = 50
			}
			if _Update.Time != timeNow {
				_Update.Previous = _Update.Current
				_Update.Current = key + "_R" + "/"
				_Update.Time = timeNow
			} else {
				_Update.Current = _Update.Current + key + "_R" + "/"
			}

			fmt.Println()
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println("                           ＜ AssetRegist ＞")
			fmt.Println("                            Regist success")
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println()
			return nil, nil
		}
	}
}

// ==================================================================================================
// AssetChange
// ==================================================================================================
func (t *CC) AssetChange(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 5 && len(args) != 6 {
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("                             ＜ AssetChange ＞")
		fmt.Println("              Incorrect number of arguments. Expecting 5 or 6")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		return nil, errors.New("[AssetChange] Incorrect number of arguments. Expecting 5 or 6")
	}
	for _, v := range args[0] {
		if v == 35 {
			conf, _ := stub.GetState(args[0])
			if conf == nil {
				fmt.Println()
				fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
				fmt.Println("                             ＜ AssetChange ＞")
				fmt.Println("                            Not exist assetkey")
				fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
				fmt.Println()
				return nil, errors.New("[AssetChange] Not exist assetkey")
			}
			asset := Asset{}
			json.Unmarshal(conf, &asset)
			if asset.ScrtKey != args[1] {
				fmt.Println()
				fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
				fmt.Println("                             ＜ AssetChange ＞")
				fmt.Println("                           Incorrect secretkey")
				fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
				fmt.Println()
				return nil, errors.New("[AssetChange] Incorrect secretkey")
			}
			asset.StartDate = args[2]
			asset.EndDate = args[3]
			asset.Except = args[4]
			if len(args) == 6 {
				var start, end int
				loc := asset.Locate
				for i, v := range _Locate[loc] {
					if v == 47 {
						end = i
						if _Locate[loc][start:end+1] == args[0]+"/" {
							if len(_Locate[loc]) == end+1 && start == 0 {
								_Locate[loc] = ""
							} else if len(_Locate[loc]) == end+1 && start != 0 {
								_Locate[loc] = _Locate[loc][:start]
							} else if len(_Locate[loc]) != end+1 && start == 0 {
								_Locate[loc] = _Locate[loc][end+1:]
							} else if len(_Locate[loc]) != end+1 && start != 0 {
								_Locate[loc] = _Locate[loc][:start] + _Locate[loc][end+1:]
							}
							break
						}
						start = end + 1
					}
				}
				asset.Locate = args[5]
				_Locate[args[5]] = _Locate[args[5]] + args[0] + "/"
			}
			assetByte, _ := json.Marshal(asset)
			stub.PutState(args[0], assetByte)

			timeNow := time.Now().Minute()
			if timeNow < 10 {
				timeNow = 0
			} else if timeNow < 20 {
				timeNow = 10
			} else if timeNow < 30 {
				timeNow = 20
			} else if timeNow < 40 {
				timeNow = 30
			} else if timeNow < 50 {
				timeNow = 40
			} else if timeNow < 60 {
				timeNow = 50
			}
			if _Update.Time != timeNow {
				_Update.Previous = _Update.Current
				_Update.Current = args[0] + "_C" + "/"
				_Update.Time = timeNow
			} else {
				_Update.Current = _Update.Current + args[0] + "_C" + "/"
			}

			fmt.Println()
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println("                             ＜ AssetChange ＞")
			fmt.Println("                           Asset change success")
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println()
			return nil, nil
		}
	}
	fmt.Println()
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("                              ＜ AssetRead ＞")
	fmt.Println("                            Incorrect assetkey")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	return nil, errors.New("[AssetChange] Incorrect assetkey")
}

// ==================================================================================================
// AssetDelete
// ==================================================================================================
func (t *CC) AssetDelete(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 2 {
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("                             ＜ AssetDelete ＞")
		fmt.Println("              Incorrect number of arguments. Expecting 2")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		return nil, errors.New("[AssetDelete] Incorrect number of arguments. Expecting 2")
	}
	for _, v := range args[0] {
		if v == 35 {
			conf, _ := stub.GetState(args[0])
			if conf == nil {
				fmt.Println()
				fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
				fmt.Println("                             ＜ AssetDelete ＞")
				fmt.Println("                            Not exist assetkey")
				fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
				fmt.Println()
				return nil, errors.New("[AssetDelete] Not exist assetkey")
			}
			asset := Asset{}
			json.Unmarshal(conf, &asset)
			if asset.ScrtKey != args[1] {
				fmt.Println()
				fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
				fmt.Println("                             ＜ AssetDelete ＞")
				fmt.Println("                           Incorrect secretkey")
				fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
				fmt.Println()
				return nil, errors.New("[AssetDelete] Incorrect secretkey")
			}

			user := User{}
			var start, end int
			confuser, _ := stub.GetState(asset.UserKey)
			json.Unmarshal(confuser, &user)
			for i, v := range user.AstList {
				if v == 47 {
					end = i
					if user.AstList[start:end+1] == args[0]+"/" {
						if len(user.AstList) == end+1 && start == 0 {
							user.AstList = ""
						} else if len(user.AstList) == end+1 && start != 0 {
							user.AstList = user.AstList[:start]
						} else if len(user.AstList) != end+1 && start == 0 {
							user.AstList = user.AstList[end+1:]
						} else if len(user.AstList) != end+1 && start != 0 {
							user.AstList = user.AstList[:start] + user.AstList[end+1:]
						}
						break
					}
					start = end + 1
				}
			}
			userByte, _ := json.Marshal(user)
			stub.PutState(asset.UserKey, userByte)
			loc := asset.Locate
			for i, v := range _Locate[loc] {
				if v == 47 {
					end = i
					if _Locate[loc][start:end+1] == args[0]+"/" {
						if len(_Locate[loc]) == end+1 && start == 0 {
							_Locate[loc] = ""
						} else if len(_Locate[loc]) == end+1 && start != 0 {
							_Locate[loc] = _Locate[loc][:start]
						} else if len(_Locate[loc]) != end+1 && start == 0 {
							_Locate[loc] = _Locate[loc][end+1:]
						} else if len(_Locate[loc]) != end+1 && start != 0 {
							_Locate[loc] = _Locate[loc][:start] + _Locate[loc][end+1:]
						}
						break
					}
					start = end + 1
				}
			}
			stub.DelState(args[0])

			timeNow := time.Now().Minute()
			if timeNow < 10 {
				timeNow = 0
			} else if timeNow < 20 {
				timeNow = 10
			} else if timeNow < 30 {
				timeNow = 20
			} else if timeNow < 40 {
				timeNow = 30
			} else if timeNow < 50 {
				timeNow = 40
			} else if timeNow < 60 {
				timeNow = 50
			}
			if _Update.Time != timeNow {
				_Update.Previous = _Update.Current
				_Update.Current = args[0] + "_D" + "/"
				_Update.Time = timeNow
			} else {
				_Update.Current = _Update.Current + args[0] + "_D" + "/"
			}

			fmt.Println()
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println("                             ＜ AssetDelete ＞")
			fmt.Println("                           Asset delete success")
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println()
			return nil, nil
		}
	}
	fmt.Println()
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("                             ＜ AssetDelete ＞")
	fmt.Println("                            Incorrect assetkey")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	return nil, errors.New("[AssetDelete] Incorrect assetkey")
}

// ==================================================================================================
// AssetRead
// ==================================================================================================
func (t *CC) AssetRead(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("                              ＜ AssetRead ＞")
		fmt.Println("              Incorrect number of arguments. Expecting 1")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		return []byte("Incorrect number of arguments. Expecting 1"), errors.New("[AssetRead] Incorrect number of arguments. Expecting 1")
	}
	for _, v := range args[0] {
		if v == 35 {
			conf, _ := stub.GetState(args[0])
			if conf == nil {
				fmt.Println()
				fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
				fmt.Println("                              ＜ AssetRead ＞")
				fmt.Println("                            Not exist assetkey")
				fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
				fmt.Println()
				return []byte("Not exist assetkey"), errors.New("[AssetRead] Not exist assetkey")
			}
			asset := Asset{}
			json.Unmarshal(conf, &asset)
			asset.ScrtKey = "unknown"
			conf, _ = json.Marshal(asset)
			fmt.Println()
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println("                              ＜ AssetRead ＞")
			fmt.Println("                              Reading success")
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println()
			return conf, nil
		}
	}
	fmt.Println()
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("                              ＜ AssetRead ＞")
	fmt.Println("                             Incorrect assetkey")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	return []byte("Incorrect assetkey"), errors.New("[AssetRead] Incorrect assetkey")
}

// ==================================================================================================
// TransactionRegist
// ==================================================================================================
func (t *CC) TransactionRegist(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 7 {
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("                        ＜ TransactionRegist ＞")
		fmt.Println("               Incorrect number of arguments. Expecting 7")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		return nil, errors.New("[TransactionRegist] Incorrect number of arguments. Expecting 7")
	}
	for _, v := range args[0] {
		if v == 95 || v == 35 {
			fmt.Println()
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println("                           ＜ AssetRegist ＞")
			fmt.Println("                           Incorrect userkey")
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println()
			return nil, errors.New("[TransactionRegist] Incorrect userkey")
		}
	}
	conf, _ := stub.GetState(args[0])
	if conf == nil {
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("                        ＜ TransactionRegist ＞")
		fmt.Println("                           Not exist userkey")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		return nil, errors.New("[TransactionRegist] Not exist userkey")
	}
	user := User{}
	json.Unmarshal(conf, &user)
	if user.ScrtKey != args[1] {
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("                        ＜ TransactionRegist ＞")
		fmt.Println("                          Incorrect secretkey")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		return nil, errors.New("[TransactionRegist] Incorrect secretkey")
	}
	transaction := Transaction{}
	transaction.ProducerKey = args[0]
	transaction.ConsumerKey = args[2]
	transaction.Type = args[3]
	transaction.StartTime = args[4]
	transaction.EndTime = args[5]
	transaction.Cost = args[6]

	transactionByte, _ := json.Marshal(transaction)
	key := args[0] + "_" + args[2] + "_" + args[3] + "_" + args[4]
	stub.PutState(key, transactionByte)
	user.TrctList = user.TrctList + key + "/"
	userByte, _ := json.Marshal(user)
	stub.PutState(args[0], userByte)

	fmt.Println()
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("                        ＜ TransactionRegist ＞")
	fmt.Println("                            Regist success")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	return nil, nil
}

// ==================================================================================================
// TransactionRead
// ==================================================================================================
func (t *CC) TransactionRead(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 4 {
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("                           ＜ TransactionRead ＞")
		fmt.Println("              Incorrect number of arguments. Expecting 4")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		return []byte("Incorrect number of arguments. Expecting 4"), errors.New("[TransactionRead] Incorrect number of arguments. Expecting 4")
	}
	key := args[0] + "_" + args[1] + "_" + args[2] + "_" + args[3]
	conf, _ := stub.GetState(key)
	if conf == nil {
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("                           ＜ TransactionRead ＞")
		fmt.Println("                           Not exist transaction")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		return []byte("Not exist transaction"), errors.New("[TransactionRead] Not exist transaction")
	}

	return conf, nil
}

// ==================================================================================================
// LocateSearch
// ==================================================================================================
func (t *CC) LocateSearch(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("                             ＜ LocateSearch ＞")
		fmt.Println("              Incorrect number of arguments. Expecting 1")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		return []byte("Incorrect number of arguments. Expecting 1"), errors.New("[LocateSearch] Incorrect number of arguments. Expecting 1")
	}
	fmt.Println()
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("                             ＜ LocateSearch ＞")
	fmt.Println("                              Reading success")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	if _Locate[args[0]] == "" {
		return []byte("Not exist asset..."), nil
	}
	return []byte(_Locate[args[0]]), nil
}

// ==================================================================================================
// GetUpdate
// ==================================================================================================
func (t *CC) GetUpdate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if _Update.Previous == "" {
		return []byte("No update record"), nil
	}
	return []byte(_Update.Previous), nil
}

func (t *CC) Test(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	for i := 0; i < 20; i++ {
		stub.PutState(strconv.Itoa(i), []byte(strconv.Itoa(i)+"T"))
	}

	return nil, nil
}
