/*
 * This file is part of the kiagnose project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2023 Red Hat, Inc.
 *
 */

package client

import (
	"context"

	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"

	snapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1"
	"kubevirt.io/client-go/kubecli"
	cdiv1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

type Client struct {
	kubecli.KubevirtClient
}

func New() (*Client, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	client, err := kubecli.GetKubevirtClientFromRESTConfig(config)
	if err != nil {
		return nil, err
	}

	return &Client{client}, nil
}

func (c *Client) ListStorageClasses(ctx context.Context) (*storagev1.StorageClassList, error) {
	return c.KubevirtClient.StorageV1().StorageClasses().List(ctx, metav1.ListOptions{})
}

func (c *Client) ListStorageProfiles(ctx context.Context) (*cdiv1.StorageProfileList, error) {
	return c.KubevirtClient.CdiClient().CdiV1beta1().StorageProfiles().List(ctx, metav1.ListOptions{})
}

func (c *Client) ListVolumeSnapshotClasses(ctx context.Context) (*snapshotv1.VolumeSnapshotClassList, error) {
	return c.KubevirtClient.KubernetesSnapshotClient().SnapshotV1().VolumeSnapshotClasses().List(ctx, metav1.ListOptions{})
}
