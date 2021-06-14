package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func DhpCompositeKey(stub shim.ChaincodeStubInterface, patientDid, method string) (string, error) {
	switch Config.DBType {
	case CouchDB:
		return stub.CreateCompositeKey("patient~method", []string{patientDid, method})
	case LevelDB:
		fallthrough
	default:
		return patientDid + method, nil
	}
}

func FetchDhpsFromLedger(stub shim.ChaincodeStubInterface, patientDid, method string) ([]Dhp, error) {
	switch {
	case Config.DBType == CouchDB && (method == "*" || strings.EqualFold(method, "any")):
		var dhps []Dhp
		st, err := stub.GetStateByPartialCompositeKey("patient~method", []string{patientDid})
		if err != nil {
			return nil, fmt.Errorf("Error retrieving DHPs from ledger using partial composite key: %w", err)
		}
		defer st.Close()
		for st.HasNext() {
			kv, err := st.Next()
			if err != nil {
				return nil, fmt.Errorf("Error accessing DHP result set: %w", err)
			}
			var dhp Dhp
			if err := json.Unmarshal(kv.Value, &dhp); err != nil {
				return nil, fmt.Errorf("Error unmarshaling DHP: %w", err)
			}
			dhps = append(dhps, dhp)
		}
		return dhps, nil
	case Config.DBType == CouchDB && method != "*":
		fallthrough
	case Config.DBType == LevelDB:
		fallthrough
	default:
		dhpCompKey, err := DhpCompositeKey(stub, patientDid, method)
		if err != nil {
			return nil, fmt.Errorf("Error determining composite key for DHP: %w", err)
		}
		dhpB, err := stub.GetState(dhpCompKey)
		if err != nil {
			return nil, fmt.Errorf("Error retrieving DHP from ledger: %w", err)
		}
		if len(dhpB) == 0 {
			return nil, fmt.Errorf("DHP for (%s, %s) not found / empty: %w", patientDid, method, err)
		}
		dhp := Dhp{}
		if err := json.Unmarshal(dhpB, &dhp); err != nil {
			return nil, fmt.Errorf("Error unmarshaling DHP: %w", err)
		}
		return []Dhp{dhp}, nil
	}
}
