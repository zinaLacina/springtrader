package validate

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)


var _ = Describe("Lab 1 Containers", func() {
	Context("Step 2", func() {
		It("should have a Dockerfile", func() {
			Expect(fileExists("../Dockerfile")).To(Succeed(), "Dockerfile was not found in the springtrader folder. Maybe the file was not saved or it was saved to a different folder.")
		})
	})

	Context("Step 3", func() {
		It("should have a skaffold.yaml file", func() {
			Expect(fileExists("../skaffold.yaml")).To(Succeed(), "skaffold.yaml was not found in the springtrader folder. Maybe the file was not saved or it was save in a different folder.")
			Fail("foo")
		})

		It("should be a valid skaffold configuration", func() {
			var skaffold interface{}
			skaffoldFile, err := ioutil.ReadFile("../skaffold.yaml")
			if err != nil {
				Skip("skaffold.yaml not found")
			}
			err = yaml.Unmarshal(skaffoldFile, &skaffold)
			Expect(err).ToNot(HaveOccurred())
			Expect(treeValue(skaffold, []interface{}{"apiVersion"})).To(Equal("skaffold/v1beta12"))
			Expect(treeValue(skaffold, []interface{}{"build", "artifacts", 0, "image"})).To(Equal("springtrader"))
			Expect(treeValue(skaffold, []interface{}{"build", "artifacts", 1, "image"})).To(Equal("sqlfdb"))
		})
	})
})
