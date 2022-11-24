// Copyright 2018 Drone.IO Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package drone

import (
	"net/http"
)

// TODO(bradrydzewski) add repo + latest build endpoint
// TODO(bradrydzewski) add queue endpoint
// TDOO(bradrydzewski) add stats endpoint
// TODO(bradrydzewski) add version endpoint

// Client is used to communicate with a Drone server.
type Client interface {
	// SetClient sets the http.Client.
	SetClient(*http.Client)

	// SetAddress sets the server address.
	SetAddress(string)

	// Self returns the currently authenticated user.
	Self() (*User, error)

	// User returns a user by login.
	User(login string) (*User, error)

	// UserList returns a list of all registered users.
	UserList() ([]*User, error)

	// UserCreate creates a new user account.
	UserCreate(user *User) (*User, error)

	// UserUpdate updates a user account.
	UserUpdate(login string, user *UserPatch) (*User, error)

	// UserDelete deletes a user account.
	UserDelete(login string) error

	// Incomplete returns a list of incomplete builds.
	Incomplete() ([]*Repo, error)

	// IncompleteV2 returns a list of builds/repos/stages that are running/pending.
	IncompleteV2() ([]*RepoBuildStage, error)

	// Repo returns a repository by name.
	Repo(namespace, name string) (*Repo, error)

	// RepoList returns a list of all repositories to which
	// the user has explicit access in the host system.
	RepoList() ([]*Repo, error)

	// RepoListSync returns a list of all repositories to which
	// the user has explicit access in the host system.
	RepoListSync() ([]*Repo, error)

	// RepoListAll returns a list of all repositories in
	// the database. This is only available to system admins.
	RepoListAll(opts ListOptions) ([]*Repo, error)

	// RepoEnable activates a repository.
	RepoEnable(namespace, name string) (*Repo, error)

	// RepoUpdate updates a repository.
	RepoUpdate(namespace, name string, repo *RepoPatch) (*Repo, error)

	// RepoChown updates a repository owner.
	RepoChown(namespace, name string) (*Repo, error)

	// RepoRepair repairs the repository hooks.
	RepoRepair(namespace, name string) error

	// RepoDisable disables a repository.
	RepoDisable(namespace, name string) error

	// RepoDelete permanetnly deletes a repository.
	RepoDelete(namespace, name string) error

	// Build returns a repository build by number.
	Build(namespace, name string, build int) (*Build, error)

	// BuildLast returns the latest build by branch. An
	// empty branch will result in the default branch.
	BuildLast(namespace, name, branch string) (*Build, error)

	// BuildList returns a list of recent builds for the
	// the specified repository.
	BuildList(namespace, name string, opts ListOptions) ([]*Build, error)

	// BuildCreate creates a new build by branch or commit.
	BuildCreate(owner, name, commit, branch string, params map[string]string) (*Build, error)

	// BuildRestart re-starts a build.
	BuildRestart(namespace, name string, build int, params map[string]string) (*Build, error)

	// BuildCancel stops the specified running job for
	// given build.
	BuildCancel(namespace, name string, build int) error

	// BuildPurge purges the build history.
	BuildPurge(namespace, name string, before int) error

	// Approve approves a blocked build stage.
	Approve(namespace, name string, build, stage int) error

	// Decline declines a blocked build stage.
	Decline(namespace, name string, build, stage int) error

	// Promote promotes a build to the target environment.
	Promote(namespace, name string, build int, target string, params map[string]string) (*Build, error)

	// Rollback reverts the target environment to an previous build.
	Rollback(namespace, name string, build int, target string, params map[string]string) (*Build, error)

	// Logs gets the logs for the specified step.
	Logs(owner, name string, build, stage, step int) ([]*Line, error)

	// LogsPurge purges the build logs for the specified step.
	LogsPurge(owner, name string, build, stage, step int) error

	// Secret returns a secret by name.
	Secret(owner, name, secret string) (*Secret, error)

	// SecretList returns a list of all repository secrets.
	SecretList(owner, name string) ([]*Secret, error)

	// SecretCreate creates a registry.
	SecretCreate(owner, name string, secret *Secret) (*Secret, error)

	// SecretUpdate updates a registry.
	SecretUpdate(owner, name string, secret *Secret) (*Secret, error)

	// SecretDelete deletes a secret.
	SecretDelete(owner, name, secret string) error

	// OrgSecret returns a secret by name.
	OrgSecret(namespace, secret string) (*Secret, error)

	// OrgSecretList returns a list of all repository secrets.
	OrgSecretList(namespace string) ([]*Secret, error)

	// OrgSecretListAll returns a list of all repository secrets.
	OrgSecretListAll() ([]*Secret, error)

	// OrgSecretCreate creates a registry.
	OrgSecretCreate(namespace string, secret *Secret) (*Secret, error)

	// OrgSecretUpdate updates a registry.
	OrgSecretUpdate(namespace string, secret *Secret) (*Secret, error)

	// OrgSecretDelete deletes a secret.
	OrgSecretDelete(namespace, name string) error

	// Cron returns a cronjob by name.
	Cron(owner, name, cron string) (*Cron, error)

	// CronList returns a list of all repository cronjobs.
	CronList(owner string, name string) ([]*Cron, error)

	// CronCreate creates a cronjob.
	CronCreate(owner, name string, in *Cron) (*Cron, error)

	// CronDelete deletes a cronjob.
	CronDelete(owner, name, cron string) error

	// CronUpdate enables a cronjob.
	CronUpdate(owner, name, cron string, in *CronPatch) (*Cron, error)

	// CronExec executes a cronjob.
	CronExec(owner, name, cron string) error

	// Sign signs the yaml file.
	Sign(owner, name, file string) (string, error)

	// Verify verifies the yaml signature.
	Verify(owner, name, file string) error

	// Encrypt returns an encrypted secret
	Encrypt(owner, name string, secret *Secret) (string, error)

	// Queue returns a list of queue items.
	Queue() ([]*Stage, error)

	// QueuePause pauses queue operations.
	QueuePause() error

	// QueueResume resumes queue operations.
	QueueResume() error

	// Node returns a node by name.
	Node(name string) (*Node, error)

	// NodeList returns a list of all nodes.
	NodeList() ([]*Node, error)

	// NodeCreate creates a node.
	NodeCreate(in *Node) (*Node, error)

	// NodeDelete deletes a node.
	NodeDelete(name string) error

	// NodeUpdate updates a node.
	NodeUpdate(name string, in *NodePatch) (*Node, error)

	//
	// Move to autoscaler-go
	//

	// Server returns the named servers details.
	Server(name string) (*Server, error)

	// ServerList returns a list of all active build servers.
	ServerList() ([]*Server, error)

	// ServerCreate creates a new server.
	ServerCreate() (*Server, error)

	// ServerDelete terminates a server.
	ServerDelete(name string, force bool) error

	// AutoscalePause pauses the autoscaler.
	AutoscalePause() error

	// AutoscaleResume resumes the autoscaler.
	AutoscaleResume() error

	// AutoscaleVersion returns the autoscaler version.
	AutoscaleVersion() (*Version, error)

	// Template returns a template by name.
	Template(namespace string, name string) (*Template, error)

	// TemplateListAll returns a list of all templates.
	TemplateListAll() ([]*Template, error)

	// TemplateList returns a list of all templates by namespace
	TemplateList(namespace string) ([]*Template, error)

	// TemplateCreate creates a template.
	TemplateCreate(namespace string, template *Template) (*Template, error)

	// TemplateUpdate updates template data.
	TemplateUpdate(namespace string, name string, template *Template) (*Template, error)

	// TemplateDelete deletes a template.
	TemplateDelete(namespace string, name string) error
}
