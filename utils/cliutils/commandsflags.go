package cliutils

import (
	"github.com/codegangsta/cli"
	"github.com/jfrog/jfrog-client-go/utils/log"
	"sort"
	"strconv"
)

const (
	// Artifactory's Commands Keys
	Config                  = "config"
	Upload                  = "upload"
	Download                = "download"
	Move                    = "move"
	Copy                    = "copy"
	Delete                  = "delete"
	Properties              = "properties"
	Search                  = "search"
	BuildPublish            = "build-publish"
	BuildAppend             = "build-append"
	BuildScan               = "build-scan"
	BuildPromote            = "build-promote"
	BuildDistribute         = "build-distribute"
	BuildDiscard            = "build-discard"
	BuildAddDependencies    = "build-add-dependencies"
	BuildAddGit             = "build-add-git"
	GitLfsClean             = "git-lfs-clean"
	Mvn                     = "mvn"
	MvnConfig               = "mvn-config"
	Gradle                  = "gradle"
	GradleConfig            = "gradle-config"
	DockerPromote           = "docker-promote"
	DockerPull              = "docker-pull"
	DockerPush              = "docker-push"
	NpmConfig               = "npm-config"
	Npm                     = "npm"
	NpmPublish              = "npmPublish"
	NugetConfig             = "nuget-config"
	Nuget                   = "nuget"
	Dotnet                  = "dotnet"
	DotnetConfig            = "dotnet-config"
	Go                      = "go"
	GoConfig                = "go-config"
	GoPublish               = "go-publish"
	GoRecursivePublish      = "go-recursive-publish"
	PipInstall              = "pip-install"
	PipConfig               = "pip-config"
	Ping                    = "ping"
	Curl                    = "curl"
	ReleaseBundleCreate     = "release-bundle-create"
	ReleaseBundleUpdate     = "release-bundle-update"
	ReleaseBundleSign       = "release-bundle-sign"
	ReleaseBundleDistribute = "release-bundle-distribute"
	ReleaseBundleDelete     = "release-bundle-delete"
	TemplateConsumer        = "template-consumer"
	RepoDelete              = "repo-delete"
	ReplicationDelete       = "replication-delete"
	PermissionTargetDelete  = "permission-target-delete"
	AccessTokenCreate       = "access-token-create"
	// MC's Commands Keys
	McConfig       = "mc-config"
	LicenseAcquire = "license-acquire"
	LicenseDeploy  = "license-deploy"
	LicenseRelease = "license-release"
	JpdAdd         = "jpd-add"
	JpdDelete      = "jpd-delete"
	// XRay's Commands Keys
	OfflineUpdate = "offline-update"

	// *** Artifactory Commands' flags ***
	// Base flags
	url         = "url"
	distUrl     = "dist-url"
	user        = "user"
	password    = "password"
	apikey      = "apikey"
	accessToken = "access-token"
	serverId    = "server-id"

	// Deprecated flags
	deprecatedPrefix      = "deprecated-"
	deprecatedUrl         = deprecatedPrefix + url
	deprecatedUser        = deprecatedPrefix + user
	deprecatedPassword    = deprecatedPrefix + password
	deprecatedApikey      = deprecatedPrefix + apikey
	deprecatedAccessToken = deprecatedPrefix + accessToken

	// Ssh flags
	sshKeyPath    = "ssh-key-path"
	sshPassPhrase = "ssh-passphrase"

	// Client certification flags
	clientCertPath    = "client-cert-path"
	clientCertKeyPath = "client-cert-key-path"
	insecureTls       = "insecure-tls"

	// Sort & limit flags
	sortBy    = "sort-by"
	sortOrder = "sort-order"
	limit     = "limit"
	offset    = "offset"

	// Spec flags
	spec     = "spec"
	specVars = "spec-vars"

	// Build info flags
	buildName   = "build-name"
	buildNumber = "build-number"
	module      = "module"

	// Generic commands flags
	excludePatterns  = "exclude-patterns"
	exclusions       = "exclusions"
	recursive        = "recursive"
	flat             = "flat"
	build            = "build"
	regexpFlag       = "regexp"
	retries          = "retries"
	dryRun           = "dry-run"
	explode          = "explode"
	includeDirs      = "include-dirs"
	props            = "props"
	excludeProps     = "exclude-props"
	failNoOp         = "fail-no-op"
	threads          = "threads"
	syncDeletes      = "sync-deletes"
	quiet            = "quiet"
	bundle           = "bundle"
	archiveEntries   = "archive-entries"
	detailedSummary  = "detailed-summary"
	syncDeletesQuiet = syncDeletes + "-" + quiet

	// Config flags
	interactive   = "interactive"
	encPassword   = "enc-password"
	basicAuthOnly = "basic-auth-only"

	// Unique upload flags
	uploadPrefix          = "upload-"
	uploadExcludePatterns = uploadPrefix + excludePatterns
	uploadExclusions      = uploadPrefix + exclusions
	uploadRecursive       = uploadPrefix + recursive
	uploadFlat            = uploadPrefix + flat
	uploadRegexp          = uploadPrefix + regexpFlag
	uploadRetries         = uploadPrefix + retries
	uploadExplode         = uploadPrefix + explode
	uploadProps           = uploadPrefix + props
	uploadSyncDeletes     = uploadPrefix + syncDeletes
	deb                   = "deb"
	symlinks              = "symlinks"

	// Unique download flags
	downloadPrefix       = "download-"
	downloadRecursive    = downloadPrefix + recursive
	downloadFlat         = downloadPrefix + flat
	downloadRetries      = downloadPrefix + retries
	downloadExplode      = downloadPrefix + explode
	downloadProps        = downloadPrefix + props
	downloadExcludeProps = downloadPrefix + excludeProps
	downloadSyncDeletes  = downloadPrefix + syncDeletes
	minSplit             = "min-split"
	splitCount           = "split-count"
	validateSymlinks     = "validate-symlinks"

	// Unique move flags
	movePrefix       = "move-"
	moveRecursive    = movePrefix + recursive
	moveFlat         = movePrefix + flat
	moveProps        = movePrefix + props
	moveExcludeProps = movePrefix + excludeProps

	// Unique copy flags
	copyPrefix       = "copy-"
	copyRecursive    = copyPrefix + recursive
	copyFlat         = copyPrefix + flat
	copyProps        = copyPrefix + props
	copyExcludeProps = copyPrefix + excludeProps

	// Unique delete flags
	deletePrefix       = "delete-"
	deleteRecursive    = deletePrefix + recursive
	deleteProps        = deletePrefix + props
	deleteExcludeProps = deletePrefix + excludeProps
	deleteQuiet        = deletePrefix + quiet

	// Unique search flags
	searchPrefix       = "search-"
	searchRecursive    = searchPrefix + recursive
	searchProps        = searchPrefix + props
	searchExcludeProps = searchPrefix + excludeProps
	count              = "count"

	// Unique properties flags
	propertiesPrefix  = "props-"
	propsRecursive    = propertiesPrefix + recursive
	propsProps        = propertiesPrefix + props
	propsExcludeProps = propertiesPrefix + excludeProps

	// Unique build-publish flags
	buildPublishPrefix = "bp-"
	bpDryRun           = buildPublishPrefix + dryRun
	envInclude         = "env-include"
	envExclude         = "env-exclude"
	buildUrl           = "build-url"

	// Unique build-add-dependencies flags
	badPrefix    = "bad-"
	badDryRun    = badPrefix + dryRun
	badRecursive = badPrefix + recursive
	badRegexp    = badPrefix + regexpFlag

	// Unique build-add-git flags
	configFlag = "config"

	// Unique build-scan flags
	fail = "fail"

	// Unique build-promote flags
	buildPromotePrefix  = "bpr-"
	bprDryRun           = buildPromotePrefix + dryRun
	bprProps            = buildPromotePrefix + props
	status              = "status"
	comment             = "comment"
	sourceRepo          = "source-repo"
	includeDependencies = "include-dependencies"
	copyFlag            = "copy"

	async = "async"

	// Unique build-distribute flags
	buildDistributePrefix = "bd-"
	bdDryRun              = buildDistributePrefix + dryRun
	bdAsync               = buildDistributePrefix + async
	sourceRepos           = "source-repos"
	passphrase            = "passphrase"
	publish               = "publish"
	override              = "override"

	// Unique build-discard flags
	buildDiscardPrefix = "bdi-"
	bdiAsync           = buildDiscardPrefix + async
	maxDays            = "max-days"
	maxBuilds          = "max-builds"
	excludeBuilds      = "exclude-builds"
	deleteArtifacts    = "delete-artifacts"

	repo = "repo"

	// Unique git-lfs-clean flags
	glcPrefix = "glc-"
	glcDryRun = glcPrefix + dryRun
	glcQuiet  = glcPrefix + quiet
	glcRepo   = glcPrefix + repo
	refs      = "refs"

	// Build tool config flags
	global          = "global"
	serverIdResolve = "server-id-resolve"
	serverIdDeploy  = "server-id-deploy"
	repoResolve     = "repo-resolve"
	repoDeploy      = "repo-deploy"

	// Unique maven-config flags
	repoResolveReleases  = "repo-resolve-releases"
	repoResolveSnapshots = "repo-resolve-snapshots"
	repoDeployReleases   = "repo-deploy-releases"
	repoDeploySnapshots  = "repo-deploy-snapshots"

	// Unique gradle-config flags
	usesPlugin          = "uses-plugin"
	useWrapper          = "use-wrapper"
	deployMavenDesc     = "deploy-maven-desc"
	deployIvyDesc       = "deploy-ivy-desc"
	ivyDescPattern      = "ivy-desc-pattern"
	ivyArtifactsPattern = "ivy-artifacts-pattern"

	// Build tool flags
	deploymentThreads = "deployment-threads"
	skipLogin         = "skip-login"

	// Unique docker promote flags
	dockerPromotePrefix = "docker-promote-"
	targetDockerImage   = "target-docker-image"
	sourceTag           = "source-tag"
	targetTag           = "target-tag"
	dockerPromoteCopy   = dockerPromotePrefix + Copy

	// Unique npm flags
	npmPrefix  = "npm-"
	npmThreads = npmPrefix + threads
	npmArgs    = "npm-args"

	// Unique nuget flags
	nugetArgs    = "nuget-args"
	solutionRoot = "solution-root"

	// Unique go flags
	deps        = "deps"
	self        = "self"
	noRegistry  = "no-registry"
	publishDeps = "publish-deps"

	// Unique release-bundle flags
	releaseBundlePrefix = "rb-"
	rbDryRun            = releaseBundlePrefix + dryRun
	rbRepo              = releaseBundlePrefix + repo
	rbPassphrase        = releaseBundlePrefix + passphrase
	sign                = "sign"
	desc                = "desc"
	releaseNotesPath    = "release-notes-path"
	releaseNotesSyntax  = "release-notes-syntax"
	distRules           = "dist-rules"
	site                = "site"
	city                = "city"
	countryCodes        = "country-codes"
	sync                = "sync"
	maxWaitMinutes      = "max-wait-minutes"
	deleteFromDist      = "delete-from-dist"

	// Template user flags
	vars = "vars"

	// Unique access-token-create flags
	groups      = "groups"
	grantAdmin  = "grant-admin"
	expiry      = "expiry"
	refreshable = "refreshable"
	audience    = "audience"

	// *** Xray Commands' flags ***
	// Unique offline-update flags
	licenseId = "license-id"
	from      = "from"
	to        = "to"
	version   = "version"
	target    = "target"

	// *** Mission Control Commands' flags ***
	missionControlPrefix = "mc-"

	// Authentication flags
	mcUrl         = missionControlPrefix + url
	mcAccessToken = missionControlPrefix + accessToken

	// Unique config flags
	mcInteractive = missionControlPrefix + interactive

	// Unique license-deploy flags
	licenseCount = "license-count"
)

