# net-echo

## docker

```shell
docker run -d --restart=unless-stopped --name net-echo -p 80:80  naturelr/net-echo:latest
```

## k8s

```shell
kubectl apply -f https://raw.githubusercontent.com/NatureLR/net-echo/master/k8s.yaml
```
