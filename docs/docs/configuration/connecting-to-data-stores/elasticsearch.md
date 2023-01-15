# ElasticSearch

If you want to use ElasticSearch as the trace data store, you can configure Tracetest to fetch trace data from ElasticSearch.

You'll configure the OpenTelemetry Collector to receive traces from your system and then send them to ElasticSearch via Elastic APM. And, you don't have to change your existing pipelines to do so.

:::tip
Examples of configuring Tracetest can be found in the [`examples` folder of the Tracetest GitHub repo](https://github.com/kubeshop/tracetest/tree/main/examples). 
:::

## Configure OpenTelemetry Collector to Send Traces to ElasticSearch

In your OpenTelemetry Collector config file, make sure to set the `exporter` to `otlp`, with the `endpoint` pointing to the Elastic APM server on port `8200`. If you are running Tracetest with Docker, the endpoint might look like this `apm-server:8200`.

```yaml
# collector.config.yaml

# If you already have receivers declared, you can just ignore
# this one and use yours instead.
receivers:
  otlp:
    protocols:
      grpc:
      http:


processors:
  batch:
    timeout: 100ms

exporters:
  otlp/elastic:
    endpoint: apm-server:8200
    tls:
      insecure: true
      insecure_skip_verify: true

service:
  pipelines:
    # You probably already have a traces pipeline, you don't have to change it.
    # Just add this one to your configuration. Just make sure to not have two
    # pipelines with the same name.
    traces/1:
      receivers: [otlp] # your receiver
      processors: [batch] # make sure to add the batch processor
      exporters: [otlp/elastic] # your exporter pointing to your Elastic APM server instance

```

## Configure Tracetest to Use ElasticSearch as a Trace Data Store

You also have to configure your Tracetest instance to make it aware that it has to fetch trace data from ElasticSearch. 

### Web UI



Edit your configuration file to include this configuration:

```yaml
# tracetest.config.yaml

postgresConnString: "host=postgres user=postgres password=postgres port=5432 sslmode=disable"

poolingConfig:
  maxWaitTimeForTrace: 10m
  retryDelay: 5s

googleAnalytics:
  enabled: true

demo:
  enabled: []

experimentalFeatures: []

telemetry:
  dataStores:
    opensearch:
      type: opensearch
      elasticsearch:
        addresses:
          - https://es01:9200
        username: elastic
        password: changeme
        index: apm-*

  exporters:
    collector:
      serviceName: tracetest
      sampling: 100 # 100%
      exporter:
        type: collector
        collector:
          endpoint: otel-collector:4317

server:
  telemetry:
    dataStore: opensearch
    exporter: collector

```
