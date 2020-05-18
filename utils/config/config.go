package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/jfrog/jfrog-cli/utils/cliutils"
	artifactoryAuth "github.com/jfrog/jfrog-client-go/artifactory/auth"
	"github.com/jfrog/jfrog-client-go/auth"
	distributionAuth "github.com/jfrog/jfrog-client-go/distribution/auth"
	"github.com/jfrog/jfrog-client-go/utils"
	"github.com/jfrog/jfrog-client-go/utils/errorutils"
	"github.com/jfrog/jfrog-client-go/utils/io/fileutils"
	"github.com/jfrog/jfrog-client-go/utils/log"
	"io/ioutil"
	"os"
	"path/filepath"
)

// This is the default server id. It is used when adding a server config without providing a server ID
const (
	DefaultServerId   = "Default-Server"
	JfrogConfigFile   = "jfrog-cli.conf"
	JfrogDependencies = "dependencies"
)

func IsArtifactoryConfExists() (bool, error) {
	conf, err := readConf()
	if err != nil {
		return false, err
	}
	return conf.Artifactory != nil && len(conf.Artifactory) > 0, nil
}

func IsMissionControlConfExists() (bool, error) {
	conf, err := readConf()
	if err != nil {
		return false, err
	}
	return conf.MissionControl != nil && *conf.MissionControl != MissionControlDetails{}, nil
}

func IsBintrayConfExists() (bool, error) {
	conf, err := readConf()
	if err != nil {
		return false, err
	}
	return conf.Bintray != nil, nil
}

func GetArtifactorySpecificConfig(serverId string) (*ArtifactoryDetails, error) {
	conf, err := readConf()
	if err != nil {
		return nil, err
	}
	details := conf.Artifactory
	if details == nil || len(details) == 0 {
		return new(ArtifactoryDetails), nil
	}
	if len(serverId) == 0 {
		details, err := GetDefaultConfiguredArtifactoryConf(details)
		return details, errorutils.CheckError(err)
	}
	return getArtifactoryConfByServerId(serverId, details)
}

// Returns the default server configuration or error if not found.
// Caller should perform the check error if required.
func GetDefaultConfiguredArtifactoryConf(configs []*ArtifactoryDetails) (*ArtifactoryDetails, error) {
	if len(configs) == 0 {
		details := new(ArtifactoryDetails)
		details.IsDefault = true
		return details, nil
	}
	for _, conf := range configs {
		if conf.IsDefault == true {
			return conf, nil
		}
	}
	return nil, errors.New("Couldn't find default server.")
}

// Returns default artifactory conf. Returns nil if default server doesn't exists.
func GetDefaultArtifactoryConf() (*ArtifactoryDetails, error) {
	configurations, err := GetAllArtifactoryConfigs()
	if err != nil {
		return nil, err
	}

	if len(configurations) == 0 {
		log.Debug("No servers were configured.")
		return nil, err
	}

	return GetDefaultConfiguredArtifactoryConf(configurations)
}

// Returns the configured server or error if the server id not found
func GetArtifactoryConf(serverId string) (*ArtifactoryDetails, error) {
	configs, err := GetAllArtifactoryConfigs()
	if err != nil {
		return nil, err
	}
	return getArtifactoryConfByServerId(serverId, configs)
}

// Returns the configured server or error if the server id not found
func getArtifactoryConfByServerId(serverId string, configs []*ArtifactoryDetails) (*ArtifactoryDetails, error) {
	for _, conf := range configs {
		if conf.ServerId == serverId {
			return conf, nil
		}
	}
	return nil, errorutils.CheckError(errors.New(fmt.Sprintf("Server ID '%s' does not exist.", serverId)))
}

func GetAndRemoveConfiguration(serverName string, configs []*ArtifactoryDetails) (*ArtifactoryDetails, []*ArtifactoryDetails) {
	for i, conf := range configs {
		if conf.ServerId == serverName {
			configs = append(configs[:i], configs[i+1:]...)
			return conf, configs
		}
	}
	return nil, configs
}

func GetAllArtifactoryConfigs() ([]*ArtifactoryDetails, error) {
	conf, err := readConf()
	if err != nil {
		return nil, err
	}
	details := conf.Artifactory
	if details == nil {
		return make([]*ArtifactoryDetails, 0), nil
	}
	return details, nil
}

func ReadMissionControlConf() (*MissionControlDetails, error) {
	conf, err := readConf()
	if err != nil {
		return nil, err
	}
	details := conf.MissionControl
	if details == nil {
		return new(MissionControlDetails), nil
	}
	return details, nil
}

func ReadBintrayConf() (*BintrayDetails, error) {
	conf, err := readConf()
	if err != nil {
		return nil, err
	}
	details := conf.Bintray
	if details == nil {
		return new(BintrayDetails), nil
	}
	return details, nil
}

