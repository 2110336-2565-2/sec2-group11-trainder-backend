package models

import (
	"io"
	"mime/multipart"

	// "os"
	"trainder-api/configs"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UploadFileToMongo(file multipart.File, handler *multipart.FileHeader) (primitive.ObjectID, error) {

	// Create a new GridFS upload stream
	uploadStream, err := configs.Bucket.OpenUploadStream(
		handler.Filename,
	)
	if err != nil {
		return primitive.NilObjectID, err
	}
	defer uploadStream.Close()

	// Copy the file data into the GridFS upload stream
	_, err = io.Copy(uploadStream, file)
	if err != nil {
		return primitive.NilObjectID, err
	}

	// Get the ID of the uploaded file
	fileID := uploadStream.FileID.(primitive.ObjectID)

	return fileID, nil
}

func retrieveFileFromMongo(fileID primitive.ObjectID, fileID_str string) (io.ReadCloser, error) {
	fileID, err := primitive.ObjectIDFromHex(fileID_str)
	if err != nil {

		return nil, err
	}

	// Open a GridFS download stream for the file
	downloadStream, err := configs.Bucket.OpenDownloadStream(fileID)
	if err != nil {
		return nil, err
	}

	return downloadStream, nil
}

func createHeader(filename string) (*multipart.FileHeader, error) {
	handler := &multipart.FileHeader{
		Filename: filename,
		Size:     1000,
	}

	return handler, nil

}

// func Upload(filename string, file multipart.File) (primitive.ObjectID, error) {
// 	// file, err := openFile(filename)
// 	// if err != nil {
// 	// 	return primitive.NilObjectID, err
// 	// }
// 	// handler, err := createHeader(filename)
// 	// if err != nil {
// 	// 	return primitive.NilObjectID, err
// 	// }

// 	// Create a new GridFS upload stream
// 	uploadStream, err := configs.Bucket.OpenUploadStream(
// 		filename,
// 	)
// 	if err != nil {
// 		return primitive.NilObjectID, err
// 	}
// 	defer uploadStream.Close()

// 	// Copy the file data into the GridFS upload stream
// 	_, err = io.Copy(uploadStream, file)
// 	if err != nil {
// 		return primitive.NilObjectID, err
// 	}

// 	// Get the ID of the uploaded file
// 	fileID := uploadStream.FileID.(primitive.ObjectID)

// 	return fileID, nil

// }

func Upload(filename string, file multipart.File) (primitive.ObjectID, error) {

	uploadStream, err := configs.Bucket.OpenUploadStream(
		filename,
	)
	if err != nil {
		return primitive.NilObjectID, err
	}
	defer uploadStream.Close()

	// Copy the file data into the GridFS upload stream
	_, err = io.Copy(uploadStream, file)
	if err != nil {
		return primitive.NilObjectID, err
	}

	// Get the ID of the uploaded file
	fileID := uploadStream.FileID.(primitive.ObjectID)

	return fileID, nil

}
