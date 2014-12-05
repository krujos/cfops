package backup

import (
	"io/ioutil"
	"os"
	"path"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Backup", func() {
	var (
		dir     string
		context *BackupContext
	)

	BeforeEach(func() {
		dir, _ = ioutil.TempDir("", "cfops-backup")
		context = New("localhost", "admin", "admin", "admin", path.Join(dir, "backup"))
	})

	AfterEach(func() {
		//os.RemoveAll(dir)
	})

	Describe("Prepare Filesystem", func() {
		Context("With an empty target", func() {
			It("should create the parent directory", func() {
				Ω(context.Target).NotTo(BeEquivalentTo(""))
				Ω(FileExists(context.Target)).To(BeFalse())
				context.initPaths()
				context.prepareFilesystem()
				Ω(FileExists(context.Target)).To(BeTrue())
			})
		})
	})
})

func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}
