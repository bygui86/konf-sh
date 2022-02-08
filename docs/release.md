
# konf-sh - Release

`NEXT VERSION:` v0.5.0

## Automatic

```sh
make release NEW_VERSION=...
```

## Manual

1. Choose a new version
   ```sh
   NEW_VERSION="v0.5.0"
   ```
2. Create a new tag with choosen version
   ```sh
   git tag -a $NEW_VERSION -m "Tag for release $NEW_VERSION"
   ```
3. Push new tag to remote, triggering `release` GitHub Action
   ```sh
   git push origin $NEW_VERSION
   ```

## Simulate

```sh
goreleaser --snapshot --skip-publish --rm-dist
```

## Available mechanisms

- goreleaser
- GitHub Actions
- GitHub Package Registry
- PackagePublishing

## GitHub Actions

| Action  | Triggered by                                                                           | Steps                                             |
|---------|----------------------------------------------------------------------------------------|---------------------------------------------------|
| build   | push to master, push to branch features/\*\*, PR to master, PR to branch features/\*\* | setup go, checkout, get dependencies, build, test |
| release | new tag creation                                                                       | setup go, checkout, unshallow, run goreleaser     |

## goreleaser

`WARN`: The first three steps will trigger the `release` GitHub Action, performing the last step (goreleaser), so be careful if you want to release manually.

1. version
   ```sh
   NEW_VERSION="v0.5.0"
   ```
2. tag
   ```sh
   git tag -a $NEW_VERSION -m "Tag for release $NEW_VERSION"
   ```
3. push
   ```sh
   git push origin $NEW_VERSION
   ```
4. release
   ```sh
   goreleaser release --rm-dist
   ```

## In code

Version is specified only in `pkg/app/utils.go`
