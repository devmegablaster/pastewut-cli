# PasteWut CLI

Command Line Interface for PasteWut, a PasteBin written in golang for easily sharing text and code just with an AlphaNumeric code

## Installation

Install pswt using homebrew

```bash
  brew tap devmegablaster/devmegablaster
  brew install pswt
```

Or install the Latest release from the [release page](https://github.com/devmegablaster/pastewut-cli/releases)

## Run Locally

Clone the project

```bash
  git clone https://github.com/devmegablaster/pastewut-cli
```

Go to the project directory

```bash
  cd pastewut-cli
```

Start the CLI

```bash
  go run .
```

## Usage/Examples

- Get help on Commands and Usage

```
pswt -h
```

- Create a new PasteWut

```
pswt n
```

- Create a new PasteWut from your clipboard content

```
pswt n -c
```

- Get PasteWut contents from the code

```
pswt g <code>
```

- Get PasteWut contents copied to your clipboard

```
pswt g -c <code>
```

## Contributing

Contributions are always welcome!

Guidlines on contributing will be updated soon, feel free to clone and make a PR

## Authors

- [@devmegablaster](https://www.github.com/devmegablaster)
