# shortlink-landing

![Version: 0.5.12](https://img.shields.io/badge/Version-0.5.12-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

Shortlink landing service

**Homepage:** <https://batazor.github.io/shortlink/>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | batazor111@gmail.com | batazor.ru |

## Source Code

* <https://github.com/batazor/shortlink>

## Requirements

Kubernetes: `>= 1.19.0 || >= v1.19.0-0`

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| deploy.affinity | list | `[]` |  |
| deploy.annotations | object | `{}` |  |
| deploy.image.pullPolicy | string | `"IfNotPresent"` |  |
| deploy.image.repository | string | `"batazor/shortlink-landing"` |  |
| deploy.image.tag | string | `"latest"` |  |
| deploy.imagePullSecrets | list | `[]` |  |
| deploy.livenessProbe.failureThreshold | int | `1` |  |
| deploy.livenessProbe.httpGet.path | string | `"/landing/"` |  |
| deploy.livenessProbe.httpGet.port | int | `80` |  |
| deploy.livenessProbe.initialDelaySeconds | int | `10` |  |
| deploy.livenessProbe.periodSeconds | int | `30` |  |
| deploy.livenessProbe.successThreshold | int | `1` |  |
| deploy.nodeSelector | list | `[]` |  |
| deploy.podSecurityContext.fsGroup | int | `1000` |  |
| deploy.readinessProbe.failureThreshold | int | `30` |  |
| deploy.readinessProbe.httpGet.path | string | `"/landing/"` |  |
| deploy.readinessProbe.httpGet.port | int | `80` |  |
| deploy.readinessProbe.initialDelaySeconds | int | `10` |  |
| deploy.readinessProbe.periodSeconds | int | `30` |  |
| deploy.readinessProbe.successThreshold | int | `1` |  |
| deploy.replicaCount | int | `1` |  |
| deploy.resources.limits.cpu | string | `"100m"` |  |
| deploy.resources.limits.memory | string | `"128Mi"` |  |
| deploy.resources.requests.cpu | string | `"100m"` |  |
| deploy.resources.requests.memory | string | `"128Mi"` |  |
| deploy.securityContext.allowPrivilegeEscalation | bool | `false` |  |
| deploy.strategy.rollingUpdate.maxSurge | int | `1` |  |
| deploy.strategy.rollingUpdate.maxUnavailable | int | `0` |  |
| deploy.strategy.type | string | `"RollingUpdate"` |  |
| deploy.terminationGracePeriodSeconds | int | `90` |  |
| deploy.tolerations | list | `[]` |  |
| enabled | bool | `true` |  |
| fullnameOverride | string | `""` |  |
| host | string | `"shortlink.ddns.net"` |  |
| ingress.annotations."cert-manager.io/cluster-issuer" | string | `"cert-manager-production"` |  |
| ingress.annotations."kubernetes.io/ingress.class" | string | `"nginx"` |  |
| ingress.annotations."kubernetes.io/tls-acme" | string | `"true"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity" | string | `"true"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/enable-opentracing" | string | `"true"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules" | string | `"true"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/rewrite-target" | string | `"/landing/$2"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/use-regex" | string | `"true"` |  |
| ingress.enabled | bool | `true` |  |
| ingress.tls[0].hosts[0] | string | `"shortlink.ddns.net"` |  |
| ingress.tls[0].secretName | string | `"shortlink-ingress-tls"` |  |
| nameOverride | string | `""` |  |
| service.port | int | `80` |  |
| service.type | string | `"ClusterIP"` |  |
| serviceAccount.create | bool | `true` |  |
| serviceAccount.name | string | `"shortlink"` |  |
