# satogram-mcp

An MCP (Model Context Protocol) server for [Satogram](https://satogram.xyz) — send custom messages with sats to Lightning Network participants directly from any MCP-compatible AI client.

## Tools

| Tool | Description |
|------|-------------|
| `get_recipient_stats` | Get statistics about Satogram recipients (pubkeys, lightning addresses, total recipients, satograms sent, signups) |
| `create_satogram` | Create a new Satogram campaign. Returns a Lightning invoice that must be paid to start delivery |
| `check_invoice_status` | Check the payment status of a Satogram invoice (`OPEN`, `SETTLED`, `CANCELED`, `ACCEPTED`) |
| `get_satogram_status` | Get delivery progress of a campaign (success/failure counts, sats spent, overall status) |

## Install

```sh
go install github.com/Satograms/satogram-mcp@latest
```

Or build from source:

```sh
git clone https://github.com/Satograms/satogram-mcp.git
cd satogram-mcp
go build -o satogram-mcp .
```

## Configuration

### For Claude Desktop

Add to your `claude_desktop_config.json`:

- macOS: `~/Library/Application Support/Claude/claude_desktop_config.json`
- Windows: `%APPDATA%\Claude\claude_desktop_config.json`

```json
{
  "mcpServers": {
    "satogram": {
      "command": "satogram-mcp"
    }
  }
}
```

### For Claude Code

Run in a separate terminal (not inside a Claude Code session), then restart Claude Code:

```sh
claude mcp add satogram -- satogram-mcp
```

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `SATOGRAM_API_URL` | `https://api.satogram.xyz` | Override the Satogram API base URL |

## Usage

Once configured, ask your AI client to interact with Satogram. Example prompts:

- "How many recipients can Satogram reach?"
- "Send a satogram with the message 'Hello from the future!' spending 1000 sats"
- "Check the status of my satogram campaign"

### Workflow

1. **Create** a campaign with `create_satogram` — you get back a Lightning invoice
2. **Pay** the invoice with your Lightning wallet
3. **Verify** payment with `check_invoice_status`
4. **Track** delivery progress with `get_satogram_status`

## License

MIT
