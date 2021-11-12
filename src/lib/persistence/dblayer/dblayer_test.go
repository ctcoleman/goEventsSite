package dblayer

import (
	"testing"
)

func TestNewPersistenceLayerDYNAMODB(t *testing.T) {
	if ans, err := NewPersistenceLayer(DYNAMODB, "test"); ans != nil || err != nil {
		t.Errorf("NewPersistenceLayer(DYNAMODB,\"test\") = %d, %d; want nil, nil", ans, err)
	}
}

func TestNewPersistenceLayerDOCUMENTDB(t *testing.T) {
	if ans, err := NewPersistenceLayer(DOCUMENTDB, "test"); ans != nil || err != nil {
		t.Errorf("NewPersistenceLayer(DOCUMENTDB,\"test\") = %d, %d; want nil, nil", ans, err)
	}
}

func TestNewPersistenceLayerMONGODB(t *testing.T) {
	if ans, err := NewPersistenceLayer(MONGODB, "localhost:27017"); ans == nil || err != nil {
		t.Errorf("NewPersistenceLayer(MONGODB, \"localhost:27017\") = %d, %d; want mongolayer, nil", ans, err)
	}
}