func SaveArtifactoryConf(details []*ArtifactoryDetails) error {
	conf, err := readConf()
	if err != nil {
		return err
	}
	conf.Artifactory = details
	return saveConfig(conf)
}

func SaveMissionControlConf(details *MissionControlDetails) error {
	conf, err := readConf()
	if err != nil {
		return err
	}
	conf.MissionControl = details
	return saveConfig(conf)
}

func SaveBintrayConf(details *BintrayDetails) error {
	config, err := readConf()
	if err != nil {
		return err
	}
	config.Bintray = details
	return saveConfig(config)
}

func saveConfig(config *ConfigV1) error {
	config.Version = cliutils.GetConfigVersion()
	b, err := json.Marshal(&config)
	if err != nil {
		return errorutils.CheckError(err)
	}
	var content bytes.Buffer
	err = json.Indent(&content, b, "", "  ")
	if err != nil {
		return errorutils.CheckError(err)
	}
	path, err := getConfFilePath()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, []byte(content.String()), 0600)
	if err != nil {
		return errorutils.CheckError(err)
	}

	return nil
}

func readConf() (*ConfigV1, error) {
	confFilePath, err := getConfFilePath()
	if err != nil {
		return nil, err
	}
	config := new(ConfigV1)
	exists, err := fileutils.IsFileExists(confFilePath, false)
	if err != nil {
		return nil, err
	}
	if !exists {
		return config, nil
	}
	content, err := fileutils.ReadFile(confFilePath)
	if err != nil {
		return nil, err
	}
	if len(content) == 0 {
		return new(ConfigV1), nil
	}
	content, err = convertIfNecessary(content)
	err = json.Unmarshal(content, &config)
	return config, err
}

// The configuration schema can change between versions, therefore we need to convert old versions to the new schema.
func convertIfNecessary(content []byte) ([]byte, error) {
	version, err := jsonparser.GetString(content, "Version")
	if err != nil {
		if err.Error() == "Key path not found" {
			version = "0"
		} else {
			return nil, errorutils.CheckError(err)
		}
	}
	switch version {
	case "0":
		result := new(ConfigV1)
		configV0 := new(ConfigV0)
		err = json.Unmarshal(content, &configV0)
		if errorutils.CheckError(err) != nil {
			return nil, err
		}
		result = configV0.Convert()
		err = saveConfig(result)
		content, err = json.Marshal(&result)
	}
	return content, err
}

func GetJfrogDependenciesPath() (string, error) {
	dependenciesDir := os.Getenv(cliutils.DependenciesDir)
	if dependenciesDir != "" {
		return utils.AddTrailingSlashIfNeeded(dependenciesDir), nil
	}
	jfrogHome, err := cliutils.GetJfrogHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(jfrogHome, JfrogDependencies), nil
}

func getConfFilePath() (string, error) {
	confPath, err := cliutils.GetJfrogHomeDir()
	if err != nil {
		return "", err
	}
	os.MkdirAll(confPath, 0777)
	return filepath.Join(confPath, JfrogConfigFile), nil
}

type ConfigV1 struct {
	Artifactory    []*ArtifactoryDetails  `json:"artifactory"`
	Bintray        *BintrayDetails        `json:"bintray,omitempty"`
	MissionControl *MissionControlDetails `json:"MissionControl,omitempty"`
	Version        string                 `json:"Version,omitempty"`
}

type ConfigV0 struct {
	Artifactory    *ArtifactoryDetails    `json:"artifactory,omitempty"`
	Bintray        *BintrayDetails        `json:"bintray,omitempty"`
	MissionControl *MissionControlDetails `json:"MissionControl,omitempty"`
}

func (o *ConfigV0) Convert() *ConfigV1 {
	config := new(ConfigV1)
	config.Bintray = o.Bintray
	config.MissionControl = o.MissionControl
	if o.Artifactory != nil {
		o.Artifactory.IsDefault = true
		o.Artifactory.ServerId = DefaultServerId
		config.Artifactory = []*ArtifactoryDetails{o.Artifactory}
	}
	return config
}

type ArtifactoryDetails struct {
	Url                  string `json:"url,omitempty"`
	SshUrl               string `json:"-"`
	DistributionUrl      string `json:"distributionUrl,omitempty"`
	User                 string `json:"user,omitempty"`
	Password             string `json:"password,omitempty"`
	SshKeyPath           string `json:"sshKeyPath,omitempty"`
	SshPassphrase        string `json:"SshPassphrase,omitempty"`
	AccessToken          string `json:"accessToken,omitempty"`
	RefreshToken         string `json:"refreshToken,omitempty"`
	TokenRefreshInterval int    `json:"tokenRefreshInterval,omitempty"`
	ClientCertPath       string `json:"clientCertPath,omitempty"`
	ClientCertKeyPath    string `json:"clientCertKeyPath,omitempty"`
	ServerId             string `json:"serverId,omitempty"`
	IsDefault            bool   `json:"isDefault,omitempty"`
	InsecureTls          bool   `json:"-"`
	// Deprecated, use password option instead.
	ApiKey string `json:"apiKey,omitempty"`
}

