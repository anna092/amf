/*
 * Namf_Communication
 *
 * AMF Communication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

 package communication

 import (
	"net/http"
 	"context"
	"encoding/json"
	"io/ioutil"
	
	"github.com/gin-gonic/gin"
 	"github.com/free5gc/amf/internal/logger"
	"github.com/free5gc/openapi/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
 )
 
 func HTTPNonUeN2InfoSubscribe(c *gin.Context) {
	logger.CommLog.Infof("Handle Non Ue N2 Info Subscribe is being implemented.")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.CommLog.Errorf("Error reading request body: %v", err)
    	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
    	return
	}
	message := models.NonUeN2InfoSubscriptionCreateData{}
	err = json.Unmarshal(body, &message)
	if err != nil {
    	logger.CommLog.Errorf("Error unmarshalling request body: %v", err)
    	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
    	return
	}

	err = insertToSubscriptionDatabase(message)
    if err != nil {
        logger.CommLog.Errorf("Error inserting to database: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database insertion failed"})
        return
    }
    c.JSON(http.StatusOK, gin.H{})
 }
 
 func insertToSubscriptionDatabase(message models.NonUeN2InfoSubscriptionCreateData) error {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        return err
    }
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        return err
    }
    collection := client.Database("local").Collection("AMFNonUeN2MessageSubscriptions")
    _, err = collection.InsertOne(context.TODO(), message)
    if err != nil {
        return err
    }
    return nil
}

 
 /*
 // NonUeN2InfoSubscribe - Namf_Communication Non UE N2 Info Subscribe service Operation
 func HTTPNonUeN2InfoSubscribe(c *gin.Context) {
	 logger.CommLog.Warnf("Handle Non Ue N2 Info Subscribe is not implemented.")
	 c.JSON(http.StatusOK, gin.H{})
 }
 */
 