package plugins

import "github.com/checkr/codeflow/server/agent"

func init() {
	agent.RegisterApi(GitPing{})
	agent.RegisterApi(GitCommit{})
	agent.RegisterApi(GitStatus{})
	agent.RegisterApi(Release{})
	agent.RegisterApi(DockerBuild{})
	agent.RegisterApi(DockerDeploy{})
	agent.RegisterApi(LoadBalancer{})
	agent.RegisterApi(WebsocketMsg{})
}

type State string

const (
	Waiting  State = "waiting"
	Running        = "running"
	Fetching       = "fetching"
	Building       = "building"
	Pushing        = "pushing"
	Complete       = "complete"
	Failed         = "failed"
	Deleting       = "deleting"
	Deleted        = "deleted"
)

type Type string

const (
	File     Type = "file"
	Env           = "env"
	Build         = "build"
	Internal      = "internal"
	External      = "external"
	Office        = "office"
)

type Action string

const (
	Create   Action = "create"
	Update          = "update"
	Destroy         = "destroy"
	Rollback        = "rollback"
	Status          = "status"
)

type Project struct {
	Slug       string `json:"slug"`
	Repository string `json:"repository"`
}

type Git struct {
	SshUrl        string `json:"gitSshUrl"`
	RsaPrivateKey string `json:"rsaPrivateKey" role:"secret"`
	RsaPublicKey  string `json:"rsaPublicKey" role:"secret"`
}

type GitCommit struct {
	Repository string `json:"repository"`
	User       string `json:"user"`
	Message    string `json:"message"`
	Ref        string `json:"ref"`
	Hash       string `json:"hash"`
	ParentHash string `json:"parentHash"`
}

type GitPing struct {
	Repository string `json:"repository"`
	User       string `json:"user"`
}

type GitStatus struct {
	Repository string `json:"repository"`
	User       string `json:"user"`
	Hash       string `json:"hash"`
	State      string `json:"state"`
	Context    string `json:"context"`
}

type Feature struct {
	Hash       string `json:"hash"`
	ParentHash string `json:"parentHash"`
	User       string `json:"user"`
	Message    string `json:"message"`
}

type Release struct {
	Id          string  `json:"id"`
	HeadFeature Feature `json:"headFeature"`
	TailFeature Feature `json:"tailFeature"`
}

type Listener struct {
	Port     int32  `json:"port"`
	Protocol string `json:"protocol"`
}

type ListenerPair struct {
	Source      Listener `json:"source"`
	Destination Listener `json:"destination"`
}

type Service struct {
	Action       Action     `json:"action"`
	Name         string     `json:"name"`
	Command      string     `json:"command"`
	Listeners    []Listener `json:"listeners"`
	Replicas     int64      `json:"replicas"`
	State        State      `json:"state"`
	StateMessage string     `json:"stateMessage"`
}

type DockerRegistry struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password" role:"secret"`
	Email    string `json:"email"`
}

type Docker struct {
	Image    string         `json:"image"`
	Registry DockerRegistry `json:"registry"`
}

type Secret struct {
	Key   string `json:"key"`
	Value string `json:"value" role:"secret"`
	Type  Type   `json:"type"`
}

type Arg struct {
	Key   string `json:"key"`
	Value string `json:"value" role:"secret"`
}

type DockerBuild struct {
	Action       Action         `json:"action"`
	State        State          `json:"state"`
	StateMessage string         `json:"stateMessage"`
	Project      Project        `json:"project"`
	Git          Git            `json:"git"`
	Feature      Feature        `json:"feature"`
	Registry     DockerRegistry `json:"registry"`
	BuildArgs    []Arg          `json:"buildArgs"`
	BuildLog     string         `json:"buildLog"`
	BuildError   string         `json:"buildError"`
	Image        string         `json:"image"`
}

type HeartBeat struct {
	Tick string `json:"tick"`
}

// Deploy
type DockerDeploy struct {
	Action             Action    `json:"action"`
	State              State     `json:"state"`
	StateMessage       string    `json:"stateMessage"`
	Project            Project   `json:"project"`
	Release            Release   `json:"release"`
	Docker             Docker    `json:"docker"`
	Services           []Service `json:"services"`
	Secrets            []Secret  `json:"secrets"`
	Timeout            int       `json:"timeout"`
	DeploymentStrategy string    `json:"deploymentStrategy"`
	Environment        string    `json:"environment"`
}

// LoadBalancer
type LoadBalancer struct {
	Action        Action         `json:"action"`
	State         State          `json:"state"`
	StateMessage  string         `json:"stateMessage"`
	Name          string         `json:"name"`
	Type          Type           `json:"type"`
	Project       Project        `json:"project"`
	Service       Service        `json:"service"`
	ListenerPairs []ListenerPair `json:"portPairs"`
	DNSName       string         `json:"dnsName"`
	Environment   string         `json:"environment"`
}

type WebsocketMsg struct {
	Channel string      `json:"channel"`
	Payload interface{} `json:"data"`
}
