helm install dapr dapr/dapr --namespace dapr-system --create-namespace --wait

helm repo add bitnami https://charts.bitnami.com/bitnami

helm install redis bitnami/redis
