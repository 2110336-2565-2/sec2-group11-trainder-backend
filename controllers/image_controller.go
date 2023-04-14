package controllers

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"trainder-api/models"
	"trainder-api/responses"
	"trainder-api/utils/tokens"

	"github.com/gin-gonic/gin"
)

type Imagedata struct {
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

// func UploadProfile() gin.HandlerFunc {
// 	fmt.Println("UploadProfile")
// 	return func(c *gin.Context) {
// 		file, err := c.FormFile("image")
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
// 			return
// 		}

// 		f, err := file.Open()
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
// 			return
// 		}
// 		defer f.Close()
// 		// Read the contents of the file into memory
// 		// data, err := ioutil.ReadAll(f)
// 		// if err != nil {
// 		// 	c.JSON(http.StatusInternalServerError, gin.H{"error3": err.Error()})
// 		// 	return
// 		// }

// 		fileID, err := models.Upload(file.Filename, f)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		log.Printf("File uploaded with ID: %v\n", fileID)

// 		// Check if the file is a JPG image
// 		// _, format, err := image.DecodeConfig(bytes.NewReader(data))
// 		// fmt.Println("format", format)
// 		// if err != nil || format != "jpeg" {
// 		// 	c.JSON(http.StatusBadRequest, gin.H{"error4": "Uploaded file is not a JPG image " + format})
// 		// 	return
// 		// }

// 		// Save the file to disk
// 		// err = ioutil.WriteFile("image.jpg", data, 0644)
// 		// if err != nil {
// 		// 	c.JSON(http.StatusInternalServerError, gin.H{"error5": err.Error()})
// 		// 	return
// 		// }

// 		// Return a success message
// 		c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})
// 		// body, _ := ioutil.ReadAll(c.Request.Body)
// 		// println(string(body))
// 		// var input Imagedata
// 		// if err := c.ShouldBindJSON(&input); err != nil {
// 		// 	c.JSON(http.StatusBadRequest, responses.ProfileResponse{
// 		// 		Status:  http.StatusBadRequest,
// 		// 		Message: err.Error(),
// 		// 	})
// 		// 	return
// 		// }

// 		// fmt.Println(input)

//		}
//	}
func UploadProfile() gin.HandlerFunc {
	fmt.Println("UploadProfile")
	return func(c *gin.Context) {
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
			return
		}

		if file.Size <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Uploaded file is empty"})
			return
		}

		f, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
			return
		}
		defer f.Close()

		fileID, err := models.Upload(file.Filename, f)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("File uploaded with ID: %v\n", fileID)

		c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})

		// update profile with avatar URI = fileID
		username, err := tokens.ExtractTokenUsername(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ProfileResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		_, err = models.UpdateAvatarUrl(username, fileID.Hex())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
			return
		}

	}
}

func GetPicture() gin.HandlerFunc {
	return func(c *gin.Context) {
		// id := c.Query("ID")
		username := c.Query("username")
		imageId, err := models.GetAvatarUrl(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
			return
		}
		downloadStream, filename, err := models.RetrieveFileFromMongo(imageId)
		if err != nil {
			log.Fatal(err)
		}
		defer downloadStream.Close()
		out, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		_, err = io.Copy(out, downloadStream)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("File saved to test_output.jpg")

	}
}