type BintrayDetails struct {
	ApiUrl            string `json:"-"`
	DownloadServerUrl string `json:"-"`
	User              string `json:"user,omitempty"`
	Key               string `json:"key,omitempty"`
	DefPackageLicense string `json:"defPackageLicense,omitempty"`
}

type MissionControlDetails struct {
	Url         string `json:"url,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
}

func (artifactoryDetails *ArtifactoryDetails) IsEmpty() bool {
	return len(artifactoryDetails.Url) == 0
}

func (artifactoryDetails *ArtifactoryDetails) SetApiKey(apiKey string) {
	artifactoryDetails.ApiKey = apiKey
}

func (artifactoryDetails *ArtifactoryDetails) SetUser(username string) {
	artifactoryDetails.User = username
}

func (artifactoryDetails *ArtifactoryDetails) SetPassword(password string) {
	artifactoryDetails.Password = password
}

func (artifactoryDetails *ArtifactoryDetails) SetAccessToken(accessToken string) {
	artifactoryDetails.AccessToken = accessToken
}

func (artifactoryDetails *ArtifactoryDetails) SetRefreshToken(refreshToken string) {
	artifactoryDetails.RefreshToken = refreshToken
}

func (artifactoryDetails *ArtifactoryDetails) SetClientCertPath(certificatePath string) {
	artifactoryDetails.ClientCertPath = certificatePath
}

func (artifactoryDetails *ArtifactoryDetails) SetClientCertKeyPath(certificatePath string) {
	artifactoryDetails.ClientCertKeyPath = certificatePath
}

func (artifactoryDetails *ArtifactoryDetails) GetApiKey() string {
	return artifactoryDetails.ApiKey
}

func (artifactoryDetails *ArtifactoryDetails) GetUrl() string {
	return artifactoryDetails.Url
}

func (artifactoryDetails *ArtifactoryDetails) GetDistributionUrl() string {
	return artifactoryDetails.DistributionUrl
}

func (artifactoryDetails *ArtifactoryDetails) GetUser() string {
	return artifactoryDetails.User
}

func (artifactoryDetails *ArtifactoryDetails) GetPassword() string {
	return artifactoryDetails.Password
}

func (artifactoryDetails *ArtifactoryDetails) GetAccessToken() string {
	return artifactoryDetails.AccessToken
}

func (artifactoryDetails *ArtifactoryDetails) GetRefreshToken() string {
	return artifactoryDetails.RefreshToken
}

func (artifactoryDetails *ArtifactoryDetails) GetClientCertPath() string {
	return artifactoryDetails.ClientCertPath
}

func (artifactoryDetails *ArtifactoryDetails) GetClientCertKeyPath() string {
	return artifactoryDetails.ClientCertKeyPath
}

func (artifactoryDetails *ArtifactoryDetails) CreateArtAuthConfig() (auth.ServiceDetails, error) {
	artAuth := artifactoryAuth.NewArtifactoryDetails()
	artAuth.SetUrl(artifactoryDetails.Url)
	return artifactoryDetails.createArtAuthConfig(artAuth)
}

func (artifactoryDetails *ArtifactoryDetails) CreateDistAuthConfig() (auth.ServiceDetails, error) {
	artAuth := distributionAuth.NewDistributionDetails()
	artAuth.SetUrl(artifactoryDetails.DistributionUrl)
	return artifactoryDetails.createArtAuthConfig(artAuth)
}

func (artifactoryDetails *ArtifactoryDetails) createArtAuthConfig(details auth.ServiceDetails) (auth.ServiceDetails, error) {
	details.SetSshUrl(artifactoryDetails.SshUrl)
	details.SetAccessToken(artifactoryDetails.AccessToken)
	// If refresh token is not empty, set a refresh handler and skip other credentials
	if artifactoryDetails.RefreshToken != "" {
		tokenRefreshServerId = artifactoryDetails.ServerId
		details.AppendPreRequestInterceptor(AccessTokenRefreshPreRequestInterceptor)
	} else {
		details.SetApiKey(artifactoryDetails.ApiKey)
		details.SetUser(artifactoryDetails.User)
		details.SetPassword(artifactoryDetails.Password)
	}
	details.SetClientCertPath(artifactoryDetails.ClientCertPath)
	details.SetClientCertKeyPath(artifactoryDetails.ClientCertKeyPath)
	details.SetSshKeyPath(artifactoryDetails.SshKeyPath)
	details.SetSshPassphrase(artifactoryDetails.SshPassphrase)
	return details, nil
}

func (missionControlDetails *MissionControlDetails) GetAccessToken() string {
	return missionControlDetails.AccessToken
}

func (missionControlDetails *MissionControlDetails) SetAccessToken(accessToken string) {
	missionControlDetails.AccessToken = accessToken
}
