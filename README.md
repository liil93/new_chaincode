# Hyperledger Fabric - New Chaincode

## 개요
유저를 등록하여 공유가 가능한 자산을 추가, 수정, 삭제하고 위치검색과 자산의 업데이트 기록을
10분단위로 조회할 수 있도록 한 코드입니다.  


체인 코드는 다음과 같은 기능을 제공합니다.  

#### **chaincode - invoke**
1. UserRegist:유저를 등록합니다.
2. UserChangeContact:유저의 연락처를 수정합니다.
3. AssetRegist:유저의 자산을 등록합니다.
4. AssetChange:유저의 자산을 변경합니다.
5. AssetDelete:유저의 자산을 삭제합니다.
6. TransactionRegist:거래 기록을 등록합니다.

#### **chaincode - query**
1. UserRead:유저의 현재 상태를 조회합니다.
2. AssetRead:유저의 자산정보를 조회합니다.
3. TransactionRead:거래 기록을 조회합니다.
4. LocateSearch:지역에 존재하는 자산들을 조회합니다.
5. GetUpdate:이전 최근 10분간의 자산 업데이트기록을 조회합니다.

### *membersrvc*
##### localhost:7050/registrar

    {
        "enrollId": "admin",
        "enrollSecret": "Xurw3yU9zI0l"
    }

admin으로 로그인하여 아래와 같은 문장이 뜨면 됩니다.

    {
      "OK": "Login successful for user 'admin'."
    }

### *init(deploy)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "deploy",
      "params": {
        "type": 1,
        "chaincodeID":{
            "name": "mycc"
        },
        "ctorMsg": {
            "args":[""]
        },
        "secureContext": "admin"
      },
      "id": 1
    }

mycc라는 체인코드를 deploy하여 줍니다. 이 때 args로는 아무것도 입력하지 않습니다.  

### *UserRegist(invoke)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "invoke",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["UserRegist", "userID", "sceretkey", "contact"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }


### *UserChangeContact(invoke)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "invoke",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["UserChangeContact", "userID", "sceretkey", "contact"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }


### *AssetRegist(invoke)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "invoke",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["AssetRegist", "userID", "sceretkey", "type", "locate"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }


### *AssetChange(invoke)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "invoke",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["AssetChange", "userID", "sceretkey", "starttime", "endtime", "except"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }

만약 위치도 바꾸고싶으면 "except" 뒤에 변경할 지역을 추가해주면 됩니다.  
예: "args":["AssetChange", "userID", "sceretkey", "starttime", "endtime", "except", "jeju"]

### *AssetDelete(invoke)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "invoke",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["AssetDelete", "userID", "sceretkey"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }


### *TransactionRegist(invoke)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "invoke",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["TransactionRegist", "producer", "secretkey", "consumer", "type", "starttime", "endtime", "cost"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }



### *UserRead(query)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "query",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["UserRead","userID"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }

### *AssetRead(query)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "query",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["AssetRead","assetkey"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }


### *TransactionRead(query)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "query",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["TransactionRead","producer", "consumer", "type", "starttime"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }


### *LocateSearch(query)*
##### localhost:7050/chaincode


    {
      "jsonrpc": "2.0",
      "method": "query",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["LocateSearch","seoul"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }

### *GetUpdate(query)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "query",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["GetUpdate"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }
