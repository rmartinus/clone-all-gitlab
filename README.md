# clone-all
Clone all my gitlab projects

## Instructions
- `export GITLAB_TOKEN=<your-personal-token>`
- `export DEV_HOME=<your-dev-path>` eg. /Users/me/dev
- `export GITLAB_NAMESPACE=<your-gitlab-namespace>` eg. mnf-group/plexus
- if you have golang >= 1.11 installed, run: `make run`
- otherwise you can run ./bin/clone-all
- You can now search on all cloned code: `grep -rIi pennytel .`
