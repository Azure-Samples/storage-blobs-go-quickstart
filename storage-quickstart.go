package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
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

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	fmt.Printf("Azure Blob storage quick start sample\n")

	// TODO: replace <storage-account-name> with your actual storage account name
	url := "https://<storage-account-name>.blob.core.windows.net/"
	ctx := context.Background()

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	client, err := azblob.NewClient(url, credential, nil)
	handleError(err)

	serviceClient := client.ServiceClient()

	// Create the container
	containerName := "quickstart-sample-container"
	fmt.Printf("Creating a container named %s\n", containerName)
	containerClient := serviceClient.NewContainerClient(containerName)
	_, err = containerClient.Create(ctx, nil)
	handleError(err)

	data := []byte("\nHello, world! This is a blob.\n")
	blobName := "sample-blob"

	blobClient := containerClient.NewBlockBlobClient(blobName)

	// Upload to data to blob storage
	fmt.Printf("Uploading a blob named %s\n", blobName)
	_, err = blobClient.UploadBuffer(ctx, data, &azblob.UploadBufferOptions{})
	handleError(err)

	// List the blobs in the container
	fmt.Println("Listing the blobs in the container:")

	pager := containerClient.NewListBlobsFlatPager(&container.ListBlobsFlatOptions{
		Include: container.ListBlobsInclude{Snapshots: true, Versions: true},
	})

	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		handleError(err)

		for _, blob := range resp.Segment.BlobItems {
			fmt.Println(*blob.Name)
		}
	}

	// Download the blob
	get, err := blobClient.DownloadStream(ctx, nil)
	handleError(err)

	downloadedData := bytes.Buffer{}
	reader := get.Body
	_, err = downloadedData.ReadFrom(reader)
	handleError(err)

	err = reader.Close()
	handleError(err)

	// Print the content of the blob we created
	fmt.Println("Blob contents:")
	fmt.Println(downloadedData.String())

	fmt.Printf("Press enter key to delete resources and exit the application.\n")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	fmt.Printf("Cleaning up.\n")

	// Delete the blob
	fmt.Printf("Deleting the blob " + blobName + "\n")

	_, err = blobClient.Delete(ctx, nil)
	handleError(err)

	// Delete the container
	fmt.Printf("Deleting the container " + containerName + "\n")
	_, err = containerClient.Delete(ctx, nil)
	handleError(err)
}
