//+build e2e

/*
Copyright 2020 Clastix Labs.

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

package e2e

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/clastix/capsule/api/v1alpha1"
)

var _ = Describe("creating a Tenant with wrong name", func() {
	tnt := &v1alpha1.Tenant{
		ObjectMeta: metav1.ObjectMeta{
			Name: "wrong-name",
		},
		Spec: v1alpha1.TenantSpec{
			Owner: v1alpha1.OwnerSpec{
				Name: "john",
				Kind: "User",
			},
			NamespacesMetadata: v1alpha1.AdditionalMetadata{},
			ServicesMetadata:   v1alpha1.AdditionalMetadata{},
			IngressClasses:     v1alpha1.IngressClassesSpec{},
			StorageClasses:     v1alpha1.StorageClassesSpec{},
			LimitRanges:        []corev1.LimitRangeSpec{},
			NamespaceQuota:     10,
			NodeSelector:       map[string]string{},
			ResourceQuota:      []corev1.ResourceQuotaSpec{},
		},
	}
	It("should fail", func() {
		Expect(k8sClient.Create(context.TODO(), tnt)).ShouldNot(Succeed())
	})
})
