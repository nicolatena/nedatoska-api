package repository


import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    . "go-base-cleancode/models"
)



type InDB struct {
    DB *gorm.DB
}

func (idb *InDB) Fetch(c *gin.Context) {

    var arr_data []Product
    var arr_meta Meta
    var response ResponseProduct

    idb.DB.Find(&arr_data)

    arr_meta = Meta{Status: true, Code: 200, Message: "Success"}

    response.Meta = arr_meta
    response.Data = arr_data
 
    c.JSON(http.StatusOK, response)
}


func (idb *InDB) Store(c *gin.Context) {
    
    var arr_data []Product
    var arr_meta Meta
    var response ResponseProduct
    sc_data := Product{}


    err := c.BindJSON(&sc_data)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  http.StatusBadRequest,
            "message": "can't bind struct",
        })
        return
    }

    sc_data.DeletedAt = nil
    
    idb.DB.Create(&sc_data)

    idb.DB.Find(&arr_data)

    arr_meta = Meta{Status: true, Code: 200, Message: "Success Inserted"}

    response.Meta = arr_meta
    response.Data = arr_data
 
    c.JSON(http.StatusOK, response)
}


func (idb *InDB) Update(c *gin.Context) {
    
    var arr_data []Product
    var arr_meta Meta
    var response ResponseProduct
    sc_data := Product{}

    
    id := c.Param("id")
    idb.DB.Where("id = ?", id).First(&sc_data)
    
    c.BindJSON(&sc_data)
    sc_data.DeletedAt = nil

    idb.DB.Save(sc_data)

    idb.DB.Find(&arr_data)

    arr_meta = Meta{Status: true, Code: 200, Message: "Success Updated"}

    response.Meta = arr_meta
    response.Data = arr_data
 
    c.JSON(http.StatusOK, response)
}


func (idb *InDB) Delete(c *gin.Context) {
    
    var arr_data []Product
    var arr_meta Meta
    var response ResponseProduct
    sc_data := Product{}

    
    id := c.Param("id")
    idb.DB.Where("id = ?", id).First(&sc_data)
    //idb.DB.Delete(sc_data) // SOFT DELETE
    idb.DB.Unscoped().Delete(sc_data) // DELETE PERMANENT

    idb.DB.Find(&arr_data)

    arr_meta = Meta{Status: true, Code: 200, Message: "Success Deleted"}

    response.Meta = arr_meta
    response.Data = arr_data
 
    c.JSON(http.StatusOK, response)
}