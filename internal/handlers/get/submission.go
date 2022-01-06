package get

import (
	"fmt"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"github.com/OJ-Graduation-Project/online-judge-backend/pkg/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertSubmission(sub entities.Submission, database string, col string, db db.DbConnection) error {
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
func RetrieveSubmission(id int, database string, col string, db db.DbConnection) (entities.Submission, error) {

	collection := db.Conn.Database(database).Collection(col)
	var sub bson.D
	err := collection.FindOne(db.Ctx, bson.M{"id": id}).Decode(&sub)
	if err != nil {
		fmt.Println(err)
		return entities.Submission{}, err
	}
	var ret entities.Submission
	bsonBytes, err := bson.Marshal(sub)
	if err != nil {
		fmt.Println(err)
	}
	bson.Unmarshal(bsonBytes, &ret)
	return ret, nil
}
