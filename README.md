# go-package-template

Go package template to reuse code and utilities.

# Install package

```shell
go get -v github.com/carseven/go-package-template@v1.0.0
```

## Install configuration for private repositories

### Configure git to use ssh credentials for private remote repository (Not Github) (Optional)

This step is only need if using this template from a non github remote server or private repositories

<details>
<summary>Check if already have the configuration</summary>

```shell
cat ~/.gitconfig
```

We should have this to be able to map example-git-repo-url.com to use SSH credentials

```shell
[url "ssh://git@example-git-repo-url.com/"]
insteadOf = https://example-git-repo-url.com/
```

</details>

```shell
git config --global ssh://git@example-git-repo-url.com/:.insteadOf https://example-git-repo-url.com/
```

### Configure GOPRIVATE

For Go modules to work (with Go 1.11 or newer), you'll also need to set the GOPRIVATE variable, to avoid using the public servers to fetch the code:

<details>
<summary>Check GOPRIVATE value</summary>

```shell
go env | grep GOPRIVATE
```

</details>

Add go private https://example-git-repo-url.com/

All repositories of a URL

```shell
go env -w GOPRIVATE="example-git-repo-url.com/*"
```

or just this repo

```shell
go env -w GOPRIVATE="example-git-repo-url.com/specific-repository-name/common"
```

# Versioning

Follow the semantic versioning for updating the package versions. For more detail, follow the GO [guide](https://go.dev/doc/modules/version-numbers).

## Tags

For updating the package version, just create new tags following SemVer.

### PR

Use the merge commit message to set the version bump. Make sure to follow this rules:

**Manual Bumping:** Any commit message that includes `#major`, `#minor`, `#patch`, or `#none` will trigger the respective version bump. If two or more are present, the highest-ranking one will take precedence.
If `#none` is contained in the merge commit message, it will skip bumping regardless `DEFAULT_BUMP`.

**Automatic Bumping:** If no `#major`, `#minor` or `#patch` tag is contained in the merge commit message, it will bump whichever `DEFAULT_BUMP` is set to (which is `patch` by default).

> **_Note:_** This action **will not** bump the tag if the `HEAD` commit has already been tagged.

### Manual

Make sure you are on main branch and the branch is up to date.

```shell
git switch main
git fetch --all
git pull
```

Create a new version to main HEAD commit.

```shell
git tag v1.0.0
```

Push new version to origin/main:

```shell
git push --tags
```

### Github releases

From github you could also create a new version and tag. For more detail, follow the github [guide](https://docs.github.com/en/repositories/releasing-projects-on-github/managing-releases-in-a-repository).
