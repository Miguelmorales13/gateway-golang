package models

import db "../conexion"
import "../entities"


func FindAllApis() (apis []entities.Api)  {
	my := db.DbConn()
	my.Where(map[string]interface{}{"active": true}).Find(&apis)
	defer my.Close()
	return
}

func FindBySecrets(apiKey string) (api entities.Api) {
	my := db.DbConn()
	my.Where(map[string]interface{}{"api_key": apiKey,"active":true}).First(&api)
	defer my.Close()

	return
}
/**
  func FindById(id primitive.ObjectID) (api entities.Api, err error) {
  	var collection = mongodb.GetMongo().GetCollection("apis")
  	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
  	err = collection.FindOne(ctx, bson.M{"id":id}).Decode(&api)
  	return
  }



  func Create(api entities.Api) (apiCreated entities.Api , err error) {
  	var collection = mongodb.GetMongo().GetCollection("apis")
  	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
  	result,err := collection.InsertOne(ctx, api)
  	if err !=nil { return }
  	apiCreated,err=FindById(result.InsertedID.(primitive.ObjectID))
  	return
  }

  func Update(id primitive.ObjectID, api entities.Api) (apiUpdated entities.Api, err error) {
  	var collection = mongodb.GetMongo().GetCollection("apis")
  	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
  	_ ,err= collection.UpdateOne(ctx, bson.M{"_id":id}, bson.M{ "$set": api })
  	if err!=nil { return }
  	apiUpdated,err=FindById(id)
  	return
  }*/
