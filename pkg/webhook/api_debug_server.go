/*
Copyright 2022 The Koordinator Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package webhook

import (
	"net/http"

	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

var debugAPIProviderMap = map[string]http.Handler{}

func RegisterDebugAPIProvider(name string, provider http.Handler) {
	debugAPIProviderMap[name] = provider
}

func InstallDebugAPIHandler(server *webhook.Server) {
	for name, provider := range debugAPIProviderMap {
		server.Register(name, provider)
		klog.Infof("Success register debug api handler, name:%v, tcpAddr:%s:%d", name, server.Host, server.Port)
	}
}