var flagsMap = map[string]cli.Flag{
	// Artifactory's commands Flags
	url: cli.StringFlag{
		Name:  url,
		Usage: "[Optional] Artifactory URL.` `",
	},
	deprecatedUrl: cli.StringFlag{
		Name:   url,
		Usage:  "[Deprecated] [Optional] Artifactory URL.` `",
		Hidden: true,
	},
	distUrl: cli.StringFlag{
		Name:  distUrl,
		Usage: "[Optional] Distribution URL.` `",
	},
	user: cli.StringFlag{
		Name:  user,
		Usage: "[Optional] Artifactory username.` `",
	},
	deprecatedUser: cli.StringFlag{
		Name:   user,
		Usage:  "[Deprecated] [Optional] Artifactory username.` `",
		Hidden: true,
	},
	password: cli.StringFlag{
		Name:  password,
		Usage: "[Optional] Artifactory password.` `",
	},
	deprecatedPassword: cli.StringFlag{
		Name:   password,
		Usage:  "[Deprecated] [Optional] Artifactory password.` `",
		Hidden: true,
	},
	apikey: cli.StringFlag{
		Name:  apikey,
		Usage: "[Optional] Artifactory API key.` `",
	},
	deprecatedApikey: cli.StringFlag{
		Name:   apikey,
		Usage:  "[Deprecated] [Optional] Artifactory API key.` `",
		Hidden: true,
	},
	accessToken: cli.StringFlag{
		Name:  accessToken,
		Usage: "[Optional] Artifactory access token.` `",
	},
	deprecatedAccessToken: cli.StringFlag{
		Name:   accessToken,
		Usage:  "[Deprecated] [Optional] Artifactory access token.` `",
		Hidden: true,
	},
	serverId: cli.StringFlag{
		Name:  serverId,
		Usage: "[Optional] Artifactory server ID configured using the config command.` `",
	},
	sshKeyPath: cli.StringFlag{
		Name:  sshKeyPath,
		Usage: "[Optional] SSH key file path.` `",
	},
	sshPassPhrase: cli.StringFlag{
		Name:  sshPassPhrase,
		Usage: "[Optional] SSH key passphrase.` `",
	},
	clientCertPath: cli.StringFlag{
		Name:  clientCertPath,
		Usage: "[Optional] Client certificate file in PEM format.` `",
	},
	clientCertKeyPath: cli.StringFlag{
		Name:  clientCertKeyPath,
		Usage: "[Optional] Private key file for the client certificate in PEM format.` `",
	},
	sortBy: cli.StringFlag{
		Name:  sortBy,
		Usage: "[Optional] A list of semicolon-separated fields to sort by. The fields must be part of the 'items' AQL domain. For more information, see https://www.jfrog.com/confluence/display/RTF/Artifactory+Query+Language#ArtifactoryQueryLanguage-EntitiesandFields` `",
	},
	sortOrder: cli.StringFlag{
		Name:  sortOrder,
		Usage: "[Default: asc] The order by which fields in the 'sort-by' option should be sorted. Accepts 'asc' or 'desc'.` `",
	},
	limit: cli.StringFlag{
		Name:  limit,
		Usage: "[Optional] The maximum number of items to fetch. Usually used with the 'sort-by' option.` `",
	},
	offset: cli.StringFlag{
		Name:  offset,
		Usage: "[Optional] The offset from which to fetch items (i.e. how many items should be skipped). Usually used with the 'sort-by' option.` `",
	},
	spec: cli.StringFlag{
		Name:  spec,
		Usage: "[Optional] Path to a File Spec.` `",
	},
	specVars: cli.StringFlag{
		Name:  specVars,
		Usage: "[Optional] List of variables in the form of \"key1=value1;key2=value2;...\" to be replaced in the File Spec. In the File Spec, the variables should be used as follows: ${key1}.` `",
	},
	buildName: cli.StringFlag{
		Name:  buildName,
		Usage: "[Optional] Providing this option will collect and record build info for this build name. Build number option is mandatory when this option is provided.` `",
	},
	buildNumber: cli.StringFlag{
		Name:  buildNumber,
		Usage: "[Optional] Providing this option will collect and record build info for this build number. Build name option is mandatory when this option is provided.` `",
	},
	module: cli.StringFlag{
		Name:  module,
		Usage: "[Optional] Optional module name for the build-info. Build name and number options are mandatory when this option is provided.` `",
	},
	excludePatterns: cli.StringFlag{
		Name:   excludePatterns,
		Usage:  "[Optional] Semicolon-separated list of exclude patterns. Exclude patterns may contain the * and the ? wildcards. Unlike the Source path, it must not include the repository name at the beginning of the path.` `",
		Hidden: true,
	},
	exclusions: cli.StringFlag{
		Name:  exclusions,
		Usage: "[Optional] Semicolon-separated list of exclusions. Exclusions can include the * and the ? wildcards.` `",
	},
	uploadExcludePatterns: cli.StringFlag{
		Name:   excludePatterns,
		Usage:  "[Optional] Semicolon-separated list of exclude patterns. Exclude patterns may contain the * and the ? wildcards or a regex pattern, according to the value of the 'regexp' option.` `",
		Hidden: true,
	},
	uploadExclusions: cli.StringFlag{
		Name:  exclusions,
		Usage: "[Optional] Semicolon-separated list of exclude patterns. Exclude patterns may contain the * and the ? wildcards or a regex pattern, according to the value of the 'regexp' option.` `",
	},
	build: cli.StringFlag{
		Name:  build,
		Usage: "[Optional] If specified, only artifacts of the specified build are matched. The property format is build-name/build-number. If you do not specify the build number, the artifacts are filtered by the latest build number.` `",
	},
	includeDirs: cli.BoolFlag{
		Name:  includeDirs,
		Usage: "[Default: false] Set to true if you'd like to also apply the source path pattern for directories and not just for files.` `",
	},
	failNoOp: cli.BoolFlag{
		Name:  failNoOp,
		Usage: "[Default: false] Set to true if you'd like the command to return exit code 2 in case of no files are affected.` `",
	},
	threads: cli.StringFlag{
		Name:  threads,
		Value: "",
		Usage: "[Default: 3] Number of working threads.` `",
	},
	insecureTls: cli.BoolFlag{
		Name:  insecureTls,
		Usage: "[Default: false] Set to true to skip TLS certificates verification.` `",
	},
	bundle: cli.StringFlag{
		Name:  bundle,
		Usage: "[Optional] If specified, only artifacts of the specified bundle are matched. The value format is bundle-name/bundle-version.` `",
	},
	archiveEntries: cli.StringFlag{
		Name:  archiveEntries,
		Usage: "[Optional] If specified, only archive artifacts containing entries matching this pattern are matched. You can use wildcards to specify multiple artifacts.` `",
	},
	detailedSummary: cli.BoolFlag{
		Name:  detailedSummary,
		Usage: "[Default: false] Set to true to include a list of the affected files in the command summary.` `",
	},
	interactive: cli.BoolTFlag{
		Name:  interactive,
		Usage: "[Default: true, unless $CI is true] Set to false if you do not want the config command to be interactive. If true, the --url option becomes optional.` `",
	},
	encPassword: cli.BoolTFlag{
		Name:  encPassword,
		Usage: "[Default: true] If set to false then the configured password will not be encrypted using Artifactory's encryption API.` `",
	},
	basicAuthOnly: cli.BoolFlag{
		Name: basicAuthOnly,
		Usage: "[Default: false] Set to true to disable replacing username and password/API key with automatically created access token that's refreshed hourly. " +
			"Username and password/API key will still be used with commands which use external tools or the JFrog Distribution service. " +
			"Can only be passed along with username and password/API key options.` `",
	},
	deb: cli.StringFlag{
		Name:  deb,
		Usage: "[Optional] Used for Debian packages in the form of distribution/component/architecture. If the value for distribution, component or architecture includes a slash, the slash should be escaped with a back-slash.` `",
	},
	uploadRecursive: cli.BoolTFlag{
		Name:  recursive,
		Usage: "[Default: true] Set to false if you do not wish to collect artifacts in sub-folders to be uploaded to Artifactory.` `",
	},
	uploadFlat: cli.BoolTFlag{
		Name:  flat,
		Usage: "[Default: true] If set to false, files are uploaded according to their file system hierarchy.` `",
	},
	uploadRegexp: cli.BoolFlag{
		Name:  regexpFlag,
		Usage: "[Default: false] Set to true to use a regular expression instead of wildcards expression to collect files to upload.` `",
	},
	uploadRetries: cli.StringFlag{
		Name:  retries,
		Usage: "[Default: " + strconv.Itoa(Retries) + "] Number of upload retries.` `",
	},
	dryRun: cli.BoolFlag{
		Name:  dryRun,
		Usage: "[Default: false] Set to true to disable communication with Artifactory.` `",
	},
	uploadExplode: cli.BoolFlag{
		Name:  explode,
		Usage: "[Default: false] Set to true to extract an archive after it is deployed to Artifactory.` `",
	},
	symlinks: cli.BoolFlag{
		Name:  symlinks,
		Usage: "[Default: false] Set to true to preserve symbolic links structure in Artifactory.` `",
	},
	uploadProps: cli.StringFlag{
		Name:  props,
		Usage: "[Optional] List of properties in the form of \"key1=value1;key2=value2,...\". Those properties will be attached to the uploaded artifacts.` `",
	},
	uploadSyncDeletes: cli.StringFlag{
		Name:  syncDeletes,
		Usage: "[Optional] Specific path in Artifactory, under which to sync artifacts after the upload. After the upload, this path will include only the artifacts uploaded during this upload operation. The other files under this path will be deleted.` `",
	},
	syncDeletesQuiet: cli.BoolFlag{
		Name:  quiet,
		Usage: "[Default: $CI] Set to true to skip the sync-deletes confirmation message.` `",
	},
	downloadRecursive: cli.BoolTFlag{
		Name:  recursive,
		Usage: "[Default: true] Set to false if you do not wish to include the download of artifacts inside sub-folders in Artifactory.` `",
	},
	downloadFlat: cli.BoolFlag{
		Name:  flat,
		Usage: "[Default: false] Set to true if you do not wish to have the Artifactory repository path structure created locally for your downloaded files.` `",
	},
	minSplit: cli.StringFlag{
		Name:  minSplit,
		Value: "",
		Usage: "[Default: " + strconv.Itoa(DownloadMinSplitKb) + "] Minimum file size in KB to split into ranges when downloading. Set to -1 for no splits.` `",
	},
	splitCount: cli.StringFlag{
		Name:  splitCount,
		Value: "",
		Usage: "[Default: " + strconv.Itoa(DownloadSplitCount) + "] Number of parts to split a file when downloading. Set to 0 for no splits.` `",
	},
	downloadRetries: cli.StringFlag{
		Name:  retries,
		Usage: "[Default: " + strconv.Itoa(Retries) + "] Number of download retries.` `",
	},
	downloadExplode: cli.BoolFlag{
		Name:  explode,
		Usage: "[Default: false] Set to true to extract an archive after it is downloaded from Artifactory.` `",
	},
	validateSymlinks: cli.BoolFlag{
		Name:  validateSymlinks,
		Usage: "[Default: false] Set to true to perform a checksum validation when downloading symbolic links.` `",
	},
	downloadProps: cli.StringFlag{
		Name:  props,
		Usage: "[Optional] List of properties in the form of \"key1=value1;key2=value2,...\". Only artifacts with these properties will be downloaded.` `",
	},
	downloadExcludeProps: cli.StringFlag{
		Name:  excludeProps,
		Usage: "[Optional] List of properties in the form of \"key1=value1;key2=value2,...\". Only artifacts without the specified properties will be downloaded.` `",
	},
	downloadSyncDeletes: cli.StringFlag{
		Name:  syncDeletes,
		Usage: "[Optional] Specific path in the local file system, under which to sync dependencies after the download. After the download, this path will include only the dependencies downloaded during this download operation. The other files under this path will be deleted.` `",
	},
	moveRecursive: cli.BoolTFlag{
		Name:  recursive,
		Usage: "[Default: true] Set to false if you do not wish to move artifacts inside sub-folders in Artifactory.` `",
	},
	moveFlat: cli.BoolFlag{
		Name:  flat,
		Usage: "[Default: false] If set to false, files are moved according to their file system hierarchy.` `",
	},
	moveProps: cli.StringFlag{
		Name:  props,
		Usage: "[Optional] List of properties in the form of \"key1=value1;key2=value2,...\". Only artifacts with these properties will be moved.` `",
	},
	moveExcludeProps: cli.StringFlag{
		Name:  excludeProps,
		Usage: "[Optional] List of properties in the form of \"key1=value1;key2=value2,...\". Only artifacts without the specified properties will be moved.` `",
	},
	copyRecursive: cli.BoolTFlag{
		Name:  recursive,
		Usage: "[Default: true] Set to false if you do not wish to copy artifacts inside sub-folders in Artifactory.` `",
	},
	copyFlat: cli.BoolFlag{
		Name:  flat,
		Usage: "[Default: false] If set to false, files are copied according to their file system hierarchy.` `",
	},
	copyProps: cli.StringFlag{
		Name:  props,
		Usage: "[Optional] List of properties in the form of \"key1=value1;key2=value2,...\". Only artifacts with these properties will be copied.` `",
	},
	copyExcludeProps: cli.StringFlag{
		Name:  excludeProps,
		Usage: "[Optional] List of properties in the form of \"key1=value1;key2=value2,...\". Only artifacts without the specified properties will be copied.` `",
	},
	deleteRecursive: cli.BoolTFlag{
		Name:  recursive,
		Usage: "[Default: true] Set to false if you do not wish to delete artifacts inside sub-folders in Artifactory.` `",
	},
	deleteProps: cli.StringFlag{
		Name:  props,
		Usage: "[Optional] List of properties in the form of \"key1=value1;key2=value2,...\". Only artifacts with these properties will be deleted.` `",
	},
	deleteExcludeProps: cli.StringFlag{
		Name:  excludeProps,
		Usage: "[Optional] List of properties in the form of \"key1=value1;key2=value2,...\". Only artifacts without the specified properties will be deleted.` `",
	},
	deleteQuiet: cli.BoolFlag{
		Name:  quiet,
		Usage: "[Default: $CI] Set to true to skip the delete confirmation message.` `",
	},
	searchRecursive: cli.BoolTFlag{
		Name:  recursive,
		Usage: "[Default: true] Set to false if you do not wish to search artifacts inside sub-folders in Artifactory.` `",
	},
	count: cli.BoolFlag{
		Name:  count,
		Usage: "[Optional] Set to true to display only the total of files or folders found.` `",
	},
	searchProps: cli.StringFlag{
		Name:  props,
		Usage: "[Optional] List of properties in the form of \"key1=value1;key2=value2,...\". Only artifacts with these properties will be returned.` `",
	},
	searchExcludeProps: cli.StringFlag{
		Name:  excludeProps,
		Usage: "[Optional] List of properties in the form of \"key1=value1;key2=value2,...\". Only artifacts without the specified properties will be returned` `",
	},
	propsRecursive: cli.BoolTFlag{
		Name:  recursive,
		Usage: "[Default: true] When false, artifacts inside sub-folders in Artifactory will not be affected.` `",
	},
	propsProps: cli.StringFlag{
		Name:  props,
		Usage: "[Optional] List of properties in the form of \"key1=value1;key2=value2,...\". Only artifacts with these properties are affected.` `",
	},
	propsExcludeProps: cli.StringFlag{
		Name:  excludeProps,
		Usage: "[Optional] List of properties in the form of \"key1=value1;key2=value2,...\". Only artifacts without the specified properties are affected` `",
	},
	buildUrl: cli.StringFlag{
		Name:  buildUrl,
		Usage: "[Optional] Can be used for setting the CI server build URL in the build-info.` `",
	},
	bpDryRun: cli.BoolFlag{
		Name:  dryRun,
		Usage: "[Default: false] Set to true to get a preview of the recorded build info, without publishing it to Artifactory.` `",
	},
	envInclude: cli.StringFlag{
		Name:  envInclude,
		Usage: "[Default: *] List of patterns in the form of \"value1;value2;...\" Only environment variables match those patterns will be included.` `",
	},
	envExclude: cli.StringFlag{
		Name:  envExclude,
		Usage: "[Default: *password*;*psw*;*secret*;*key*;*token*] List of case insensitive patterns in the form of \"value1;value2;...\". Environment variables match those patterns will be excluded.` `",
	},
	badRecursive: cli.BoolTFlag{
		Name:  recursive,
		Usage: "[Default: true] Set to false if you do not wish to collect artifacts in sub-folders to be added to the build info.` `",
	},
	badRegexp: cli.BoolFlag{
		Name:  regexpFlag,
		Usage: "[Default: false] Set to true to use a regular expression instead of wildcards expression to collect files to be added to the build info.` `",
	},
	badDryRun: cli.BoolFlag{
		Name:  dryRun,
		Usage: "[Default: false] Set to true to only get a summery of the dependencies that will be added to the build info.` `",
	},
	configFlag: cli.StringFlag{
		Name:  configFlag,
		Usage: "[Optional] Path to a configuration file.` `",
	},
	fail: cli.BoolTFlag{
		Name:  fail,
		Usage: "[Default: true] Set to false if you do not wish the command to return exit code 3, even if the 'Fail Build' rule is matched by Xray.` `",
	},
	status: cli.StringFlag{
		Name:  status,
		Usage: "[Optional] Build promotion status.` `",
	},
	comment: cli.StringFlag{
		Name:  comment,
		Usage: "[Optional] Build promotion comment.` `",
	},
	sourceRepo: cli.StringFlag{
		Name:  sourceRepo,
		Usage: "[Optional] Build promotion source repository.` `",
	},
	includeDependencies: cli.BoolFlag{
		Name:  includeDependencies,
		Usage: "[Default: false] If set to true, the build dependencies are also promoted.` `",
	},
	copyFlag: cli.BoolFlag{
		Name:  copyFlag,
		Usage: "[Default: false] If set true, the build artifacts and dependencies are copied to the target repository, otherwise they are moved.` `",
	},
	bprDryRun: cli.BoolFlag{
		Name:  dryRun,
		Usage: "[Default: false] If true, promotion is only simulated. The build is not promoted.` `",
	},
	bprProps: cli.StringFlag{
		Name:  props,
		Usage: "[Optional] List of properties in the form of \"key1=value1;key2=value2,...\". A list of properties to attach to the build artifacts.` `",
	},
	targetDockerImage: cli.StringFlag{
		Name:  "target-docker-image",
		Usage: "[Optional] Docker target image name.` `",
	},
	sourceTag: cli.StringFlag{
		Name:  "source-tag",
		Usage: "[Optional] The tag name to promote.` `",
	},
	targetTag: cli.StringFlag{
		Name:  "target-tag",
		Usage: "[Optional] The target tag to assign the image after promotion.` `",
	},
	dockerPromoteCopy: cli.BoolFlag{
		Name:  "copy",
		Usage: "[Default: false] If set true, the Docker image is copied to the target repository, otherwise it is moved.` `",
	},
	sourceRepos: cli.StringFlag{
		Name:  sourceRepos,
		Usage: "[Optional] List of local repositories in the form of \"repo1,repo2,...\" from which build artifacts should be deployed.` `",
	},
	passphrase: cli.StringFlag{
		Name:  passphrase,
		Usage: "[Optional] If specified, Artifactory will GPG sign the build deployed to Bintray and apply the specified passphrase.` `",
	},
	publish: cli.BoolTFlag{
		Name:  publish,
		Usage: "[Default: true] If true, builds are published when deployed to Bintray.` `",
	},
	override: cli.BoolFlag{
		Name:  override,
		Usage: "[Default: false] If true, Artifactory overwrites builds already existing in the target path in Bintray.` `",
	},
	bdAsync: cli.BoolFlag{
		Name:  async,
		Usage: "[Default: false] If true, the build will be distributed asynchronously.` `",
	},
	bdDryRun: cli.BoolFlag{
		Name:  dryRun,
		Usage: "[Default: false] If true, distribution is only simulated. No files are actually moved.` `",
	},
	maxDays: cli.StringFlag{
		Name:  maxDays,
		Usage: "[Optional] The maximum number of days to keep builds in Artifactory.` `",
	},
	maxBuilds: cli.StringFlag{
		Name:  maxBuilds,
		Usage: "[Optional] The maximum number of builds to store in Artifactory.` `",
	},
	excludeBuilds: cli.StringFlag{
		Name:  excludeBuilds,
		Usage: "[Optional] List of build numbers in the form of \"value1,value2,...\", that should not be removed from Artifactory.` `",
	},
	deleteArtifacts: cli.BoolFlag{
		Name:  deleteArtifacts,
		Usage: "[Default: false] If set to true, automatically removes build artifacts stored in Artifactory.` `",
	},
	bdiAsync: cli.BoolFlag{
		Name:  async,
		Usage: "[Default: false] If set to true, build discard will run asynchronously and will not wait for response.` `",
	},
	refs: cli.StringFlag{
		Name:  refs,
		Usage: "[Default: refs/remotes/*] List of Git references in the form of \"ref1,ref2,...\" which should be preserved.` `",
	},
	glcRepo: cli.StringFlag{
		Name:  repo,
		Usage: "[Optional] Local Git LFS repository which should be cleaned. If omitted, this is detected from the Git repository.` `",
	},
	glcDryRun: cli.BoolFlag{
		Name:  dryRun,
		Usage: "[Default: false] If true, cleanup is only simulated. No files are actually deleted.` `",
	},
	glcQuiet: cli.BoolFlag{
		Name:  quiet,
		Usage: "[Default: $CI] Set to true to skip the delete confirmation message.` `",
	},
	global: cli.BoolFlag{
		Name:  global,
		Usage: "[Default: false] Set to true if you'd like the configuration to be global (for all projects). Specific projects can override the global configuration.` `",
	},
	serverIdResolve: cli.StringFlag{
		Name:  serverIdResolve,
		Usage: "[Optional] Artifactory server ID for resolution. The server should configured using the 'jfrog rt c' command.` `",
	},
	serverIdDeploy: cli.StringFlag{
		Name:  serverIdDeploy,
		Usage: "[Optional] Artifactory server ID for deployment. The server should configured using the 'jfrog rt c' command.` `",
	},
	repoResolveReleases: cli.StringFlag{
		Name:  repoResolveReleases,
		Usage: "[Optional] Resolution repository for release dependencies.` `",
	},
	repoResolveSnapshots: cli.StringFlag{
		Name:  repoResolveSnapshots,
		Usage: "[Optional] Resolution repository for snapshot dependencies.` `",
	},
	repoDeployReleases: cli.StringFlag{
		Name:  repoDeployReleases,
		Usage: "[Optional] Deployment repository for release artifacts.` `",
	},
	repoDeploySnapshots: cli.StringFlag{
		Name:  repoDeploySnapshots,
		Usage: "[Optional] Deployment repository for snapshot artifacts.` `",
	},
	repoResolve: cli.StringFlag{
		Name:  repoResolve,
		Usage: "[Optional] Repository for dependencies resolution.` `",
	},
	repoDeploy: cli.StringFlag{
		Name:  repoDeploy,
		Usage: "[Optional] Repository for artifacts deployment.` `",
	},
	usesPlugin: cli.BoolFlag{
		Name:  usesPlugin,
		Usage: "[Default: false] Set to true if the Gradle Artifactory Plugin is already applied in the build script.` `",
	},
	useWrapper: cli.BoolFlag{
		Name:  useWrapper,
		Usage: "[Default: false] Set to true if you'd like to use the Gradle wrapper.` `",
	},
	deployMavenDesc: cli.BoolTFlag{
		Name:  deployMavenDesc,
		Usage: "[Default: true] Set to false if you do not wish to deploy Maven descriptors.` `",
	},
	deployIvyDesc: cli.BoolTFlag{
		Name:  deployIvyDesc,
		Usage: "[Default: true] Set to false if you do not wish to deploy Ivy descriptors.` `",
	},
	ivyDescPattern: cli.StringFlag{
		Name:  ivyDescPattern,
		Usage: "[Default: '[organization]/[module]/ivy-[revision].xml' Set the deployed Ivy descriptor pattern.` `",
	},
	ivyArtifactsPattern: cli.StringFlag{
		Name:  ivyArtifactsPattern,
		Usage: "[Default: '[organization]/[module]/[revision]/[artifact]-[revision](-[classifier]).[ext]' Set the deployed Ivy artifacts pattern.` `",
	},
	deploymentThreads: cli.StringFlag{
		Name:  threads,
		Value: "",
		Usage: "[Default: 3] Number of threads for uploading build artifacts.` `",
	},
	skipLogin: cli.BoolFlag{
		Name:  skipLogin,
		Usage: "[Default: false] Set to true if you'd like the command to skip performing docker login.` `",
	},
	npmArgs: cli.StringFlag{
		Name:   npmArgs,
		Usage:  "[Deprecated] [Optional] A list of npm arguments and options in the form of \"--arg1=value1 --arg2=value2\"` `",
		Hidden: true,
	},
	npmThreads: cli.StringFlag{
		Name:  threads,
		Value: "",
		Usage: "[Default: 3] Number of working threads for build-info collection.` `",
	},
	nugetArgs: cli.StringFlag{
		Name:   nugetArgs,
		Usage:  "[Deprecated] [Optional] A list of NuGet arguments and options in the form of \"arg1 arg2 arg3\"` `",
		Hidden: true,
	},
	solutionRoot: cli.StringFlag{
		Name:   solutionRoot,
		Usage:  "[Deprecated] [Default: .] Path to the root directory of the solution. If the directory includes more than one sln files, then the first argument passed in the --nuget-args option should be the name (not the path) of the sln file.` `",
		Hidden: true,
	},
	deps: cli.StringFlag{
		Name:  deps,
		Value: "",
		Usage: "[Optional] List of project dependencies in the form of \"dep1-name:version,dep2-name:version...\" to be published to Artifactory. Use \"ALL\" to publish all dependencies.` `",
	},
	self: cli.BoolTFlag{
		Name:  self,
		Usage: "[Default: true] Set false to skip publishing the project package zip file to Artifactory..` `",
	},
	noRegistry: cli.BoolFlag{
		Name:   noRegistry,
		Usage:  "[Deprecated] [Default: false] Set to true if you don't want to use Artifactory as your proxy` `",
		Hidden: true,
	},
	publishDeps: cli.BoolFlag{
		Name:   publishDeps,
		Usage:  "[Deprecated] [Default: false] Set to true if you wish to publish missing dependencies to Artifactory` `",
		Hidden: true,
	},
	rbDryRun: cli.BoolFlag{
		Name:  dryRun,
		Usage: "[Default: false] Set to true to disable communication with JFrog Distribution.` `",
	},
	sign: cli.BoolFlag{
		Name:  sign,
		Usage: "[Default: false] If set to true, automatically signs the release bundle version.` `",
	},
	desc: cli.StringFlag{
		Name:  desc,
		Usage: "[Optional] Description of the release bundle.` `",
	},
	releaseNotesPath: cli.StringFlag{
		Name:  releaseNotesPath,
		Usage: "[Optional] Path to a file describes the release notes for the release bundle version.` `",
	},
	releaseNotesSyntax: cli.StringFlag{
		Name:  "release-notes-syntax",
		Usage: "[Default: plain_text] The syntax for the release notes. Can be one of 'markdown', 'asciidoc', or 'plain_text` `",
	},
	rbPassphrase: cli.StringFlag{
		Name:  passphrase,
		Usage: "[Optional] The passphrase for the signing key. ` `",
	},
	rbRepo: cli.StringFlag{
		Name:  repo,
		Usage: "[Optional] A repository name at source Artifactory to store release bundle artifacts in. If not provided, Artifactory will use the default one.` `",
	},
	distRules: cli.StringFlag{
		Name:  distRules,
		Usage: "Path to distribution rules.` `",
	},
	site: cli.StringFlag{
		Name:  site,
		Usage: "[Default: '*'] Wildcard filter for site name. ` `",
	},
	city: cli.StringFlag{
		Name:  city,
		Usage: "[Default: '*'] Wildcard filter for site city name. ` `",
	},
	countryCodes: cli.StringFlag{
		Name:  countryCodes,
		Usage: "[Default: '*'] Semicolon-separated list of wildcard filters for site country codes. ` `",
	},
	sync: cli.BoolFlag{
		Name:  sync,
		Usage: "[Default: false] Set to true to enable sync distribution (the command execution will end when the distribution process ends).` `",
	},
	maxWaitMinutes: cli.StringFlag{
		Name:  maxWaitMinutes,
		Usage: "[Default: 60] Max minutes to wait for sync distribution. ` `",
	},
	deleteFromDist: cli.BoolFlag{
		Name:  deleteFromDist,
		Usage: "[Default: false] Set to true to delete release bundle version in JFrog Distribution itself after deletion is complete in the specified Edge node/s.` `",
	},
	vars: cli.StringFlag{
		Name:  vars,
		Usage: "[Optional] List of variables in the form of \"key1=value1;key2=value2;...\" to be replaced in the template. In the template, the variables should be used as follows: ${key1}.` `",
	},
	groups: cli.StringFlag{
		Name: groups,
		Usage: "[Default: *] A list of comma-separated groups for the access token to be associated with. " +
			"Specify * to indicate that this is a 'user-scoped token', i.e., the token provides the same access privileges that the current subject has, and is therefore evaluated dynamically. " +
			"A non-admin user can only provide a scope that is a subset of the groups to which he belongs` `",
	},
	grantAdmin: cli.BoolFlag{
		Name:  grantAdmin,
		Usage: "[Default: false] Set to true to provides admin privileges to the access token. This is only available for administrators.` `",
	},
	expiry: cli.StringFlag{
		Name:  expiry,
		Usage: "[Default: " + strconv.Itoa(TokenExpiry) + "] The time in seconds for which the token will be valid. To specify a token that never expires, set to zero. Non-admin can only set a value that is equal to or less than the default 3600.` `",
	},
	refreshable: cli.BoolFlag{
		Name:  refreshable,
		Usage: "[Default: false] Set to true if you'd like the the token to be refreshable. A refresh token will also be returned in order to be used to generate a new token once it expires.` `",
	},
	audience: cli.StringFlag{
		Name:  audience,
		Usage: "[Optional] A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs, as obtained by the 'jfrog rt curl api/system/service_id' command.` `",
	},
	// Xray's commands Flags
	licenseId: cli.StringFlag{
		Name:  licenseId,
		Usage: "[Mandatory] Xray license ID` `",
	},
	from: cli.StringFlag{
		Name:  from,
		Usage: "[Optional] From update date in YYYY-MM-DD format.` `",
	},
	to: cli.StringFlag{
		Name:  to,
		Usage: "[Optional] To update date in YYYY-MM-DD format.` `",
	},
	version: cli.StringFlag{
		Name:  version,
		Usage: "[Optional] Xray API version.` `",
	},
	target: cli.StringFlag{
		Name:  target,
		Usage: "[Default: ./] Path for downloaded update files.` `",
	},
	// Mission Control's commands Flags
	mcUrl: cli.StringFlag{
		Name:  url,
		Usage: "[Optional] Mission Control URL.` `",
	},
	mcAccessToken: cli.StringFlag{
		Name:  accessToken,
		Usage: "[Optional] Mission Control Admin token.` `",
	},
	mcInteractive: cli.BoolTFlag{
		Name:  interactive,
		Usage: "[Default: true] Set to false if you do not want the config command to be interactive. If true, the other command options become optional.",
	},
	licenseCount: cli.StringFlag{
		Name:  licenseCount,
		Value: "",
		Usage: "[Default: " + strconv.Itoa(DefaultLicenseCount) + "] The number of licenses to deploy. Minimum value is 1.` `",
	},
}

