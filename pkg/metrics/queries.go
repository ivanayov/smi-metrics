package metrics

var (
	resourceQueries = map[string]string{
		"p99_response_latency": `histogram_quantile(
			0.99,
			sum(
				irate(
					response_latency_ms_bucket{
						namespace=~"{{default ".+" .namespace }}",
						{{lower .kind}}=~"{{default ".+" .name }}"
					}[{{.window}}]
				)
			) by (
				{{lower .kind}},
				namespace,
				le
			)
		)`,
		"p90_response_latency": `histogram_quantile(
			0.90,
			sum(
				irate(
					response_latency_ms_bucket{
						namespace=~"{{default ".+" .namespace }}",
						{{lower .kind}}=~"{{default ".+" .name }}"
					}[{{.window}}]
				)
			) by (
				{{lower .kind}},
				namespace,
				le
			)
		)`,
		"p50_response_latency": `histogram_quantile(
			0.50,
			sum(
				irate(
					response_latency_ms_bucket{
						namespace=~"{{default ".+" .namespace }}",
						{{lower .kind}}=~"{{default ".+" .name }}"
					}[{{.window}}]
				)
			) by (
				{{lower .kind}},
				namespace,
				le
			)
		)`,
		"success_count": `sum(
			increase(
				response_total{
					classification="success",
					namespace=~"{{default ".+" .namespace }}",
					{{lower .kind}}=~"{{default ".+" .name }}"
				}[{{.window}}]
			)
		) by (
			{{lower .kind}},
			namespace
		)`,
		"failure_count": `sum(
			increase(
				response_total{
					classification="failure",
					namespace=~"{{default ".+" .namespace }}",
					{{lower .kind}}=~"{{default ".+" .name }}"
				}[{{.window}}]
			)
		) by (
			{{lower .kind}},
			namespace
		)`,
	}

	edgeQueries = map[string]string{
		"p99_response_latency": `histogram_quantile(
			0.99,
			sum(
				irate(
					response_latency_ms_bucket{
						namespace=~"{{.namespace}}",
						{{lower .kind}}=~"{{default ".+" .fromName}}",
						dst_{{lower .kind}}=~"{{default ".+" .toName}}"
					}[{{.window}}]
				)
			) by (
				{{lower .kind}},
				dst_{{lower .kind}},
				namespace,
				dst_namespace,
				le
			)
		)`,
		"p90_response_latency": `histogram_quantile(
			0.90,
			sum(
				irate(
					response_latency_ms_bucket{
						namespace=~"{{.namespace}}",
						{{lower .kind}}=~"{{default ".+" .fromName}}",
						dst_{{lower .kind}}=~"{{default ".+" .toName}}"
					}[{{.window}}]
				)
			) by (
				{{lower .kind}},
				dst_{{lower .kind}},
				namespace,
				dst_namespace,
				le
			)
		)`,
		"p50_response_latency": `histogram_quantile(
			0.50,
			sum(
				irate(
					response_latency_ms_bucket{
						namespace=~"{{.namespace}}",
						{{lower .kind}}=~"{{default ".+" .fromName}}",
						dst_{{lower .kind}}=~"{{default ".+" .toName}}"
					}[{{.window}}]
				)
			) by (
				{{lower .kind}},
				dst_{{lower .kind}},
				namespace,
				dst_namespace,
				le
			)
		)`,
		"success_count": `sum(
			increase(
				response_total{
					classification="success",
					namespace=~"{{.namespace}}",
					{{lower .kind}}=~"{{default ".+" .fromName}}",
					dst_{{lower .kind}}=~"{{default ".+" .toName}}"
				}[{{.window}}]
			)
		) by (
			{{lower .kind}},
			dst_{{lower .kind}},
			namespace,
			dst_namespace
		)`,
		"failure_count": `sum(
			increase(
				response_total{
					classification="failure",
					namespace=~"{{.namespace}}",
					{{lower .kind}}=~"{{default ".+" .fromName}}",
					dst_{{lower .kind}}=~"{{default ".+" .toName}}"
				}[{{.window}}]
			)
		) by (
			{{lower .kind}},
			dst_{{lower .kind}},
			namespace,
			dst_namespace
		)`,
	}
)
