package cluster

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
    p.AddResourceConfigurator("confluent_kafka_cluster", func(r *config.Resource) {
        r.ShortGroup = "KafkaCluster"
    })
}
