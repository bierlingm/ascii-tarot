# ASCII Tarot

A terminal tarot card reader written in Go, with ASCII art rendering and JSON output for scripting.

Forked from [asyapluggedin/ascii-tarot](https://github.com/asyapluggedin/ascii-tarot) -- the original project by Asya, who designed the card art and built the initial version. This fork adds a CLI interface with spreads, card lookup, and machine-readable output.

## Usage

```
tarot draw [n]          Draw n random cards (default: 1)
tarot spread [type]     Do a spread (three, celtic)
tarot list              List all cards
tarot card <id>         Show a specific card

Flags:
  --json                Output as JSON
```

## Examples

```sh
# Draw a single card
tarot draw

# Three-card past/present/future spread
tarot spread three

# Celtic cross
tarot spread celtic

# JSON output for piping into other tools
tarot draw 3 --json
```

## Building from source

Requires Go 1.19+.

```sh
go build -o tarot .
```

## Credits

Original card art and project by [@asyapluggedin](https://github.com/asyapluggedin). Text extraction script by [@MOCSABNIMAJNEB](https://github.com/MOCSABNIMAJNEB). Technology guidance and pair programming by [@dwimmertxt](https://github.com/dwimmertxt).
