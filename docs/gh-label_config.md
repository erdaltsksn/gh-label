## gh-label config

Configure the application

### Synopsis

Set up GitHub token and any configurations needed for this app works.

```
gh-label config [flags]
```

### Examples

```
gh-label config --token <GITHUB_TOKEN>
```

### Options

```
  -h, --help           help for config
  -t, --token string   A personal access token is required to access private repositories.
                       You can generate your token here: https://github.com/settings/tokens/new
```

### Options inherited from parent commands

```
  -r, --repo string   Repository which its labels will be generated or exported into a file.
                      Please use 'username/repo-name' format.
```

### SEE ALSO

* [gh-label](gh-label.md)	 - This app helps you manage GitHub issue labels.

