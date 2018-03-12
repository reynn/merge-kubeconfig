# merge-kubeconfig

Merge configurations for multiple Kubernetes clusters into one file. This allows using `kubectl config use-context` to switch between clusters instead of manually switching configs.

## Usage

### Docker

```shell
docker run -v "$(pwd):/tmp:rw" quay.io/reynn/merge-kubeconfig /tmp/config1.yaml /tmp/config2.yaml /tmp/config3.yaml
```

### Local

```shell
go get github.com/reynn/merge-kubeconfig

merge-kubeconfig -outpath ~/.kube/config ~/.kube/config1.yaml ~/.kube/config2.yaml ~/.kube/config3.yaml
```

## Flags

| Flag      | Description                                                                               |
|-----------|-------------------------------------------------------------------------------------------|
| out       | What type of output, `yaml` will write out to a .yaml file, `text` will output to stdout. |
| outpath   | Path to write a `yaml` file to.                                                           |
| namespace | If provided will add a namespace to the output yaml config.                               |
