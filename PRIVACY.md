# Privacy Policy — labctl (unofficial Windows build)

_Last updated: 2026-11-05_

This policy covers the unofficial Windows build of `labctl` published from
[rodolphobrock/labctl](https://github.com/rodolphobrock/labctl) (including the
`RodolphoBrock.labctl` WinGet package and the `rodolphobrock/labctl` Scoop bucket).
It is maintained by Rodolpho Brock and is not affiliated with iximiuz Labs.

## Summary

- `labctl` has **no telemetry, analytics, crash reporting or tracking** of any kind.
- It only sends data to the **iximiuz Labs** service (`labs.iximiuz.com`), and only to
  carry out the commands you run.
- The maintainer of this build **does not receive, collect or store any of your data**.

## What is sent, and where

`labctl` is a command-line client for [iximiuz Labs](https://labs.iximiuz.com).
When you run a command, it talks to the iximiuz Labs API and its related endpoints:

| Data | When | Sent to |
|---|---|---|
| Session ID and access token (in the `Authorization` header) | Every API request after `labctl auth login` | iximiuz Labs API |
| `labctl/<version>` (in the `User-Agent` header) | Every API request | iximiuz Labs API |
| SSH public key generated on your machine | Opening a connection to a playground (`ssh`, `ide`, `port-forward`, …) | iximiuz Labs API |
| Terminal input/output and copied files | `ssh`, `cp`, `port-forward`, `expose` and similar commands | Your own playground, through iximiuz Labs |
| Content files you choose to publish | `labctl content push` | iximiuz Labs and the storage upload URLs it provides |

No other host is contacted by `labctl` itself. Commands such as `labctl ide` may start
local tools (for example your IDE), which follow their own privacy policies.

How iximiuz Labs handles your data (account data, cookies, third parties) is described in
section 3, "Data Usage and Privacy", of its terms:
<https://labs.iximiuz.com/terms-and-conditions>.

## What is stored on your machine

- `%USERPROFILE%\.iximiuz\labctl\config.yaml` — API endpoint, session ID and access token.
- `%USERPROFILE%\.iximiuz\labctl\plays\` — local state for playgrounds you use.
- `%USERPROFILE%\.ssh\iximiuz_labs_user` (and `.pub`) — the SSH key pair used to reach playgrounds.

`labctl auth logout` ends the session on iximiuz Labs, removes it from the config file and
deletes the generated SSH key pair. Uninstalling the package does not delete these files; you can remove the
`.iximiuz\labctl` folder and the key pair manually.

## Retention and deletion

This build keeps nothing outside your machine. Data held by iximiuz Labs is governed by
its terms; contact iximiuz Labs to access or delete your account data there.

## Contact

Questions about this build: contact the maintainer through GitHub,
<https://github.com/rodolphobrock>.
