# Operator Notes

# Commands

```
OPERATOR_NAME=app-operator

operator-sdk new $OPERATOR_NAME --repo github.com/sachinmaharana/appsoperator

operator-sdk add api --api-version=sachinmaharana.com/v1 --kind=AppsOperator

operator-sdk add api --api-version=sachinmaharana.com/v1 --kind=AppsOperator

// changes in pkg/apis/sachinmaharana/v1/appsoperator_types.go -> Spec, Status

operator-sdk generate k8s

operator-sdk add controller --api-version=sachinmaharana.com/v1 --kind=AppsOperator

// make changes


kubectl apply -f deploy/crds/*_crd.yaml 

operator-sdk up local --namespace default 

OPERATOR_NAME=app-operator 

k apply -f deploy/crds/*_cr.yaml 
```