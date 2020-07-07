## gh-label export

Export GitHub label list into a file

### Synopsis

Find relevant repository and export label list into a file. You can
state which directory the file will be out using parameter.

```
gh-label export [flags]
```

### Examples

```
# Export the labels into a file at the current directory
gh-label export --repo erdaltsksn/playground

# Export the labels into a file by specifying absolute file path
gh-label export --repo erdaltsksn/playground --out ~/Desktop/mylabels.json
```

### Options

```
  -h, --help         help for export
  -o, --out string   Output file which contains label list will be save here.
                     Use 'directory/filename.json' format
```

### Options inherited from parent commands

```
  -r, --repo string   Repository which its labels will be generated or exported into a file.
                      Please use 'username/repo-name' format.
```

### SEE ALSO

* [gh-label](gh-label.md)	 - This app helps you manage GitHub issue labels.

