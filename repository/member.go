package repository

import (
	"context"

	"test/driver"
	"test/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindAllMember() ([]models.Member, error) {
	var members []models.Member
	db, client, err := driver.ConnectMongo("member")
	if err != nil {
		return members, err
	}
	cur, err := db.Find(context.TODO(), bson.M{})
	if err != nil {
		return members, err
	}
	for cur.Next(context.Background()) {
		var member models.Member
		cur.Decode(&member)
		members = append(members, member)
	}
	err = client.Disconnect(context.Background())
	if err != nil {
		return members, err
	}
	return members, nil
}

func FindMember(id primitive.ObjectID) (models.Member, error) {
	var member models.Member
	db, client, err := driver.ConnectMongo("member")
	if err != nil {
		return member, err
	}
	err = db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&member)
	if err != nil {
		return member, err
	}
	err = client.Disconnect(context.Background())

	if err != nil {
		return member, err
	}
	return member, nil
}

func FindMemberUsernamePassword(auth models.AuthenMember) (models.Member, error) {
	var Member models.Member
	db, client, err := driver.ConnectMongo("member")
	if err != nil {
		return Member, err
	}
	err = db.FindOne(context.TODO(), bson.M{"username": auth.Username, "password": auth.Password}).Decode(&Member)
	if err != nil {
		return Member, err
	}
	err = client.Disconnect(context.Background())

	if err != nil {
		return Member, err
	}
	return Member, nil
}

func FindMemberLastID() (models.Member, error) {
	var Member models.Member
	db, client, err := driver.ConnectMongo("member")
	if err != nil {
		return Member, err
	}
	opts := options.FindOne().SetSort(bson.D{{"_id", -1}})
	err = db.FindOne(context.TODO(), bson.M{}, opts).Decode(&Member)
	if err != nil {
		return Member, err
	}
	err = client.Disconnect(context.Background())

	if err != nil {
		return Member, err
	}
	return Member, nil
}

func InsertMember(member models.Member) error {
	db, client, err := driver.ConnectMongo("member")
	if err != nil {
		return err
	}
	_, err = db.InsertOne(context.Background(), member)
	if err != nil {
		return err
	}
	err = client.Disconnect(context.Background())

	if err != nil {
		return err
	}
	return nil
}

func UpdateMember(Member models.PutMember, id primitive.ObjectID) error {
	db, client, err := driver.ConnectMongo("member")
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", Member}}
	_, err = db.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		return err
	}
	err = client.Disconnect(context.Background())

	if err != nil {
		return err
	}
	return nil
}

func DeleteMember(id primitive.ObjectID) error {
	db, client, err := driver.ConnectMongo("member")
	if err != nil {
		return err
	}

	_, err = db.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	err = client.Disconnect(context.Background())

	if err != nil {
		return err
	}
	return err
}
