## Get started

You can start a mock server of [container registry](https://distribution.github.io/distribution/) with below command:

```shell
atest mock --files mock/image-registry.yaml --prefix /
```

then, you can pull images from it:

```shell
docker pull localhost:6060/anyRepo/anyName:anyTag
```
