⚠️ Currently in development, it's expected to have bugs and missing features, no builds will be provided during the development process. ⚠️

---

# UAD Plugin Manager

UAD Plugin Manager is a simple tool used to manage Universal Audio DSP powered plugins.

## Why?

When you install or upgrade the drivers for your Universal Audio interface, the installer reinstalls every plugins. Even if you don't own them. This makes browsing plugins in a DAW harder since you need to remember exactly which one you own.

You could hide or remove the plugins by hand, but, with every update it becomes a time consuming process. This tool was created to simplify this process.

## How to use (development)

1. Have [GO](https://go.dev/) installed and working
2. Clone or download the repository
3. Open the terminal as an administrator in the directory you just downloaded
4. Install the project dependency ` go get .`
5. Run de project using `go run .`, you'll see available commands in the terminal.

### Commands

- Remove all unlicensed plugins `remove -u`
- Remove specific collections of plugins `remove -n "UAD Ampeg B15N Bass Amplifier" -n "UAD SPL Vitalizer MK2-T"`
- Display all available collections and the plugins they contain `list -a`

### Examples (development)

- `go run . remove -u`
- `go run . remove -n "UAD Ampeg B15N Bass Amplifier" -n "UAD SPL Vitalizer MK2-T"`
- `go run . list -a`
