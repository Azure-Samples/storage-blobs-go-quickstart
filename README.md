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

1. Go to the [Azure portal](https://portal.azure.com/#create/Microsoft.StorageAccount-ARM) create a storage account menu.
2. Enter a unique name for your storage account. Keep these rules in mind for naming your storage account:
    - The name must be between 3 and 24 characters in length.
    - The name may contain numbers and lowercase letters only.
3. Select your subscription.
4. For **Resource group**, create a new one or use an existing resource group.
5. Select the **Location** to use for your storage account.
6. Click **Create** to create your storage account.

## Sign in with Azure CLI

To support local development, the `DefaultAzureCredential` can authenticate as the user signed into the Azure CLI.

Run the following command to sign into the Azure CLI.

```azurecli
az login
```

## Assign RBAC permissions to the storage account

Azure storage accounts require explicit permissions to perform read and write operations. In order to use the storage account, you must assign permissions to the account. To do that you'll need to assing an appropriate RBAC role to your account. To get the `objectID` of the currently signed in user, run `az ad signed-in-user show --query objectId`.

Run the following AzureCli command to assign the storage account permissions:

```azurecli
az role assignment create --assignee "<ObjectID>" --role "Storage Blob Data Contributor" --scope "<StorageAccountResourceID>"
```

Learn more about Azure's built-in RBAC roles, click [here](https://docs.microsoft.com/azure/role-based-access-control/built-in-roles).

> Note: Azure Cli has built in helper fucntions that retrieve the storage access keys when permissions are not detected. That functionally does not transfer to the DefaultAzureCredential, which is the reason for assiging RBAC roles to your account.

## Download and Install the Azure Storage Blob SDK for Go

From your GOPATH, execute the following command:

```bash
go get github.com/Azure/azure-sdk-for-go/sdk/azidentity

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
