package validate

import (
  "io/ioutil"
  "log"

  "gopkg.in/yaml.v2"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)


var _ = Describe("Lab 1 Containers", func() {
	var failMessage string

	BeforeEach(func() {
		failMessage = ""
	})

	Context("Step 2", func() {
		It("should have a Dockerfile", func() {
			failMessage = "Dockerfile Doesn't Exist or is in the wrong location\n"
			Expect(fileExists("../Dockerfile")).To(Succeed(), failMessage)
		})
	})

	Context("Step 3", func() {
		It("should have a skaffold.yaml file", func() {
			failMessage = "skaffold.yaml Doesn't Exist or is in the wrong location\n"
			Expect(fileExists("../skaffold.yaml")).To(Succeed(), failMessage)
		})

		It("should be a valid skaffold configuration", func() {
			var skaffold interface{}
			skaffoldFile, err := ioutil.ReadFile("../skaffold.yaml")
			if err != nil {
				Skip("skaffold.yaml not found")
			}

			err = yaml.Unmarshal(skaffoldFile, &skaffold)
			Expect(err).ToNot(HaveOccurred())
			//failures := InterceptGomegaFailures(func() {

			failMessage = "Incorrect apiVersion in skaffold.yaml\n"

			Expect(treeValue(skaffold, []interface{}{"apiVersion"})).To(Equal("skaffold/v1beta12"), failMessage)
			failMessage = "First build artifact in skaffold.yaml should be \"springtrader\"\n"

			Expect(treeValue(skaffold, []interface{}{"build", "artifacts", 0, "image"})).To(Equal("springtrader"), failMessage)
			failMessage = "Second build artifact in skaffold.yaml should be \"sqlfdb\"\n"

			Expect(treeValue(skaffold, []interface{}{"build", "artifacts", 1, "image"})).To(Equal("sqlfdb"), failMessage)
			//})
			//log.Printf(failures[0])
		})
	})

	AfterEach(func() {
		log.Printf("%v\n", CurrentGinkgoTestDescription())
		if CurrentGinkgoTestDescription().Failed {
			ConcatenatedMessage += failMessage
		}
	})
})
