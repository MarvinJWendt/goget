# goget

## Usage
> A simple go module downloader

goget

## Description

```
Goget is a simple go module downloader.
```
## Examples

```bash
goget
goget pterm
```

## Flags
|Flag|Usage|
|----|-----|
|`--debug`|enable debug messages|
|`--disable-update-checks`|disables update checks|
|`--raw`|print unstyled raw output (set it if output is written to a file)|

## Commands
|Command|Usage|
|-------|-----|
|`goget completion`|Generate the autocompletion script for the specified shell|
|`goget date`|Prints the current date.|
|`goget help`|Help about any command|
# ... completion
`goget completion`

## Usage
> Generate the autocompletion script for the specified shell

goget completion

## Description

```
Generate the autocompletion script for goget for the specified shell.
See each sub-command's help for details on how to use the generated script.

```

## Commands
|Command|Usage|
|-------|-----|
|`goget completion bash`|Generate the autocompletion script for bash|
|`goget completion fish`|Generate the autocompletion script for fish|
|`goget completion powershell`|Generate the autocompletion script for powershell|
|`goget completion zsh`|Generate the autocompletion script for zsh|
# ... completion bash
`goget completion bash`

## Usage
> Generate the autocompletion script for bash

goget completion bash

## Description

```
Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(goget completion bash)

To load completions for every new session, execute once:

#### Linux:

	goget completion bash > /etc/bash_completion.d/goget

#### macOS:

	goget completion bash > /usr/local/etc/bash_completion.d/goget

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion fish
`goget completion fish`

## Usage
> Generate the autocompletion script for fish

goget completion fish

## Description

```
Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	goget completion fish | source

To load completions for every new session, execute once:

	goget completion fish > ~/.config/fish/completions/goget.fish

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion powershell
`goget completion powershell`

## Usage
> Generate the autocompletion script for powershell

goget completion powershell

## Description

```
Generate the autocompletion script for powershell.

To load completions in your current shell session:

	goget completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion zsh
`goget completion zsh`

## Usage
> Generate the autocompletion script for zsh

goget completion zsh

## Description

```
Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions for every new session, execute once:

#### Linux:

	goget completion zsh > "${fpath[1]}/_goget"

#### macOS:

	goget completion zsh > /usr/local/share/zsh/site-functions/_goget

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... date
`goget date`

## Usage
> Prints the current date.

goget date

## Flags
|Flag|Usage|
|----|-----|
|`-f, --format string`|specify a custom date format (default "02 Jan 06")|
# ... help
`goget help`

## Usage
> Help about any command

goget help [command]

## Description

```
Help provides help for any command in the application.
Simply type goget help [path to command] for full details.
```


---
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 15 March 2022**