var commandFlags = map[string][]string{
	Config: {
		interactive, encPassword, url, distUrl, user, password, apikey, accessToken, sshKeyPath, clientCertPath,
		clientCertKeyPath, basicAuthOnly, insecureTls,
	},
	Upload: {
		url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath, serverId, clientCertPath,
		clientCertKeyPath, spec, specVars, buildName, buildNumber, module, uploadExcludePatterns, uploadExclusions, deb,
		uploadRecursive, uploadFlat, uploadRegexp, uploadRetries, dryRun, uploadExplode, symlinks, includeDirs,
		uploadProps, failNoOp, threads, uploadSyncDeletes, syncDeletesQuiet, insecureTls,
	},
	Download: {
		url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath, serverId, clientCertPath,
		clientCertKeyPath, spec, specVars, buildName, buildNumber, module, excludePatterns, exclusions, sortBy,
		sortOrder, limit, offset, downloadRecursive, downloadFlat, build, minSplit, splitCount, downloadRetries, dryRun,
		downloadExplode, validateSymlinks, bundle, includeDirs, downloadProps, downloadExcludeProps, failNoOp, threads,
		archiveEntries, downloadSyncDeletes, syncDeletesQuiet, insecureTls, detailedSummary,
	},
	Move: {
		url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath, serverId, clientCertPath,
		clientCertKeyPath, spec, specVars, excludePatterns, exclusions, sortBy, sortOrder, limit, offset, moveRecursive,
		moveFlat, dryRun, build, moveProps, moveExcludeProps, failNoOp, archiveEntries, insecureTls,
	},
	Copy: {
		url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath, serverId, clientCertPath,
		clientCertKeyPath, spec, specVars, excludePatterns, exclusions, sortBy, sortOrder, limit, offset, copyRecursive,
		copyFlat, dryRun, build, bundle, copyProps, copyExcludeProps, failNoOp, archiveEntries, insecureTls,
	},
	Delete: {
		url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath, serverId, clientCertPath,
		clientCertKeyPath, spec, specVars, excludePatterns, exclusions, sortBy, sortOrder, limit, offset,
		deleteRecursive, dryRun, build, deleteQuiet, deleteProps, deleteExcludeProps, failNoOp, threads, archiveEntries,
		insecureTls,
	},
	Search: {
		url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath, serverId, clientCertPath,
		clientCertKeyPath, spec, specVars, excludePatterns, exclusions, sortBy, sortOrder, limit, offset,
		searchRecursive, build, count, bundle, includeDirs, searchProps, searchExcludeProps, failNoOp, archiveEntries,
		insecureTls,
	},
	Properties: {
		url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath, serverId, clientCertPath,
		clientCertKeyPath, spec, specVars, excludePatterns, exclusions, sortBy, sortOrder, limit, offset,
		propsRecursive, build, bundle, includeDirs, failNoOp, threads, archiveEntries, propsProps, propsExcludeProps,
		insecureTls,
	},
	BuildPublish: {
		url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath, serverId, buildUrl, bpDryRun,
		envInclude, envExclude, insecureTls,
	},
	BuildAppend: {
		url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath, serverId, buildUrl, bpDryRun,
		envInclude, envExclude, insecureTls,
	},
	BuildAddDependencies: {
		spec, specVars, uploadExcludePatterns, uploadExclusions, badRecursive, badRegexp, badDryRun,
	},
	BuildAddGit: {
		configFlag, serverId,
	},
	BuildScan: {
		url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath, serverId, fail, insecureTls,
	},
	BuildPromote: {
		url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath, serverId, status, comment,
		sourceRepo, includeDependencies, copyFlag, bprDryRun, bprProps, insecureTls,
	},
	BuildDistribute: {
		url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath, serverId, sourceRepos, passphrase,
		publish, override, bdAsync, bdDryRun, insecureTls,
	},
	BuildDiscard: {
		url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath, serverId, maxDays, maxBuilds,
		excludeBuilds, deleteArtifacts, bdiAsync, insecureTls,
	},
	GitLfsClean: {
		url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath, serverId, refs, glcRepo, glcDryRun,
		glcQuiet, insecureTls,
	},
	MvnConfig: {
		global, serverIdResolve, serverIdDeploy, repoResolveReleases, repoResolveSnapshots, repoDeployReleases, repoDeploySnapshots,
	},
	GradleConfig: {
		global, serverIdResolve, serverIdDeploy, repoResolve, repoDeploy, usesPlugin, useWrapper, deployMavenDesc,
		deployIvyDesc, ivyDescPattern, ivyArtifactsPattern,
	},
	Mvn: {
		buildName, buildNumber, deploymentThreads, insecureTls,
	},
	Gradle: {
		buildName, buildNumber, deploymentThreads,
	},
	DockerPromote: {
		targetDockerImage, sourceTag, targetTag, dockerPromoteCopy, url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath,
		serverId,
	},
	DockerPush: {
		buildName, buildNumber, module, url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath,
		serverId, skipLogin, threads,
	},
	DockerPull: {
		buildName, buildNumber, module, url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath,
		serverId, skipLogin,
	},
	NpmConfig: {
		global, serverIdResolve, serverIdDeploy, repoResolve, repoDeploy,
	},
	Npm: {
		npmArgs, deprecatedUrl, deprecatedUser, deprecatedPassword, deprecatedApikey, deprecatedAccessToken, buildName,
		buildNumber, module, npmThreads,
	},
	NpmPublish: {
		npmArgs, deprecatedUrl, deprecatedUser, deprecatedPassword, deprecatedApikey, deprecatedAccessToken, buildName,
		buildNumber, module,
	},
	NugetConfig: {
		global, serverIdResolve, repoResolve,
	},
	Nuget: {
		nugetArgs, solutionRoot, deprecatedUrl, deprecatedUser, deprecatedPassword, deprecatedApikey,
		deprecatedAccessToken, buildName, buildNumber, module,
	},
	DotnetConfig: {
		global, serverIdResolve, repoResolve,
	},
	Dotnet: {
		buildName, buildNumber, module,
	},
	GoConfig: {
		global, serverIdResolve, serverIdDeploy, repoResolve, repoDeploy,
	},
	GoPublish: {
		deps, self, url, user, password, apikey, accessToken, serverId, buildName, buildNumber, module,
	},
	Go: {
		noRegistry, publishDeps, deprecatedUrl, deprecatedUser, deprecatedPassword, deprecatedApikey,
		deprecatedAccessToken, buildName, buildNumber, module,
	},
	GoRecursivePublish: {
		url, user, password, apikey, accessToken, serverId,
	},
	Ping: {
		url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath, serverId, clientCertPath,
		clientCertKeyPath, insecureTls,
	},
	Curl: {
		serverId,
	},
	PipConfig: {
		global, serverIdResolve, repoResolve,
	},
	PipInstall: {
		buildName, buildNumber, module,
	},
	ReleaseBundleCreate: {
		url, distUrl, user, password, apikey, accessToken, sshKeyPath, sshPassPhrase, serverId, spec, specVars,
		rbDryRun, sign, desc, exclusions, releaseNotesPath, releaseNotesSyntax, rbPassphrase, rbRepo, insecureTls,
	},
	ReleaseBundleUpdate: {
		url, distUrl, user, password, apikey, accessToken, sshKeyPath, sshPassPhrase, serverId, spec, specVars,
		rbDryRun, sign, desc, exclusions, releaseNotesPath, releaseNotesSyntax, rbPassphrase, rbRepo, insecureTls,
	},
	ReleaseBundleSign: {
		url, distUrl, user, password, apikey, accessToken, sshKeyPath, sshPassPhrase, serverId, rbPassphrase, rbRepo,
		insecureTls,
	},
	ReleaseBundleDistribute: {
		url, distUrl, user, password, apikey, accessToken, sshKeyPath, sshPassPhrase, serverId, rbDryRun, distRules,
		site, city, countryCodes, sync, maxWaitMinutes, insecureTls,
	},
	ReleaseBundleDelete: {
		url, distUrl, user, password, apikey, accessToken, sshKeyPath, sshPassPhrase, serverId, rbDryRun, distRules,
		site, city, countryCodes, sync, maxWaitMinutes, insecureTls, deleteFromDist, deleteQuiet,
	},
	TemplateConsumer: {
		url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath, serverId, clientCertPath,
		clientCertKeyPath, vars,
	},
	RepoDelete: {
		url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath, serverId, clientCertPath,
		clientCertKeyPath, deleteQuiet,
	},
	ReplicationDelete: {
		url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath, serverId, clientCertPath,
		clientCertKeyPath, deleteQuiet,
	},
	PermissionTargetDelete: {
		url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath, serverId, clientCertPath,
		clientCertKeyPath, deleteQuiet,
	},
	AccessTokenCreate: {
		url, user, password, apikey, accessToken, sshPassPhrase, sshKeyPath, serverId, clientCertPath,
		clientCertKeyPath, groups, grantAdmin, expiry, refreshable, audience,
	},
	// Xray's commands
	OfflineUpdate: {
		licenseId, from, to, version, target,
	},
	// Mission Control's commands
	McConfig: {
		mcUrl, mcAccessToken, mcInteractive,
	},
	LicenseAcquire: {
		mcUrl, mcAccessToken,
	},
	LicenseDeploy: {
		mcUrl, mcAccessToken, licenseCount,
	},
	LicenseRelease: {
		mcUrl, mcAccessToken,
	},
	JpdAdd: {
		mcUrl, mcAccessToken,
	},
	JpdDelete: {
		mcUrl, mcAccessToken,
	},
}

