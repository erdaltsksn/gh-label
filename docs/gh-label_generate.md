## gh-label generate

Generate labels using a list

### Synopsis

Generate labels using predefined label list or a custom label file.

```
gh-label generate [flags]
```

### Examples

```
# Generate the labels using a predefined list
gh-label generate --repo erdaltsksn/playground --list "insane"

# User custom file as a list to generate the labels
gh-label generate --repo erdaltsksn/playground --file my-labels.json

# DANGER: Remove all the labels before generating the labels
gh-label generate --repo erdaltsksn/playground --list "insane" --force
```

### Options

```
      --file string   Use file as a label list. User --list "file.json"
  -f, --force         This will remove all labels before generate the labels.
  -h, --help          help for generate
  -l, --list string   Predefined label list name. Use --list "ABC"
```

### Options inherited from parent commands

```
  -r, --repo string   Repository which its labels will be generated or exported into a file.
                      Please use 'username/repo-name' format.
```

### SEE ALSO

* [gh-label](gh-label.md)	 - This app helps you manage GitHub issue labels.

