# Discordo &middot; [![discord](https://img.shields.io/discord/1297292231299956788?color=5865F2&logo=discord&logoColor=white)](https://discord.com/invite/VzF9UFn2aB) [![ci](https://github.com/ayn2op/discordo/actions/workflows/ci.yml/badge.svg)](https://github.com/ayn2op/discordo/actions/workflows/ci.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/ayn2op/discordo)](https://goreportcard.com/report/github.com/ayn2op/discordo) [![license](https://img.shields.io/github/license/ayn2op/discordo?logo=github)](https://github.com/ayn2op/discordo/blob/master/LICENSE)

Discordo is a lightweight, secure, and feature-rich Discord terminal client. Heavily work-in-progress, expect breaking changes.

![Preview](.github/preview.png)

## Added Features
- **Bot token authentication** - Uses bot tokens instead of user tokens to comply with Discord's Terms of Service

## Installation

### Building from source

```bash
git clone https://github.com/DevKyurria/discordo-bot
cd discordo-bot
go build .
```

### Wayland clipboard support

`wl-clipboard` is required for clipboard support.

## Usage

### Bot Token (UI, recommended)

1. Run the `discordo` executable with no arguments.

2. Enter your bot token and click on the "Login" button to save it.

> [!NOTE]
> Discordo uses bot tokens rather than user tokens to comply with Discord's Terms of Service. Create a bot application at Discord Developer Portal and use its token.

## Configuration

The configuration file allows you to configure and customize the behavior, keybindings, and theme of the application.

- Unix: `$XDG_CONFIG_HOME/discordo/config.toml` or `$HOME/.config/discordo/config.toml`
- Darwin: `$HOME/Library/Application Support/discordo/config.toml`
- Windows: `%AppData%/discordo/config.toml`

Discordo uses the default configuration if a configuration file is not found in the aforementioned path; however, the default configuration file is not written to the path. [The default configuration can be found here](./internal/config/config.toml).

## Troubleshooting

## Bot shows no servers
- Ensure your bot has been invited to at least one server
- Verify the bot is online in the server member list

##Can't see members or mentions
- Check that Server Members Intent is enabled in Developer Portal
- Re-invite the bot after enabling the intent

##Can't read messages
- Verify Message Content Intent is enabled
- Ensure the bot has Send Messages and Read Message History permissions

>[!IMPORTANT]
>This client is designed for bot tokens to comply with Discord's Terms of Service.
>Your bot must have the required permissions and intents configured correctly before using Discordo.

## License

Copyright (C) 2026-present ayn2op

This project is licensed under the GNU General Public License v3.0 (GPL-3.0).
See the [LICENSE](./LICENSE) file for the full license text.
