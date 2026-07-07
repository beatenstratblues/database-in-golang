package codec

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type DataTransforming interface {
	Marshall(data []byte) (bson.Raw, error)
	Unmarshall(bsonData bson.Raw) ([]byte, error)
}

type DataTransformer struct{}

func (DataTransformer) Marshall(rawData []byte) (bson.Raw, error) {
	var bsonBytes bson.Raw
	err := bson.UnmarshalExtJSON(rawData, false, &bsonBytes)
	if err != nil {
		return nil, err
	}
	return bsonBytes, nil
}

func (DataTransformer) Unmarshall(bsonData bson.Raw) ([]byte, error) {
	data, err := bson.MarshalExtJSON(bsonData, false, true)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NewDataTransformer() DataTransforming {
	return &DataTransformer{}
}
