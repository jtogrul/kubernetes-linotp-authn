curl -X POST -k -H "Content-Type: application/json" -H "Cache-Control: no-cache" -H "Postman-Token: 002892dc-352f-3a27-8824-95bf459c3518" -d '{
  "apiVersion": "authentication.k8s.io/v1beta1",
  "kind": "TokenReview",
  "spec": {
    "token": "dGphZmFybGk6Njg1NTY0"
  }
}' "https://127.0.0.1:8765/"