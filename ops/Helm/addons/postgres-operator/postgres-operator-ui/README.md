# postgres-operator-ui

![Version: 1.6.0](https://img.shields.io/badge/Version-1.6.0-informational?style=flat-square) ![AppVersion: 1.6.0](https://img.shields.io/badge/AppVersion-1.6.0-informational?style=flat-square)

Postgres Operator UI provides a graphical interface for a convenient database-as-a-service user experience

**Homepage:** <https://github.com/zalando/postgres-operator>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| Zalando | opensource@zalando.de |  |

## Source Code

* <https://github.com/zalando/postgres-operator>

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| envs.operatorApiUrl | string | `"http://postgres-operator:8080"` |  |
| envs.operatorClusterNameLabel | string | `"cluster-name"` |  |
| envs.resourcesVisible | string | `"False"` |  |
| envs.targetNamespace | string | `"default"` |  |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.registry | string | `"registry.opensource.zalan.do"` |  |
| image.repository | string | `"acid/postgres-operator-ui"` |  |
| image.tag | string | `"v1.6.0"` |  |
| ingress.annotations."cert-manager.io/cluster-issuer" | string | `"cert-manager-production"` |  |
| ingress.annotations."kubernetes.io/ingress.class" | string | `"nginx"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/app-root" | string | `"/"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity" | string | `"true"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/enable-opentracing" | string | `"true"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules" | string | `"true"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/rewrite-target" | string | `"/"` |  |
| ingress.enabled | bool | `true` |  |
| ingress.hosts[0].host | string | `"shortlink.ddns.net"` |  |
| ingress.hosts[0].paths[0] | string | `"/psql-ui"` |  |
| ingress.tls[0].hosts[0] | string | `"shortlink.ddns.net"` |  |
| ingress.tls[0].secretName | string | `"shortlink-ingress-tls"` |  |
| rbac.create | bool | `true` |  |
| replicaCount | int | `1` |  |
| resources.limits.cpu | string | `"200m"` |  |
| resources.limits.memory | string | `"200Mi"` |  |
| resources.requests.cpu | string | `"100m"` |  |
| resources.requests.memory | string | `"100Mi"` |  |
| service.port | string | `"80"` |  |
| service.type | string | `"LoadBalancer"` |  |
| serviceAccount.create | bool | `true` |  |
| serviceAccount.name | string | `nil` |  |
