# AGENTS.md — developing the vaibe-os-installer

Read [README.md](README.md) first; the system contract lives in the vAIbe-OS core
repository (`vaibe-studio/vaibe-os`, `docs/`).

## Layout

- `cmd/vaibe-os/` — CLI entry point (binary name: `vaibe-os`).
- `internal/installer/` — init logic + embedded instance skeleton.
- `install/` — installer scripts (bash / PowerShell).
- `scripts/sync-skeleton.sh` — skeleton sync from the core checkout.

## Rules for this repository

- This is the **installer only**: no canon content, no rules or skills — those live in the
  vAIbe-OS core. The embedded skeleton is a **synced mirror** — never hand-edit it.
- Go project layout: `cmd/` + `internal/`; keep functions small and imperative.
- GitLab CI builds binaries on tags (`v*`) and publishes them to GitHub Releases
  (`GITHUB_TOKEN` CI variable).
- Changes via PR; merge and tag (semver) by the owner. `CHANGELOG.md` per release.
