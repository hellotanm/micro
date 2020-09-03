package router

import (
	"github.com/micro/go-micro/v3/registry"
)

type ServiceFilter func([]*registry.Service) []*registry.Service

// FilterLabel is a label based Select Filter which will
// only return services with the label specified.
func FilterLabel(key, val string) ServiceFilter {
	return func(old []*registry.Service) []*registry.Service {
		var services []*registry.Service

		for _, service := range old {
			serv := new(registry.Service)
			var nodes []*registry.Node

			for _, node := range service.Nodes {
				if node.Metadata == nil {
					continue
				}

				if node.Metadata[key] == val {
					nodes = append(nodes, node)
				}
			}

			// only add service if there's some nodes
			if len(nodes) > 0 {
				// copy
				*serv = *service
				serv.Nodes = nodes
				services = append(services, serv)
			}
		}

		return services
	}
}
