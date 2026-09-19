package vm

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Factory", func() {
	Describe("cgroupBind", func() {
		It("returns true only when all three flags are true", func() {
			Expect(cgroupBind(true, true, true)).To(BeTrue())
		})

		It("returns false when startContainersWithSystemD is false", func() {
			Expect(cgroupBind(false, true, true)).To(BeFalse())
		})

		It("returns false when MountCgroupfs is false", func() {
			Expect(cgroupBind(true, false, true)).To(BeFalse())
		})

		It("returns false by default when UseHostCgroupNamespace is false", func() {
			Expect(cgroupBind(true, true, false)).To(BeFalse())
		})
	})
})
