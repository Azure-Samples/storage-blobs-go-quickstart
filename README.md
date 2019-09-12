---
page_type: sample
languages:
- go
products:
- azure
description: "This repository contains a simple sample project to help you getting started with Azure storage using Go as the development language."
urlFragment: storage-blobs-go-quickstart
---

# Transfer objects to and from Azure Blob storage using Go

This repository contains a simple sample project to help you getting started with Azure storage using Go as the development language.

## Prerequisites

To complete this tutorial:

* Install [Go](https://golang.org/dl/) 1.8 or later 

If you don't have an Azure subscription, create a [free account](https://portal.azure.com/#create/Microsoft.StorageAccount-ARM) before you begin.

## Create a storage account using the Azure portal

First, create a new general-purpose storage account to use for this quickstart.

1. Go to the [Azure portal](https://portal.azure.com/#create/Microsoft.StorageAccount-ARM) account creation menu and log in using your Azure account. 
2. Enter a unique name for your storage account. Keep these rules in mind for naming your storage account:
    - The name must be between 3 and 24 characters in length.
    - The name may contain numbers and lowercase letters only.
3. Select your subscription. 
4. For **Resource group**, create a new one or use an existing resource group. 
5. Select the **Location** to use for your storage account.
6. Check **Pin to dashboard** and click **Create** to create your storage account. 

After your storage account is created, it's pinned to the dashboard. Select it to open it. Under Settings, select **Access keys**. Copy and paste the Storage account name and the Key under **key1** into a text editor for later use.

## Put the account name and key in environment variables

This solution requires your storage account name and key to be stored in environment variables securely on the machine running the sample. Follow one of the examples below depending on your operating System to create the environment variables. If using Windows, close out of your open IDE or shell and restart to ensure that the environment variables are initialized.

### Linux

```bash
export AZURE_STORAGE_ACCOUNT="<youraccountname>"
export AZURE_STORAGE_ACCESS_KEY="<youraccountkey>"
```
### Windows

```cmd
setx AZURE_STORAGE_ACCOUNT "<youracountname>"
setx AZURE_STORAGE_ACCESS_KEY "<youraccountkey>"
```

## Download and Install the Azure Storage Blob SDK for Go

From your GOPATH, execute the following command:
```
go get github.com/Azure/azure-storage-blob-go/2016-05-31/azblob
```

At this point, you can run this application. It creates its own file to upload, and then cleans up after itself by deleting everything at the end.

## Run the application

Navigate to your application directory and run the application with the go run command.

```
go run storage-quickstart.go
```

## More information

The [Azure storage documentation](https://docs.microsoft.com/azure/storage/) includes a rich set of tutorials and conceptual articles, which serve as a good complement to the samples. For more samples on the Azure Storage SDK for GO, check out the examples [here](https://godoc.org/github.com/Azure/azure-storage-blob-go/2016-05-31/azblob).

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.
