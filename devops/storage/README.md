Based (mostly) on https://www.digitalocean.com/community/tutorials/how-to-deploy-postgres-to-kubernetes-cluster

```bash
$ kubectl apply -f persistent-volume.yaml
persistentvolume/postgres-pv created
$ kubectl apply -f persistent-volume-claim.yaml
persistentvolumeclaim/postgres-pvc created
```

