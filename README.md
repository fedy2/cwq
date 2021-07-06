# cwq
A small tool to import and export CloudWatch query definitions.

Withi this tool you can:
* backup your query definitions
* restore your query definitions
* copy the query definitions between different AWS accounts
* generate/produce your query definitions locally and then import them

## Installation
Download and manually install the executable from the [release page](https://github.com/fedy2/cwq/releases).

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
  -h, --help                       Show context-sensitive help.
      --debug                      Enable debug mode.
      --version                    Print version information and quit

  -o, --output=file                Write to file instead of stdout
  -f, --format="json"              Output format: csv or json. Default: json
      --prefix=STRING              Export the query definitions with a name that starts with the specified prefix.
      --[no-]csv-include-header    Add an header to the CSV file. Default: true
      --csv-delimiter=,            Field separator in the CSV file. Default: ,
```

### Import command
The `import` command imports a list of CloudWatch query definitions into an AWS account.
Here the available options:
```bash
Usage: cwq import <file>

Imports query definitions.

Arguments:
  <file>    File containing the query descriptions

Flags:
  -h, --help                   Show context-sensitive help.
      --debug                  Enable debug mode.
      --version                Print version information and quit

  -f, --format="json"          File format: csv or json. Default: json
      --clear-ids              Strip out the ids from the query definitions
      --[no-]csv-has-header    The CSV file has an header. Default: true
      --csv-delimiter=,        Field separator in the CSV file. Default: ,
      --csv-comment=#          Comment character in the CSV file. Default: #
```

#### Limitations
* The tool tries to import the query one by one, if one of the imports fails the import process is stopped with a risk of incomplete imports.
* The columns for the CSV format at the moment have a fixed position.


## Disclaimer
This is a personal project to learn Go.
