package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"api-gin/database"
	"api-gin/models"
)

func All(c *gin.Context) {
	var customer []models.Customer
	database.DB.Find(&customer)
	c.JSON(200, customer)
}

func Congrats(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza?",
	})
}

func Create(c *gin.Context) {

	var customer models.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidateCustomer(&customer); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&customer)
	c.JSON(http.StatusOK, customer)
}

func FindById(c *gin.Context) {
	var customer models.Customer
	id := c.Params.ByName("id")
	database.DB.First(&customer, id)

	if customer.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "customer não encontrado"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func Delete(c *gin.Context) {
	var customer models.Customer
	id := c.Params.ByName("id")
	database.DB.Delete(&customer, id)
	c.JSON(http.StatusOK, gin.H{"data": "customer deletado com sucesso"})
}

func Edit(c *gin.Context) {
	var customer models.Customer
	id := c.Params.ByName("id")
	database.DB.First(&customer, id)

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidateCustomer(&customer); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&customer).UpdateColumns(customer)
	c.JSON(http.StatusOK, customer)
}

func FindByCpf(c *gin.Context) {
	var customer models.Customer
	cpf := c.Param("cpf")
	database.DB.Where(&models.Customer{CPF: cpf}).First(&customer)

	if customer.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "customer não encontrado"})
		return
	}

	c.JSON(http.StatusOK, customer)
}