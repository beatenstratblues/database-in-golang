package main

import (
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DataTransforming interface {
	Marshall(data []byte) (*[]byte, error)
	Unmarshall(data []byte) ([]byte, error)
}

type DataTransformer struct{}

func (DataTransformer) Marshall(rawData []byte) (*bson.Raw, error) {
	var bsonBytes bson.Raw
	err := bson.UnmarshalExtJSON(rawData, false, &bsonBytes)
	if err != nil {
		return nil, err
	}
	return &bsonBytes, nil
}

func (DataTransformer) Unmarshall(rawData []byte) {
	fmt.Println("This function implements unmarshalling")
}

func NewDataTransformer() *DataTransformer {
	return &DataTransformer{}
}

func main() {
	data, _ := os.ReadFile("test.json")
	test := NewDataTransformer()
	mData, err := test.Marshall(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mData)
}
