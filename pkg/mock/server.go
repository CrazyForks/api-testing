/*
Copyright 2024 API Testing Authors.

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
package mock

import (
	"io"
	"net/http"
)

type Loadable interface {
	Load() error
}

type DynamicServer interface {
	Start(reader Reader, prefix string) error
	SetupHandler(reader Reader, prefix string) (http.Handler, error)
	WithTLS(certFile, keyFile string) DynamicServer
	WithLogWriter(writer io.Writer) DynamicServer
	GetTLS() (certFile, keyFile string)
	Stop() error
	GetPort() string
	EnableMetrics()
	Loadable
}

const (
	headerMockServer = "Mock-Server"
)
