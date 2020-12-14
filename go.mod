module github.com/jfrog/jfrog-cli

go 1.14

require (
	github.com/buger/jsonparser v0.0.0-20180910192245-6acdf747ae99
	github.com/codegangsta/cli v1.20.0
	github.com/jfrog/gocmd v0.1.17
	github.com/jfrog/gofrog v1.0.6
	github.com/jfrog/jfrog-cli-core v1.1.2
	github.com/jfrog/jfrog-client-go v0.16.0
	github.com/mholt/archiver v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1
	github.com/vbauerster/mpb/v4 v4.7.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/jfrog/jfrog-client-go => github.com/jfrog/jfrog-client-go v0.16.1-0.20201214134149-9fbfa31cf8b1

replace github.com/jfrog/jfrog-cli-core => github.com/jfrog/jfrog-cli-core v1.1.3-0.20201214130915-fb405dd53615

// replace github.com/jfrog/gocmd => github.com/jfrog/gocmd master

// replace github.com/jfrog/gofrog => github.com/jfrog/gofrog v1.0.6