func GetCommandFlags(cmd string) []cli.Flag {
	flagList, ok := commandFlags[cmd]
	if !ok {
		log.Error("The command \"", cmd, "\" does not found in commands flags map.")
		return nil
	}
	return buildAndSortFlags(flagList)
}

func buildAndSortFlags(keys []string) (flags []cli.Flag) {
	for _, flag := range keys {
		flags = append(flags, flagsMap[flag])
	}
	sort.Slice(flags, func(i, j int) bool { return flags[i].GetName() < flags[j].GetName() })
	return
}

// This function is used for mvn and gradle command validation
func GetBasicBuildToolsFlags() (flags []cli.Flag) {
	basicBuildToolsFlags := []string{url, distUrl, user, password, apikey, accessToken, serverId}
	return buildAndSortFlags(basicBuildToolsFlags)
}

var deprecatedFlags = []string{deprecatedUrl, deprecatedUser, deprecatedPassword, deprecatedApikey, deprecatedAccessToken}

// This function is used for legacy (deprecated) nuget command validation
func GetLegacyNugetFlags() (flags []cli.Flag) {
	legacyNugetFlags := []string{nugetArgs, solutionRoot}
	legacyNugetFlags = append(legacyNugetFlags, deprecatedFlags...)
	return buildAndSortFlags(legacyNugetFlags)
}

// This function is used for legacy (deprecated) npm command validation
func GetLegacyNpmFlags() (flags []cli.Flag) {
	legacyNpmFlags := append(deprecatedFlags, npmArgs)
	return buildAndSortFlags(legacyNpmFlags)
}

// This function is used for legacy (deprecated) go command validation
func GetLegacyGoFlags() (flags []cli.Flag) {
	legacyGoFlags := []string{noRegistry, publishDeps}
	legacyGoFlags = append(legacyGoFlags, deprecatedFlags...)
	return buildAndSortFlags(legacyGoFlags)
}
