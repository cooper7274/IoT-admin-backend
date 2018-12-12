package handler

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"IoT-admin-backend/models"

	"github.com/gin-gonic/gin"
)

// List all mappings
func ListMapping(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var mappings []models.Mapping
	query := bson.M{
		"productId": bson.ObjectIdHex(c.Param("_id")),
	}
	err := db.C(models.CollectionMapping).Find(query).All(&mappings)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, mappings)
}

// Get a mapping
func GetMapping(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var mapping models.Mapping

	err := db.C(models.CollectionMapping).
		FindId(bson.ObjectIdHex(c.Param("_id"))).
		One(&mapping)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, mapping)
}

// Create a mapping
func CreateMapping(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var mapping models.Mapping
	err := c.BindJSON(&mapping)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionProduct).Insert(mapping)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, mapping)
}

// Delete mapping
func DeleteMapping(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
	err := db.C(models.CollectionMapping).Remove(query)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, nil)
}

// Update mapping
func UpdateMapping(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var mapping models.Mapping
	err := c.BindJSON(&mapping)
	if err != nil {
		c.Error(err)

		return
	}

	// 查找原来的文档
	query := bson.M{
		"_id": bson.ObjectIdHex(c.Param("_id")),
	}

	// 更新
	err = db.C(models.CollectionProduct).Update(query, mapping)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, mapping)
}

