package contract

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Contract struct {
	contractapi.Contract
}

// GetWorldState takes a key and gets the world state and returns values
func GetWorldState(ctx contractapi.TransactionContextInterface, key string, value interface{}) error {
	result, err := ctx.GetStub().GetState(key)
	if err != nil {
		return fmt.Errorf("failed to get world state: %v", err)
	}

	return json.Unmarshal(result, value)
}

// SetWorldState takes a key and value and sets on the world state
func SetWorldState(сtx contractapi.TransactionContextInterface, key string, value interface{}) error {
	b, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("error marshalling: %v", err)
	}

	return сtx.GetStub().PutState(key, b)
}

// CreateBook creating new book and set in world state
func (c *Contract) CreateBook(ctx contractapi.TransactionContextInterface, title, author, pubkisher string) error {
	bookDetails := Book{
		ID:        GenKeyID("book@"),
		Title:     title,
		Author:    author,
		Publisher: pubkisher,
	}

	if err := SetWorldState(ctx, bookDetails.ID, bookDetails); err != nil {
		return fmt.Errorf("error creating book: %v", err)
	}

	return nil
}

// DeleteBook delete the book from world state
func (c *Contract) DeleteBook(ctx contractapi.TransactionContextInterface, id string) error {
	if err := ctx.GetStub().DelState(id); err != nil {
		return fmt.Errorf("failed to delete in world state: %v", err)
	}

	return nil
}

// GetBookById returns book by id from world state
func (c *Contract) GetBookById(ctx contractapi.TransactionContextInterface, id string) (*Book, error) {
	bookDetails := new(Book)
	err := GetWorldState(ctx, id, bookDetails)
	if err != nil {
		return nil, err
	}

	return bookDetails, nil
}

// GetBooks returns books by all id from world state
func (c *Contract) GetBooks(ctx contractapi.TransactionContextInterface) ([]*Book, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to get world state: %v", err)
	}
	defer resultsIterator.Close()

	sliceBooks := []*Book{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		bookDetails := new(Book)
		err = json.Unmarshal(queryResponse.Value, bookDetails)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal book: %v", err)
		}
		sliceBooks = append(sliceBooks, bookDetails)
	}

	return sliceBooks, nil
}
