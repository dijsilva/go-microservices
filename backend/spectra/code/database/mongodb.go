package database

import (
	"context"
	"fmt"
	"log"
	"spectra/commom"
	appErrors "spectra/errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type SpectraAbsorbanceColumn struct {
	Pos   int     `bson:"pos"`
	Value float64 `bson:"value"`
}
type SpectraFileRow struct {
	Row      int                       `bson:"row"`
	IsHeader bool                      `bson:"isHeader"`
	Values   []SpectraAbsorbanceColumn `bson:"values"`
}

type SpectraDTO struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	SampleName    string             `bson:"sample_name" binding:"required"`
	UserOwner     string             `bson:"user_owner"`
	NSamples      int                `bson:"n_samples" binding:"required"`
	EquipmentUsed string             `bson:"equipment_used" binding:"required"`
	Rows          []SpectraFileRow   `bson:"data"`
}

type SpectrasResponse struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	SampleName    string             `bson:"sample_name" binding:"required"`
	UserOwner     string             `bson:"user_owner"`
	NSamples      int                `bson:"n_samples" binding:"required"`
	EquipmentUsed string             `bson:"equipment_used" binding:"required"`
}

type MongoDb struct {
	Client *mongo.Client
}

func InitDatabase() (IDatabase, error) {
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/?maxPoolSize=%s&w=majority",
		commom.Envs.MongoDbUser,
		commom.Envs.MongoDbPass,
		commom.Envs.MongoDbHost,
		commom.Envs.MongoDbPort,
		commom.Envs.MaxPoolSize,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		return nil, err
	}
	return &MongoDb{Client: client}, nil
}

func (m *MongoDb) DisconnectDatabse() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := m.Client.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func (m *MongoDb) Create(input SpectraDTO) (string, appErrors.ErrorResponse) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Println("Getting spectra_request collection")
	collection := m.Client.Database(commom.Envs.MongoDbDatabaseName).Collection("spectra_request")
	input.ID = primitive.NewObjectID()
	dataInserted, err := collection.InsertOne(ctx, input)
	if err != nil {
		return "", appErrors.InternalServerError(fmt.Sprintf("Error to store data - %s", err.Error()))
	}
	hexId := dataInserted.InsertedID.(primitive.ObjectID).Hex() // @TODO: estudar mais
	return hexId, appErrors.ErrorResponse{}
}

func (m *MongoDb) ListByOwner(usernameOwner string) ([]SpectrasResponse, appErrors.ErrorResponse) {
	var results []SpectrasResponse
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Println("Getting spectra_request collection for list by owner")
	filter := bson.D{{"user_owner", usernameOwner}}
	collection := m.Client.Database(commom.Envs.MongoDbDatabaseName).Collection("spectra_request")
	cursor, err := collection.Find(ctx, filter)
	defer func() {
		cursorErr := cursor.Close(ctx)
		if cursorErr != nil {
			panic(cursorErr)
		}
	}()
	if err != nil {
		return results, appErrors.InternalServerError(fmt.Sprintf("Error to store data - %s", err.Error()))
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results, appErrors.ErrorResponse{}
}

func (m *MongoDb) GetById(id string) (SpectraDTO, appErrors.ErrorResponse) {
	var result SpectraDTO
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Println("Getting spectra_request collection for get by id")

	objectId, errParseId := primitive.ObjectIDFromHex(id)
	if errParseId != nil {
		return result, appErrors.BadRequest("invalid id")
	}

	fmt.Println(objectId)
	collection := m.Client.Database(commom.Envs.MongoDbDatabaseName).Collection("spectra_request")
	err := collection.FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&result)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return result, appErrors.NotFound("not found")
		}
		return result, appErrors.InternalServerError(fmt.Sprintf("Error to store data - %s", err.Error()))
	}

	return result, appErrors.ErrorResponse{}
}
