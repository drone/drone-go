package drone

type User struct {
	ID       int64  `json:"-"`
	Remote   string `json:"remote"`
	Login    string `json:"login"`
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"`
	Gravatar string `json:"gravatar"`
	Admin    bool   `json:"admin"`
	Active   bool   `json:"active"`
	Syncing  bool   `json:"syncing"`
	Created  int64  `json:"created_at"`
	Updated  int64  `json:"updated_at"`
}

type Repo struct {
	Remote      string `json:"remote"`
	Host        string `json:"host"`
	Owner       string `json:"owner"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	CloneURL    string `json:"clone_url"`
	GitURL      string `json:"git_url"`
	SSHURL      string `json:"ssh_url"`
	Active      bool   `json:"active"`
	Private     bool   `json:"private"`
	Privileged  bool   `json:"privileged"`
	PostCommit  bool   `json:"post_commits"`
	PullRequest bool   `json:"pull_requests"`
	PrivateKey  string `json:"private_key"`
	Timeout     int64  `json:"timeout"`
	Created     int64  `json:"created_at"`
	Updated     int64  `json:"updated_at"`
	// NOTE: these fields will be blank unless the token used has admin privileges
	Params    string `json:"params"`
	PublicKey string `json:"public_key"`
}

type Commit struct {
	ID          int64  `json:"id"`
	Status      string `json:"status"`
	Started     int64  `json:"started_at"`
	Finished    int64  `json:"finished_at"`
	Duration    int64  `json:"duration"`
	Sha         string `json:"sha"`
	Branch      string `json:"branch"`
	PullRequest string `json:"pull_request"`
	Author      string `json:"author"`
	Gravatar    string `json:"gravatar"`
	Timestamp   string `json:"timestamp"`
	Message     string `json:"message"`
	Created     int64  `json:"created_at"`
	Updated     int64  `json:"updated_at"`
}

// Returns the Short (--short) Commit Hash.
func (c *Commit) ShaShort() string {
	if len(c.Sha) > 8 {
		return c.Sha[:8]
	} else {
		return c.Sha
	}
}
