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

* Install [Go](https://golang.org/dl/) 1.17.3 or later

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

## Populate Environment with Access Keys

> Note that in this step security credentials are handled that must be threaded carefully just like username/passwords

The application will fetch it's ***Access key*** from the environment. Export the strings you have copied in the last step:

```
export AZURE_STORAGE_ACCOUNT_NAME="<your_account_name>"
export AZURE_STORAGE_ACCOUNT_KEY="<your_secret_key>"
```

## Download and Install the Azure Storage Blob SDK for Go

From your GOPATH, execute the following command:

```bash
go get github.com/Azure/azure-sdk-for-go/sdk/storage/azblob
```

At this point, you can run this application. It creates an Azure storage container and blob object then cleans up after itself by deleting everything at the end.

## Run the application

Open the `storage-quickstart.go` file.

Replace `<StorageAccountName>` with the name of your Azure storage account.

Run the application with the `go run` command:

```bash
go run storage-quickstart.go
```

## More information

The [Azure storage documentation](https://docs.microsoft.com/azure/storage/) includes a rich set of tutorials and conceptual articles, which serve as a good complement to the samples. For more samples on the Azure Storage SDK for GO, check out the examples [here](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/storage/azblob).

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.
