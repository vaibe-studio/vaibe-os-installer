# vaibe-os-installer

> The installer for [vAIbe-OS](https://github.com/vaibe-studio/vaibe-os) — creates a new
> instance workspace from the embedded instance skeleton.

The core of vAIbe-OS is **declarative** (documentation, instance skeleton, canon) — this
repository is its executable counterpart: a small Go CLI that installs and initializes
workspaces.

## Usage

```bash
vaibe-os init ~/Projects/my-os [--name "My OS"]
```

Creates the workspace: `projects/ · inbox/ · wiki/ · teams/` + task format v2
(`README.md` for humans, `AGENTS.md` for the agent, `attachments/` via git-lfs,
`results/v{N}/`), instance canon stubs and `os/config.yml` — then `git init`s the folder.

## Install

```bash
curl -fsSL https://os.vaibe.studio/install | bash
```

> Not live yet — the landing page ships with v0.2.0. Manual: download the binary from
> [Releases](https://github.com/vaibe-studio/vaibe-os-installer/releases) and put it on
> your `PATH`.

## Layout

- `cmd/vaibe-os/` — CLI entry point
- `internal/installer/` — init logic + embedded skeleton
- `install/` — installer scripts (bash / PowerShell): download the binary from Releases
- `scripts/sync-skeleton.sh` — sync the embedded skeleton from the vAIbe-OS core checkout

## Skeleton sync

`internal/installer/skeleton/` mirrors `skeleton/` of the core repository
(`vaibe-studio/vaibe-os`) — never hand-edit it; sync via
`scripts/sync-skeleton.sh <core-checkout>` when the core skeleton changes.

## Release flow

GitLab CI builds binaries on tags (`v*`) and publishes them to GitHub Releases
(requires the `GITHUB_TOKEN` CI variable). Versioning starts at `v0.2.0` — the
fresh-main era of vAIbe-OS.
