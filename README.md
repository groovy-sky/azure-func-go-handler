# Azure Functions custom handler in Go

This repository contains an example of an Azure Functions custom handler in Go. 

You can use Azure Cloud Shell in Bash mode for deploying and building Functions environment. Just execute script below:

```
deploy_region="westeurope"                                                                                    
deploy_group="go-custom-function"                                                                            
az group create -l $deploy_region -n $deploy_group                                                            
function=$(az deployment group create --resource-group $deploy_group --template-uri https://raw.githubusercontent.com/groovy-sky/azure-func-go-handler/master/Template/azuredeploy.json | jq -r '. | .properties | .dependencies | .[] | .resourceName')                                                     
[ ! -d "azure-func-go-handler/.git" ] && git clone https://github.com/groovy-sky/azure-func-go-handler        
cd azure-func-go-handler/Function && git pull                                                                 
go build *.go && func azure functionapp publish $function --no-build --force


```

If the publishing was successful you can validate a result by accessing a newly created function.

Exact instruction how-to run you can find [here](https://github.com/groovy-sky/azure/tree/master/func-custom-handler-00#introduction)

