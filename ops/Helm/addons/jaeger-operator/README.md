# jaeger-operator

![Version: 0.2.8](https://img.shields.io/badge/Version-0.2.8-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | batazor111@gmail.com | batazor.ru |

## Requirements

Kubernetes: `>= 1.19.0 || >= v1.19.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://jaegertracing.github.io/helm-charts | jaeger-operator | 2.18.3 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| ingress.annotations."cert-manager.io/cluster-issuer" | string | `"cert-manager-production"` |  |
| ingress.annotations."kubernetes.io/ingress.class" | string | `"nginx"` |  |
| ingress.annotations."kubernetes.io/tls-acme" | string | `"true"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity" | string | `"true"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/enable-opentracing" | string | `"true"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules" | string | `"true"` |  |
| ingress.enabled | bool | `true` |  |
| ingress.tls[0].hosts[0] | string | `"shortlink.ddns.net"` |  |
| ingress.tls[0].secretName | string | `"shortlink-ingress-tls"` |  |
| jaeger-operator.enabled | bool | `true` |  |
