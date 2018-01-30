package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/Azure/azure-pipeline-go/pipeline"
	"github.com/Azure/azure-storage-blob-go/2016-05-31/azblob"
)

/// <summary>
/// Azure Storage Quickstart Sample - Demonstrate how to upload, list, download, and delete blobs.
///
///
/// Documentation References:
/// - What is a Storage Account - https://docs.microsoft.com/azure/storage/common/storage-create-storage-account
/// - Blob Service Concepts - https://docs.microsoft.com/rest/api/storageservices/Blob-Service-Concepts
/// - Blob Service Go SDK API - https://godoc.org/github.com/Azure/azure-storage-blob-go
/// - Blob Service REST API - https://docs.microsoft.com/rest/api/storageservices/Blob-Service-REST-API
/// - Scalability and performance targets - https://docs.microsoft.com/azure/storage/common/storage-scalability-targets
/// - Azure Storage Performance and Scalability checklist https://docs.microsoft.com/azure/storage/common/storage-performance-checklist
/// - Storage Emulator - https://docs.microsoft.com/azure/storage/common/storage-use-emulator
/// </summary>

func randomString() (random string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	random = strconv.Itoa(r.Int())
	return
}

func handleErrors(err error) {
	if err != nil {
		if serr, ok := err.(azblob.StorageError); ok { // This error is a Service-specific
			switch serr.ServiceCode() { // Compare serviceCode to ServiceCodeXxx constants
			case azblob.ServiceCodeContainerAlreadyExists:
				fmt.Println("Received 409. Container already exists")
				break
			default:
				// Handle other errors ...
				log.Fatal(err)
			}
		}
	}
}

func main() {

	fmt.Printf("Azure Blob storage quick start sample\n")

	// From the Azure portal, get your Storage account's name and key and set environment variables.
	accountName, accountKey := os.Getenv("AZURE_STORAGE_ACCOUNT"), os.Getenv("AZURE_STORAGE_ACCESS_KEY")
	if len(accountName) == 0 || len(accountKey) == 0 {
		log.Fatal("AZURE_STORAGE_ACCOUNT and AZURE_STORAGE_ACCESS_KEY environment variables are not set")
	}

	// Create a request pipeline using your Storage account's name and account key.
	credential := azblob.NewSharedKeyCredential(accountName, accountKey)
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	// Create a random string for the quick start container
	containerName := fmt.Sprintf("%s%s", "quickstart-", randomString())

	// From the Azure portal, get your Storage account blob service URL endpoint.
	URL, _ := url.Parse(
		fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName))

	// Create a ContainerURL object that wraps the container URL and a request
	// pipeline to make requests.
	containerURL := azblob.NewContainerURL(*URL, p)

	// Create the container
	fmt.Printf("Creating a container named %s\n", containerName)
	ctx := context.Background() // This example uses a never-expiring context
	_, err := containerURL.Create(ctx, azblob.Metadata{}, azblob.PublicAccessNone)
	handleErrors(err)

	// Create a file to test the upload and download.
	fmt.Printf("Creating a dummy file to test the upload and download\n")
	data := []byte("hello world\nthis is a blob\n")
	fileName := randomString()
	err = ioutil.WriteFile(fileName, data, 0700)
	handleErrors(err)

	// Here's how to upload a blob.
	blobURL := containerURL.NewBlockBlobURL(fileName)
	file, err := os.Open(fileName)
	handleErrors(err)

	// Use the PutBlob API to upload the file
	// Note that PutBlob can upload up to 256MB data in one shot. Details: https://docs.microsoft.com/en-us/rest/api/storageservices/put-blob
	fmt.Printf("Uploading the file with blob name: %s\n", fileName)
	_, err = blobURL.PutBlob(ctx, file, azblob.BlobHTTPHeaders{}, azblob.Metadata{}, azblob.BlobAccessConditions{})
	handleErrors(err)

	// Alternatively you can use the high level API UploadFileToBlockBlob function to upload blocks in parallel.
	// This function calls PutBlock/PutBlockLlist for files larger 256 MBs, and calls PutBlob for any file smaller
	// Note this will overwrite the file uploaded by PutBlob in the previous line
	_, err = azblob.UploadFileToBlockBlob(ctx, file, blobURL, azblob.UploadToBlockBlobOptions{
		BlockSize:   4 * 1024 * 1024,
		Parallelism: 16})
	handleErrors(err)

	// List the blobs in the container
	marker := azblob.Marker{}
	for marker.NotDone() {
		// Get a result segment starting with the blob indicated by the current Marker.
		listBlob, err := containerURL.ListBlobs(ctx, marker, azblob.ListBlobsOptions{})
		handleErrors(err)

		// ListBlobs returns the start of the next segment; you MUST use this to get
		// the next segment (after processing the current result segment).
		marker = listBlob.NextMarker

		// Process the blobs returned in this result segment (if the segment is empty, the loop body won't execute)
		for _, blobInfo := range listBlob.Blobs.Blob {
			fmt.Print("Blob name: " + blobInfo.Name + "\n")
		}
	}

	// Here's how to download the blob
	get, err := blobURL.GetBlob(ctx, azblob.BlobRange{}, azblob.BlobAccessConditions{}, false)
	handleErrors(err)

	// Wrap the response body in a ResponseBodyProgress and pass a callback function
	// for progress reporting.
	responseBody := pipeline.NewResponseBodyProgress(get.Body(),
		func(bytesTransferred int64) {
			fmt.Printf("Read %d of %d bytes.\n", bytesTransferred, get.ContentLength())
		})
	downloadedData := &bytes.Buffer{}
	downloadedData.ReadFrom(responseBody)
	// The downloaded blob data is in downloadData's buffer. :Let's print it
	fmt.Printf("Downloaded the blob: " + downloadedData.String())

	// Cleaning up the quick start by deleting the container and the file created locally
	fmt.Printf("Press a key to delete the sample files, example container, and exit the application.\n")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	fmt.Printf("Cleaning up.\n")
	containerURL.Delete(ctx, azblob.ContainerAccessConditions{})
	file.Close()
	os.Remove(fileName)

}
