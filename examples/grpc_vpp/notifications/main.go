// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"io"
	"os"
	"time"

	"github.com/ligato/cn-infra/core"
	"github.com/ligato/cn-infra/flavors/local"
	"github.com/ligato/cn-infra/logging"
	log "github.com/ligato/cn-infra/logging/logrus"
	"github.com/ligato/vpp-agent/plugins/defaultplugins/common/model/rpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const defaultAddress = "localhost:9111"

var address = defaultAddress

// init sets the default logging level
func init() {
	log.DefaultLogger().SetOutput(os.Stdout)
	log.DefaultLogger().SetLevel(logging.DebugLevel)
}

// Start Agent plugins selected for this example.
func main() {
	// Init close channel to stop the example.
	closeChannel := make(chan struct{}, 1)

	flag.StringVar(&address, "address", defaultAddress, "address of GRPC server")

	// Example plugin
	agent := local.NewAgent(local.WithPlugins(func(flavor *local.FlavorLocal) []*core.NamedPlugin {
		examplePlugin := &core.NamedPlugin{PluginName: PluginID, Plugin: &ExamplePlugin{}}

		return []*core.NamedPlugin{{examplePlugin.PluginName, examplePlugin}}
	}))

	core.EventLoopWithInterrupt(agent, closeChannel)
}

// PluginID of example plugin
const PluginID core.PluginName = "example-plugin"

// ExamplePlugin demonstrates the use of the remoteclient to locally transport example configuration into the default VPP plugins.
type ExamplePlugin struct {
	conn *grpc.ClientConn
}

// Init initializes example plugin.
func (plugin *ExamplePlugin) Init() (err error) {
	// Set up connection to the server.
	plugin.conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return err
	}

	// Start notification watcher.
	go plugin.watchNotifications()

	log.DefaultLogger().Info("Initialization of the example plugin has completed")
	return err
}

// Close does nothing
func (plugin *ExamplePlugin) Close() error {
	return nil
}

// Get is an implementation of client-side statistics streaming.
func (plugin *ExamplePlugin) watchNotifications() {
	// Index of last received notification
	var lastIndex uint32

	for {
		// Get client for notification service
		client := rpc.NewNotificationServiceClient(plugin.conn)
		// Prepare request with the last index
		request := &rpc.FromIndex{
			Index: uint32(lastIndex),
		}
		// Get stream object
		stream, err := client.Get(context.Background(), request)
		if err != nil {
			log.DefaultLogger().Error(err)
			return
		}
		// Receive all message from the stream
		log.DefaultLogger().Info("Sending request ... ")
		for {
			notification, err := stream.Recv()
			if err == io.EOF {
				log.DefaultLogger().Info("All new notifications received")
				break
			}
			if err != nil {
				log.DefaultLogger().Error(err)
				return
			}

			log.DefaultLogger().Infof("Received notification: %v (index: %d)", notification.IfNotif, notification.Index)
			lastIndex = notification.Index
		}

		// Wait till next request
		time.Sleep(3 * time.Second)

	}
}
