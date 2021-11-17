package handler

// import (
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/zenklot/backend-zenknote/database"
// 	"github.com/zenklot/backend-zenknote/model"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// func GetEmployee(c *fiber.Ctx) error {
// 	query := bson.D{{}}
// 	cursor, err := database.MI.DB.Collection("employees").Find(c.Context(), query)
// 	if err != nil {
// 		return c.Status(500).SendString(err.Error())
// 	}

// 	var employees []model.Employee = make([]model.Employee, 0)

// 	// iterate the cursor and decode each item into an Employee
// 	if err := cursor.All(c.Context(), &employees); err != nil {
// 		return c.Status(500).SendString(err.Error())

// 	}
// 	// return employees list in JSON format
// 	return c.JSON(employees)
// }

// func PostEmployee(c *fiber.Ctx) error {
// 	collection := database.MI.DB.Collection("employees")

// 	employee := new(model.Employee)
// 	err := c.BodyParser(employee)
// 	if err != nil {
// 		return c.Status(400).SendString(err.Error())
// 	}

// 	employee.ID = ""

// 	insertResult, err := collection.InsertOne(c.Context(), employee)
// 	if err != nil {
// 		return c.Status(500).SendString(err.Error())
// 	}

// 	filter := bson.D{{Key: "_id", Value: insertResult.InsertedID}}
// 	createRecord := collection.FindOne(c.Context(), filter)

// 	createdEmployee := &model.Employee{}
// 	createRecord.Decode(createdEmployee)
// 	return c.Status(201).JSON(createdEmployee)
// }

// func PutEmployee(c *fiber.Ctx) error {
// 	idParam := c.Params("id")
// 	employeeID, err := primitive.ObjectIDFromHex(idParam)

// 	// the provided ID might be invalid ObjectID
// 	if err != nil {
// 		return c.SendStatus(400)
// 	}

// 	employee := new(model.Employee)
// 	// Parse body into struct
// 	if err := c.BodyParser(employee); err != nil {
// 		return c.Status(400).SendString(err.Error())
// 	}

// 	// Find the employee and update its data
// 	query := bson.D{{Key: "_id", Value: employeeID}}
// 	update := bson.D{
// 		{Key: "$set",
// 			Value: bson.D{
// 				{Key: "name", Value: employee.Name},
// 				{Key: "age", Value: employee.Age},
// 				{Key: "salary", Value: employee.Salary},
// 			},
// 		},
// 	}
// 	err = database.MI.DB.Collection("employees").FindOneAndUpdate(c.Context(), query, update).Err()

// 	if err != nil {
// 		// ErrNoDocuments means that the filter did not match any documents in the collection
// 		if err == mongo.ErrNoDocuments {
// 			return c.SendStatus(404)
// 		}
// 		return c.SendStatus(500)
// 	}

// 	// return the updated employee
// 	employee.ID = idParam
// 	return c.Status(200).JSON(employee)
// }

// func DeleteEmployee(c *fiber.Ctx) error {
// 	employeeID, err := primitive.ObjectIDFromHex(
// 		c.Params("id"),
// 	)

// 	// the provided ID might be invalid ObjectID
// 	if err != nil {
// 		return c.SendStatus(400)
// 	}

// 	// find and delete the employee with the given ID
// 	query := bson.D{{Key: "_id", Value: employeeID}}
// 	result, err := database.MI.DB.Collection("employees").DeleteOne(c.Context(), &query)

// 	if err != nil {
// 		return c.SendStatus(500)
// 	}

// 	// the employee might not exist
// 	if result.DeletedCount < 1 {
// 		return c.SendStatus(404)
// 	}

// 	// the record was deleted
// 	return c.SendStatus(204)
// }
