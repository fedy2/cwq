# cwq
A small tool to import and export CloudWatch query definitions.

## Installation
Download and manually install the executable from the release tab.

### Installation via package manager
Coming soon.

## Usage
To get the a description of the available commands run:
```bash
$cwq --help
Usage: cwq <command>

A small tool to import and export CloudWatch query definitions.

Flags:
  -h, --help       Show context-sensitive help.
      --debug      Enable debug mode.
      --version    Print version information and quit

Commands:
  export
    Exports query definitions.

  import <file>
    Imports query definitions.

Run "cwq <command> --help" for more information on a command.
```

### AWS account
The tool uses the default AWS account configured for the user.

### Export command
The `export` command exports the list of CloudWatch query definitions present in an AWS account.
Here the available options:
```bash
Usage: cwq export

Exports query definitions.

Flags:
  -h, --help             Show context-sensitive help.
      --debug            Enable debug mode.
      --version          Print version information and quit

  -o, --output=file      Write to file instead of stdout
      --format="json"    Output format: csv or json
```
#### Limitations
* The current version does not support the paging of the response, so if there are too many definitions the export can result incomplete.

### Import command
The `import` command imports a list of CloudWatch query definitions into an AWS account.
Here the available options:
```bash
Usage: cwq import <file>

Imports query definitions.

Arguments:
  <file>    File containing the query descriptions

Flags:
  -h, --help             Show context-sensitive help.
      --debug            Enable debug mode.
      --version          Print version information and quit

      --format="json"    File format: csv or json
```

#### Limitations
* If the query definition entry contains a `QueryDefinitionId` this is passed to AWS that will try to update an existing query definition.
* The tool tries to import the query one by one, if one of the imports fails the import process is stopped with a risk of incomplete imports.
* The columns for the CSV format at the moment have a fixed position and the header is required.


## Disclaimer
This is a personal project to learn Go.
