# Jayteal

> Convert JTL XML reports into CSV files

## Installation

From source:

```
go get github.com/stayradiated/jayteal
```

Or using the included binary (mac os only)

```
git clone https://github.com/stayradiated/jayteal
cd jayteal
./jayteal
```

## Usage

```
jayteal --src report.jst --dst out.jst
```

You can set the progress bar maximum (by default 100000) using `-n`

```
jayteal --src report.jst --dst out.jst -n 1234
```
