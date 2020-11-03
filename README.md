<p align="center">
  <h3 align="center">EsCharGo</h3>
  <h4 align="center">Escape Regex Metacharacters, made in Go</h3>
</p>

---

![preview](https://user-images.githubusercontent.com/31771429/97931648-b6106f00-1d33-11eb-8dfc-dd972e76ccb5.gif)

* [Installation](#installation)

* [Usage](#usage)

* [Limitation](#limitation)

* [License](#license)

## Installation

```bash
# brew
brew install p-rs/eschargo-homebrew-tap/formula

# scoop
scoop bucket add eschargo https://github.com/p-rs/eschargo-scoop-bucket.git

# via go build
go build -o ecg github.com/p-rs/eschargo

# or via releases
tar xzf file.tgz /usr/local/bin/ecg
```

## Usage

Run any standard in/out through `ecg`:

```bash
# output: \[02/Nov/2020:21:50:22 \+0000\]
ecg '[02/Nov/2020:21:50:22 +0000]'

# outputs \[02/Nov/2020:21:50:22 \+0000\] to MacOS clipboard
ecg '[02/Nov/2020:21:50:22 +0000]' | pbcopy

# output: escapes characters in example.txt
cat example.txt | ecg

# outputs escapes characters in example.txt and outputs to MacOS clipboard
cat example.txt | tr -d '\n' | ecg | pbcopy
```

## Limitation

`ecg` does not account for bash/history escaping, so make sure to use single quotes.

## License

MIT. See `LICENSE` for more details.
