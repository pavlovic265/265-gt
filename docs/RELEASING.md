# Release Process

This project uses [release-please](https://github.com/googleapis/release-please) for automated releases.

## How It Works

1. Merge PRs to `main` using [conventional commits](https://www.conventionalcommits.org/)
2. release-please automatically creates/updates a Release PR
3. When ready to release, merge the Release PR
4. release-please creates a GitHub Release (with tag)
5. The `release: [published]` event triggers GoReleaser
6. GoReleaser builds binaries and publishes to npm

## Flow

```
merge PR to main → Release PR created/updated
                        ↓
              merge Release PR → GitHub Release created (with tag)
                        ↓
              release event triggers release.yml
                        ↓
              GoReleaser → binaries built + npm publish
```

## Workflows

| Workflow | Trigger | Purpose |
|----------|---------|---------|
| `release-please.yml` | Push to main | Creates/updates Release PR |
| `release.yml` | GitHub Release published | Builds binaries, publishes to npm |

## Conventional Commits

| Prefix | Version Bump | Example |
|--------|--------------|---------|
| `feat:` | Minor (0.1.0) | `feat: add clone command` |
| `fix:` | Patch (0.0.1) | `fix: resolve SSH bug` |
| `feat!:` or `fix!:` | Major (1.0.0) | `feat!: breaking API change` |
| `chore:`, `docs:`, etc. | No bump | `chore: update deps` |

## Manual Release (Emergency)

If needed, you can still release manually:

```bash
make patch   # 0.62.0 → 0.62.1
make minor   # 0.62.0 → 0.63.0
make major   # 0.62.0 → 1.0.0
```
