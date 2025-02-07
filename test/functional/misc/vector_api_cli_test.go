// go:build !fluentd
package misc

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	logging "github.com/openshift/cluster-logging-operator/apis/logging/v1"
	"github.com/openshift/cluster-logging-operator/internal/constants"
	"github.com/openshift/cluster-logging-operator/test/framework/functional"
	testfw "github.com/openshift/cluster-logging-operator/test/functional"
)

var _ = Describe("[Functional][Misc][API_CLI] Functional test", func() {

	if testfw.LogCollectionType != logging.LogCollectionTypeVector {
		defer GinkgoRecover()
		Skip("skip for non-vector")
	}

	var framework *functional.CollectorFunctionalFramework

	BeforeEach(func() {
		Expect(testfw.LogCollectionType).To(Equal(logging.LogCollectionTypeVector))
		framework = functional.NewCollectorFunctionalFrameworkUsingCollector(logging.LogCollectionTypeVector)
		functional.NewClusterLogForwarderBuilder(framework.Forwarder).FromInput(logging.InputNameInfrastructure).ToHttpOutput()
	})

	AfterEach(func() {
		framework.Cleanup()
	})

	Context("invoking vector CLI commands that talk to the vector API", func() {
		It("should work", func() {
			Expect(framework.Deploy()).To(BeNil())
			out, _ := framework.RunCommand(constants.CollectorName, `curl`, `-sv`, `-m`, `5`, `--connect-timeout`, `3`, `http://127.0.0.1:8686/health`)
			Expect(out).To(ContainSubstring(`{"ok":true}`))
		})
	})
})
