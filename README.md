Nothing to see here. Just experimentation.

## Prerequisites
1. Kubernetes cluster
2. Latest running Tiller `helm init`
3. Ingress Controller
```
helm install --namespace kube-system --set dashboard.domain=traefik.helm.tucker.me stable/traefik
```
4. Wildcard domain name pointed to ingress service load balancer IP.

## Building
```bash
make build
```

## Deploying
```bash
helm install --set dashboard.domain=traefik.helm.tucker.me charts/helm-ui
```


// helm install --name traefik --namespace kube-system --set dashboard.enabled=true,dashboard.domain=traefik.helm.tucker.me stable/traefik
