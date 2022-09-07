# uuid

kubectl delete secret -n staging-dev client-token --ignore-not-found
kubectl create secret -n staging-dev generic --dry-run=client client-token --from-file=uuid=token.tmp