# logsync-crd

## How to use it
### yaml
```
apiVersion: logging.yo.logsync/v1
kind: LogSync
metadata:
  name: logsync-sample
spec:
  projectID: projectID
  destination: storage.googleapis.com/okan-tfstate
  filter: severity >= ERROR
```

### command
kubectl apply -f config/samples/logging_v1_logsync.yaml
