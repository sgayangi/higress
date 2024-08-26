// Copyright (c) 2022 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package option

import (
	"fmt"
	"os"
)

const optionYAML = `# File generated by hgctl. Modify as required.

version: 1.0.0

build:
  # The official builder image version
  builder:
    go: 1.19
    tinygo: 0.28.1
    oras: 1.0.0
  # The WASM plugin project directory
  input: ./
  # The output of the build products
  output:
  # Choose between 'files' and 'image'
    type: files
    # Destination address: when type=files, specify the local directory path, e.g., './out' or
    # type=image, specify the remote docker repository, e.g., 'docker.io/<your_username>/<your_image>'
    dest: ./out
  # The authentication configuration for pushing image to the docker repository
  docker-auth: ~/.docker/config.json
  # The directory for the WASM plugin configuration structure
  model-dir: ./
  # The WASM plugin configuration structure name
  model: PluginConfig
  # Enable debug mode
  debug: false

test:
  # Test environment name, that is a docker compose project name
  name: wasm-test
  # The output path to build products, that is the source of test configuration parameters
  from-path: ./out
  # The test configuration source
  test-path: ./test
  # Docker compose configuration, which is empty, looks for the following files from 'test-path':
  # compose.yaml, compose.yml, docker-compose.yml, docker-compose.yaml
  compose-file:
  # Detached mode: Run containers in the background
  detach: false

install:
  # The namespace of the installation
  namespace: higress-system
  # Use to validate WASM plugin configuration when install by yaml
  spec-yaml: ./out/spec.yaml
  # Installation source. Choose between 'from-yaml' and 'from-go-project'
  from-yaml: ./test/plugin-conf.yaml
  # If 'from-go-src' is non-empty, the output type of the build option must be 'image'
  from-go-src:
  # Enable debug mode
  debug: false
`

func GenOptionYAML(dir string) error {
	path := fmt.Sprintf("%s/option.yaml", dir)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(optionYAML); err != nil {
		return err
	}

	return nil
}