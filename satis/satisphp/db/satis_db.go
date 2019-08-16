package db

type SatisDb struct {
	Abandoned              map[string]string `json:"abandoned,omitempty"`
	Archive                SatisArchive      `json:"archive,omitempty"`
	Comment                string            `json:"_comment,omitempty"`
	Config                 SatisConfig       `json:"config,omitempty"`
	Description            string            `json:"description,omitempty"`
	Homepage               string            `json:"homepage"`
	IncludeFilename        string            `json:"include-filename,omitempty"`
	MinimumStability       string            `json:"minimum-stability,omitempty"`
	Name                   string            `json:"name"`
	NotifyBatch            string            `json:"notify-batch,omitempty"`
	OutputDir              string            `json:"output-dir,omitempty"`
	OutputHTML             bool              `json:"output-html,omitempty"`
	Providers              bool              `json:"providers,omitempty"`
	Repositories           []SatisRepository `json:"repositories,omitempty"`
	Require                map[string]string `json:"require,omitempty"`
	RequireAll             bool              `json:"require-all,omitempty"`
	RequireDependencies    bool              `json:"require-dependencies,omitempty"`
	RequireDevDependencies bool              `json:"require-dev-dependencies,omitempty"`
	TwigTemplate           string            `json:"twig-template,omitempty"`
}

type SatisRepository struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type SatisArchive struct {
	AbsoluteDirectory string   `json:"absolute-directory,omitempty"`
	Blacklist         []string `json:"blacklist,omitempty"`
	Checksum          bool     `json:"checksum,omitempty"`
	Directory         string   `json:"directory,omitempty"`
	Format            string   `json:"format,omitempty"`
	IgnoreFilters     bool     `json:"ignore-filters,omitempty"`
	PrefixURL         string   `json:"prefix-url,omitempty"`
	SkipDev           bool     `json:"skip-dev,omitempty"`
	Whitelist         []string `json:"whitelist,omitempty"`
}

type SatisConfig struct {
	ProcessTimeout        int                    `json:"process-timeout,omitempty"`
	UseIncludePath        bool                   `json:"use-include-path,omitempty"`
	PreferredInstall      string                 `json:"preferred-install,omitempty"`
	StoreAuths            bool                   `json:"store-auths,omitempty"`
	GithubProtocols       []string               `json:"github-protocols,omitempty"`
	GithubOauth           map[string]interface{} `json:"github-oauth,omitempty"`
	GitlabOauth           map[string]interface{} `json:"gitlab-oauth,omitempty"`
	GitlabToken           map[string]interface{} `json:"gitlab-token,omitempty"`
	DisableTLS            bool                   `json:"disable-tls,omitempty"`
	SecureHTTP            bool                   `json:"secure-http,omitempty"`
	BitbucketOauth        map[string]interface{} `json:"bitbucket-oauth-oauth,omitempty"`
	Cafile                string                 `json:"cafile,omitempty"`
	Capath                string                 `json:"capath,omitempty"`
	HTTPBasic             map[string]interface{} `json:"http-basic,omitempty"`
	Platform              map[string]interface{} `json:"platform,omitempty"`
	VendorDir             string                 `json:"vendor-dir,omitempty"`
	BinDir                string                 `json:"bin-dir,omitempty"`
	DataDir               string                 `json:"data-dir,omitempty"`
	CacheDir              string                 `json:"cache-dir,omitempty"`
	CacheFilesDir         string                 `json:"cache-files-dir,omitempty"`
	CacheRepoDir          string                 `json:"cache-repo-dir,omitempty"`
	CacheVcsDir           string                 `json:"cache-vcs-dir,omitempty"`
	CacheFilesTTL         int                    `json:"cache-files-ttl,omitempty"`
	CacheFilesMaxsize     string                 `json:"cache-files-maxsize,omitempty"`
	BinCompat             string                 `json:"bin-compat,omitempty"`
	PrependAutoloader     bool                   `json:"prepend-autoloader,omitempty"`
	AutoloaderSuffix      string                 `json:"autoloader-suffix,omitempty"`
	OptimizeAutoloader    bool                   `json:"optimize-autoloader,omitempty"`
	SortPackages          bool                   `json:"sort-packages,omitempty"`
	ClassmapAuthoritative bool                   `json:"classmap-authoritative,omitempty"`
	ApcuAutoloader        bool                   `json:"apcu-autoloader,omitempty"`
	GithubDomains         []string               `json:"github-domains,omitempty"`
	GithubExposeHostname  bool                   `json:"github-expose-hostname,omitempty"`
	GitlabDomains         []string               `json:"gitlab-domains,omitempty"`
	UseGithubAPI          bool                   `json:"use-github-api,omitempty"`
	NotifyOnInstall       bool                   `json:"notify-on-install,omitempty"`
	DiscardChanges        bool                   `json:"discard-changes,omitempty"`
	ArchiveFormat         string                 `json:"archive-format,omitempty"`
	ArchiveDir            string                 `json:"archive-dir,omitempty"`
	HtaccessProtect       bool                   `json:"htaccess-protect,omitempty"`
}
