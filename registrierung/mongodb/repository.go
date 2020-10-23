package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"training-fellow.de/registrierung"
)

//NewRepo erzeugt ein neues RegistrierungsRepository für MongoDB
func NewRepo(url, database, collection string) registrierung.RegistrierungsRepository {
	return &mongoDBRepositoy{url, collection, database}
}

type mongoDBRepositoy struct {
	url        string
	collection string
	database   string
}

type mongoCall func(*mongo.Collection) error

//SaveRegistrierung speichert die übergebene Registrierung
func (m *mongoDBRepositoy) SaveRegistrierung(registrierung *registrierung.Registrierung) error {
	fmt.Println("Save")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	return m.executeInClient(ctx, func(collection *mongo.Collection) error {
		registrierung.ID = primitive.NewObjectID().Hex()
		_, err := collection.InsertOne(ctx, registrierung)
		return err
	})

}

//GetUnconfirmedRegistrierungen lieferte eine Liste aller bestätigter Registrierungen
func (m *mongoDBRepositoy) GetUnconfirmedRegistrierungen() ([]*registrierung.Registrierung, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	registrations := make([]*registrierung.Registrierung, 0)
	err := m.executeInClient(ctx, func(collection *mongo.Collection) error {
		cursor, err := collection.Find(ctx, bson.M{"confirmed": false})
		if err != nil {
			return err
		}
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			reg := &registrierung.Registrierung{}
			cursor.Decode(reg)
			registrations = append(registrations, reg)
		}
		return nil
	})

	return registrations, err
}

//ConfirmedRegistrierung bestätigt eine Registrierung
func (m *mongoDBRepositoy) ConfirmedRegistrierung(registrierungsID string) (*registrierung.Registrierung, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	registrierung := &registrierung.Registrierung{}
	err := m.executeInClient(ctx, func(collection *mongo.Collection) error {
		result := collection.FindOneAndUpdate(ctx,
			bson.M{"_id": registrierungsID},
			bson.M{"$set": bson.M{"confirmed": true}})
		return result.Decode(&registrierung)
	})
	return registrierung, err
}

//Template Funktion zum Ausführen von MongoDB Aufrufen
func (m *mongoDBRepositoy) executeInClient(ctx context.Context, callback mongoCall) error {
	client, err := mongo.NewClient(
		options.Client().
			ApplyURI(m.url))
	if err != nil {
		return err
	}

	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database(m.database).Collection(m.collection)
	return callback(collection)

}
