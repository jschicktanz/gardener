/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/gardener/gardener/pkg/client/extensions/clientset/versioned/typed/extensions/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeExtensionsV1alpha1 struct {
	*testing.Fake
}

func (c *FakeExtensionsV1alpha1) BackupBuckets() v1alpha1.BackupBucketInterface {
	return &FakeBackupBuckets{c}
}

func (c *FakeExtensionsV1alpha1) BackupEntries() v1alpha1.BackupEntryInterface {
	return &FakeBackupEntries{c}
}

func (c *FakeExtensionsV1alpha1) Bastions(namespace string) v1alpha1.BastionInterface {
	return &FakeBastions{c, namespace}
}

func (c *FakeExtensionsV1alpha1) Clusters() v1alpha1.ClusterInterface {
	return &FakeClusters{c}
}

func (c *FakeExtensionsV1alpha1) ContainerRuntimes(namespace string) v1alpha1.ContainerRuntimeInterface {
	return &FakeContainerRuntimes{c, namespace}
}

func (c *FakeExtensionsV1alpha1) ControlPlanes(namespace string) v1alpha1.ControlPlaneInterface {
	return &FakeControlPlanes{c, namespace}
}

func (c *FakeExtensionsV1alpha1) DNSRecords(namespace string) v1alpha1.DNSRecordInterface {
	return &FakeDNSRecords{c, namespace}
}

func (c *FakeExtensionsV1alpha1) Extensions(namespace string) v1alpha1.ExtensionInterface {
	return &FakeExtensions{c, namespace}
}

func (c *FakeExtensionsV1alpha1) Infrastructures(namespace string) v1alpha1.InfrastructureInterface {
	return &FakeInfrastructures{c, namespace}
}

func (c *FakeExtensionsV1alpha1) Networks(namespace string) v1alpha1.NetworkInterface {
	return &FakeNetworks{c, namespace}
}

func (c *FakeExtensionsV1alpha1) OperatingSystemConfigs(namespace string) v1alpha1.OperatingSystemConfigInterface {
	return &FakeOperatingSystemConfigs{c, namespace}
}

func (c *FakeExtensionsV1alpha1) Workers(namespace string) v1alpha1.WorkerInterface {
	return &FakeWorkers{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeExtensionsV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}