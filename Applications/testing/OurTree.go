package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Record
type SmartContract struct {
	contractapi.Contract
}

type Record struct {
	ID        string `json:"ID"`
	Header    string `json:"header"`
	Data      string `json:"data"`
	Traceback string `json:"traceback"`
}

// InitLedger adds a base set of records to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	records := []Record{
		{ID: "record1", Header: "blue", Data: "5", Traceback: "None"},
	}

	for _, record := range records {
		recordJSON, err := json.Marshal(record)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(record.ID, recordJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateRecord issues a new record to the world state with given details.
func (s *SmartContract) CreateRecord(ctx contractapi.TransactionContextInterface, id string, header string, data string, traceback string) error {
	exists, err := s.RecordExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the record %s already exists", id)
	}

	record := Record{
		ID:        id,
		Header:    header,
		Data:      data,
		Traceback: traceback,
	}
	recordJSON, err := json.Marshal(record)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, recordJSON)
}

// ReadRecord returns the record stored in the world state with given id.
func (s *SmartContract) ReadRecord(ctx contractapi.TransactionContextInterface, id string) (*Record, error) {
	recordJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if recordJSON == nil {
		return nil, fmt.Errorf("the record %s does not exist", id)
	}

	var record Record
	err = json.Unmarshal(recordJSON, &record)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

// // UpdateRecord updates an existing record in the world state with provided parameters.
// func (s *SmartContract) UpdateRecord(ctx contractapi.TransactionContextInterface, id string, color string, size int, owner string, appraisedValue int) error {
// 	exists, err := s.RecordExists(ctx, id)
// 	if err != nil {
// 		return err
// 	}
// 	if !exists {
// 		return fmt.Errorf("the record %s does not exist", id)
// 	}

// 	// overwriting original record with new record
// 	record := Record{
// 		ID:             id,
// 		Color:          color,
// 		Size:           size,
// 		Owner:          owner,
// 		AppraisedValue: appraisedValue,
// 	}
// 	recordJSON, err := json.Marshal(record)
// 	if err != nil {
// 		return err
// 	}

// 	return ctx.GetStub().PutState(id, recordJSON)
// }

// DeleteRecord deletes an given record from the world state.
func (s *SmartContract) DeleteRecord(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.RecordExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the record %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// RecordExists returns true when record with given ID exists in world state
func (s *SmartContract) RecordExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	recordJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return recordJSON != nil, nil
}

// // TransferRecord updates the owner field of record with given id in world state.
// func (s *SmartContract) TransferRecord(ctx contractapi.TransactionContextInterface, id string, newOwner string) error {
// 	record, err := s.ReadRecord(ctx, id)
// 	if err != nil {
// 		return err
// 	}

// 	record.Owner = newOwner
// 	recordJSON, err := json.Marshal(record)
// 	if err != nil {
// 		return err
// 	}

// 	return ctx.GetStub().PutState(id, recordJSON)
// }

// GetAllRecords returns all records found in world state
func (s *SmartContract) GetAllRecords(ctx contractapi.TransactionContextInterface) ([]*Record, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all records in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var records []*Record
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var record Record
		err = json.Unmarshal(queryResponse.Value, &record)
		if err != nil {
			return nil, err
		}
		records = append(records, &record)
	}

	return records, nil
}
