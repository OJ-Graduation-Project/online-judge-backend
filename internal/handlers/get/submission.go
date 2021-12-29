package get

import (
	"fmt"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"go.mongodb.org/mongo-driver/bson"
)

type Submission struct {
	ID         int
	Accepted   bool
	Language   string
	Date       string
	FailedCase struct {
		ID          int
		Reason      string
		User_output string
	}
}

func InsertSubmission(sub Submission, database string, col string, db db.DbConnection) error {
	collection := db.Conn.Database(database).Collection(col)

	bsonBytes, _ := bson.Marshal(sub)
	result, err := collection.InsertOne(db.Ctx, bsonBytes)
	if err != nil {
		fmt.Println("Error in InsertOne()")
		fmt.Println(err)
		return err
	}
	fmt.Println("Inserted Successfully", result)
	return nil
}
func RetrieveSubmission(id int, database string, col string, db db.DbConnection) (Submission, error) {

	collection := db.Conn.Database(database).Collection(col)
	var sub bson.D
	err := collection.FindOne(db.Ctx, bson.M{"id": id}).Decode(&sub)
	if err != nil {
		fmt.Println(err)
		return Submission{}, err
	}
	var ret Submission
	bsonBytes, err := bson.Marshal(sub)
	if err != nil {
		fmt.Println(err)
	}
	bson.Unmarshal(bsonBytes, &ret)
	return ret, nil
}
