/*
Copyright 2020 The SuperEdge Authors.

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

package steps

import (
	"github.com/superedge/superedge/pkg/edgeadm/common"
	"github.com/superedge/superedge/pkg/edgeadm/constant"
	"github.com/superedge/superedge/pkg/util"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow"
	cmdUtil "k8s.io/kubernetes/cmd/kubeadm/app/cmd/util"
)

var (
	containerDExample = cmdUtil.Examples(`
		# Install containerD container runtime.
		  kubeadm init phase container`)
)

//install containerD container runtime
func NewContainerDPhase() workflow.Phase {
	return workflow.Phase{
		Name:         "containerD",
		Short:        "Install containerD container runtime",
		Long:         "Install containerD container runtime",
		Example:      containerDExample,
		Run:          installContainerD,
		InheritFlags: []string{},
	}
}

func installContainerD(c workflow.RunData) error {

	klog.V(4).Infof("Start install containerD container runtime")
	if err := common.UnzipPackage(EdgeadmConf.WorkerPath+constant.ZipContainerPath, EdgeadmConf.WorkerPath+constant.UnZipContainerDstPath); err != nil {
		klog.Errorf("Unzip containerD container runtime Package: %s, error: %v", EdgeadmConf.WorkerPath+constant.UnZipContainerDstPath, err)
		return err
	}
	if err := common.UnzipPackage(EdgeadmConf.WorkerPath+constant.ZipContainerPath, EdgeadmConf.WorkerPath+constant.UnZipContainerDstPath); err != nil {
		klog.Errorf("Unzip containerD container runtime Package: %s, error: %v", EdgeadmConf.WorkerPath+constant.UnZipContainerDstPath, err)
		return err
	}
	if _, _, err := util.RunLinuxCommand(EdgeadmConf.WorkerPath + constant.ContainerDInstallShell); err != nil {
		klog.Errorf("Run containerD container runtime install shell: %s, error: %v",
			EdgeadmConf.WorkerPath+constant.UnZipContainerDstPath, err)
		return err
	}

	klog.V(4).Infof("Installed containerD container runtime successfully")
	return nil
}
