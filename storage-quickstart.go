package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

// Azure Storage Quickstart Sample - Demonstrate how to upload, list, download, and delete blobs.
//
// Documentation References:
// - What is a Storage Account - https://docs.microsoft.com/azure/storage/common/storage-create-storage-account
// - Blob Service Concepts - https://docs.microsoft.com/rest/api/storageservices/Blob-Service-Concepts
// - Blob Service Go SDK API - https://godoc.org/github.com/Azure/azure-storage-blob-go
// - Blob Service REST API - https://docs.microsoft.com/rest/api/storageservices/Blob-Service-REST-API
// - Scalability and performance targets - https://docs.microsoft.com/azure/storage/common/storage-scalability-targets
// - Azure Storage Performance and Scalability checklist https://docs.microsoft.com/azure/storage/common/storage-performance-checklist
// - Storage Emulator - https://docs.microsoft.com/azure/storage/common/storage-use-emulator

func randomString() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(r.Int())
}

func main() {
	fmt.Printf("Azure Blob storage quick start sample\n")

	// Replace `<StorageAccountName` with your storage account
	url := "https://storageblobsgo.blob.core.windows.net/" //replace with <StorageAccountName>
	ctx := context.Background()

	// Create a default request pipeline using your storage account name and account key.
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatal("Invalid credentials with error: " + err.Error())
	}

	serviceClient, err := azblob.NewServiceClient(url, credential, nil)
	if err != nil {
		log.Fatal("Invalid credentials with error: " + err.Error())
	}

	// Create the container
	containerName := fmt.Sprintf("quickstart-%s", randomString())
	fmt.Printf("Creating a container named %s\n", containerName)
	containerClient := serviceClient.NewContainerClient(containerName)
	_, err = containerClient.Create(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Creating a dummy file to test the upload and download\n")

	data := []byte("\nhello world this is a blob\n")
	blobName := "blob" + "-" + randomString()

	var blockOptions azblob.HighLevelUploadToBlockBlobOption

	blobClient, err := azblob.NewBlockBlobClient(url+containerName+"/"+blobName, credential, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Upload to data to blob storage
	_, err = blobClient.UploadBufferToBlockBlob(ctx, data, blockOptions)

	if err != nil {
		log.Fatalf("Failure to upload to blob: %+v", err)
	}

	// List the blobs in the container
	fmt.Println("Listing the blobs in the container:")

	pager := containerClient.ListBlobsFlat(nil)

	for pager.NextPage(ctx) {
		resp := pager.PageResponse()

		for _, v := range resp.ContainerListBlobFlatSegmentResult.Segment.BlobItems {
			fmt.Println(*v.Name)
		}
	}

	if err = pager.Err(); err != nil {
		log.Fatalf("Failure to list blobs: %+v", err)
	}

	// Download the blob
	get, err := blobClient.Download(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	downloadedData := &bytes.Buffer{}
	reader := get.Body(azblob.RetryReaderOptions{})
	_, err = downloadedData.ReadFrom(reader)
	if err != nil {
		log.Fatal(err)
	}
	err = reader.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(downloadedData.String())

	// Delete the blob
	fmt.Printf("Deleting the blob " + blobName + "\n")

	_, err = blobClient.Delete(ctx, nil)
	if err != nil {
		log.Fatalf("Failure: %+v", err)
	}

	// Delete the container
	fmt.Printf("Deleting the blob " + containerName + "\n")
	_, err = containerClient.Delete(ctx, nil)

	if err != nil {
		log.Fatalf("Failure: %+v", err)
	}
}
