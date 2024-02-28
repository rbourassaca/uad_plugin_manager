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
3. Place your `UADSystemProfile.txt` in the same folder as the executable [^1]
4. Open the terminal in the directory you just downloaded
5. Install the project dependency `go get .`
6. Run de project using `go run .`, you'll see available commands in the terminal.

### Commands

- Remove all unlicensed plugins `remove -u` [^2]
- Remove specific collections of plugins `remove -n "UAD Ampeg B15N Bass Amplifier" -n "UAD SPL Vitalizer MK2-T"` [^2]
- Display all available collections and the plugins they contain `list -a`

### Examples

- `go run . remove -u` [^2]
- `go run . remove -n "UAD Ampeg B15N Bass Amplifier" -n "UAD SPL Vitalizer MK2-T"` [^2]
- `go run . list -a`


[^1]: To get your UADSystemProfile.txt be sure your audio interface is open and connected, open `UAD Meter & Control Panel`, click on the hamburger menu and click `System Info...`. In the new window, be sure to be in the System Info tab and click `Save Detailed System Profile` at the bottom of the page
[^2]: Some commands may need to run with elevated privileges since they handle protected files. On windows you'll need to run the terminal as an administrator before using the commands and on Mac OS X you'll need to prepend sudo to the commands.
