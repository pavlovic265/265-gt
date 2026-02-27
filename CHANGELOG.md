# Changelog

## [0.71.0](https://github.com/pavlovic265/265-gt/compare/v0.70.0...v0.71.0) (2026-02-27)


### Features

* accept token and GPG key as direct CLI arguments in account edit ([27d681c](https://github.com/pavlovic265/265-gt/commit/27d681c333e672bb387bb66d35ff6d6d03df2c9c))
* add submit-stack (ss) command to push and create PRs for entire stack ([beac83b](https://github.com/pavlovic265/265-gt/commit/beac83b67a604d8f74c89543ec64cc3d25e1388f))
* add viewport scrolling to list component to prevent terminal overflow ([e168eb6](https://github.com/pavlovic265/265-gt/commit/e168eb6164525463eca6067c8cd222bda6a37f98))
* show merge-queued status in PR list ([ddb5fab](https://github.com/pavlovic265/265-gt/commit/ddb5fabad7cb0c20a834209cfcaa25d313bbaeb9))


### Bug Fixes

* add trailing newline to status output to prevent zsh % indicator ([716f541](https://github.com/pavlovic265/265-gt/commit/716f541016298214f9db112f3469c5a8a27f3ea2))
* **pr:** show merge-queued indicator for GitHub PRs ([b101202](https://github.com/pavlovic265/265-gt/commit/b101202ef9b014fae8ff4822c7515be57583966c))
* trigger release-please ([1d75fde](https://github.com/pavlovic265/265-gt/commit/1d75fde33da1c10dced480400396f164d4885ca7))

## [0.70.0](https://github.com/pavlovic265/265-gt/compare/v0.69.0...v0.70.0) (2026-02-10)


### Features

* show PR approval status instead of merge conflict status in pr list ([21b3103](https://github.com/pavlovic265/265-gt/commit/21b310377e8b386b11be72349559abcd1d7e45e9))


### Bug Fixes

* wrap long function signature to satisfy lll linter' ([3bb4100](https://github.com/pavlovic265/265-gt/commit/3bb41000cbc19ccbccd3dd4ff75357b913895c2e))

## [0.69.0](https://github.com/pavlovic265/265-gt/compare/v0.68.0...v0.69.0) (2026-02-08)


### Features

* add ([97f91c0](https://github.com/pavlovic265/265-gt/commit/97f91c037bd6c8de318c74c3c3bfc3a66b7dec3e))
* add ([0350283](https://github.com/pavlovic265/265-gt/commit/0350283fe2a2854e6ae4116d043b58e42059f162))
* add automated releases with release-please ([f4aa369](https://github.com/pavlovic265/265-gt/commit/f4aa36933a38b1f8fe6a32f038b7f69fbcdbe008))
* add platform field to public.yml ([3a9dcb0](https://github.com/pavlovic265/265-gt/commit/3a9dcb0dd49130f7ffdcb7456cf8bf9c48e94dbc))
* add SSH config integration and clone command ([75bd3af](https://github.com/pavlovic265/265-gt/commit/75bd3af145591da95df7df4076c259bf201c7561))
* auto-attach account after clone ([36a10e4](https://github.com/pavlovic265/265-gt/commit/36a10e4df2f872350c4d10dc9f8ca2cdd74bb5f6))
* move global config to ~/.config/gt/config.yml ([1c900c5](https://github.com/pavlovic265/265-gt/commit/1c900c50abf802b79ecf299c399567829d8892e1))


### Bug Fixes

* simplify TrimPrefix usage in GetRemoteBranches ([750fae5](https://github.com/pavlovic265/265-gt/commit/750fae595878153e508a4b4ea8baf44eed7e9dac))
* standardize error handling patterns ([986fd04](https://github.com/pavlovic265/265-gt/commit/986fd0436634646bd2b93c1f619fb466c0267d1e))
* trigger release on GitHub Release publish ([b2c8b30](https://github.com/pavlovic265/265-gt/commit/b2c8b30110765e462718a812fd2333ad3ba59747))
* trigger release-please ([179b36a](https://github.com/pavlovic265/265-gt/commit/179b36a0b2474a940d028cb6ab67f56fcfc116ce))
* trigger release-please ([51ab51b](https://github.com/pavlovic265/265-gt/commit/51ab51ba2d48dd388e9434d133dfdd20e367464c))
* trigger release-please ([b21cbb5](https://github.com/pavlovic265/265-gt/commit/b21cbb55d644bdc2070a28c3684394090fbd6562))
* use authenticated GitHub API for version check ([02d29b6](https://github.com/pavlovic265/265-gt/commit/02d29b6aa021f4560a878c6b865cbc859434c6fe))
* use PAT for release-please to trigger other workflows ([28893ac](https://github.com/pavlovic265/265-gt/commit/28893ac35cf8ac89dc3696cafd9790d77f40cd65))
* use PAT for release-please workflow ([9bda931](https://github.com/pavlovic265/265-gt/commit/9bda93178fb8a3d201768c253178a3a4d0ff60b8))

## [0.68.0](https://github.com/pavlovic265/265-gt/compare/v0.67.0...v0.68.0) (2026-02-08)


### Features

* add platform field to public.yml ([3a9dcb0](https://github.com/pavlovic265/265-gt/commit/3a9dcb0dd49130f7ffdcb7456cf8bf9c48e94dbc))

## [0.67.0](https://github.com/pavlovic265/265-gt/compare/v0.66.0...v0.67.0) (2026-02-07)


### Features

* add ([97f91c0](https://github.com/pavlovic265/265-gt/commit/97f91c037bd6c8de318c74c3c3bfc3a66b7dec3e))
* add ([0350283](https://github.com/pavlovic265/265-gt/commit/0350283fe2a2854e6ae4116d043b58e42059f162))
* add automated releases with release-please ([f4aa369](https://github.com/pavlovic265/265-gt/commit/f4aa36933a38b1f8fe6a32f038b7f69fbcdbe008))
* add SSH config integration and clone command ([75bd3af](https://github.com/pavlovic265/265-gt/commit/75bd3af145591da95df7df4076c259bf201c7561))
* auto-attach account after clone ([36a10e4](https://github.com/pavlovic265/265-gt/commit/36a10e4df2f872350c4d10dc9f8ca2cdd74bb5f6))
* move global config to ~/.config/gt/config.yml ([1c900c5](https://github.com/pavlovic265/265-gt/commit/1c900c50abf802b79ecf299c399567829d8892e1))


### Bug Fixes

* simplify TrimPrefix usage in GetRemoteBranches ([750fae5](https://github.com/pavlovic265/265-gt/commit/750fae595878153e508a4b4ea8baf44eed7e9dac))
* standardize error handling patterns ([986fd04](https://github.com/pavlovic265/265-gt/commit/986fd0436634646bd2b93c1f619fb466c0267d1e))
* trigger release on GitHub Release publish ([b2c8b30](https://github.com/pavlovic265/265-gt/commit/b2c8b30110765e462718a812fd2333ad3ba59747))
* trigger release-please ([51ab51b](https://github.com/pavlovic265/265-gt/commit/51ab51ba2d48dd388e9434d133dfdd20e367464c))
* trigger release-please ([b21cbb5](https://github.com/pavlovic265/265-gt/commit/b21cbb55d644bdc2070a28c3684394090fbd6562))
* use authenticated GitHub API for version check ([02d29b6](https://github.com/pavlovic265/265-gt/commit/02d29b6aa021f4560a878c6b865cbc859434c6fe))
* use PAT for release-please to trigger other workflows ([28893ac](https://github.com/pavlovic265/265-gt/commit/28893ac35cf8ac89dc3696cafd9790d77f40cd65))
* use PAT for release-please workflow ([9bda931](https://github.com/pavlovic265/265-gt/commit/9bda93178fb8a3d201768c253178a3a4d0ff60b8))

## [0.66.0](https://github.com/pavlovic265/265-gt/compare/v0.65.0...v0.66.0) (2026-02-07)


### Features

* add ([97f91c0](https://github.com/pavlovic265/265-gt/commit/97f91c037bd6c8de318c74c3c3bfc3a66b7dec3e))
* add ([0350283](https://github.com/pavlovic265/265-gt/commit/0350283fe2a2854e6ae4116d043b58e42059f162))
* move global config to ~/.config/gt/config.yml ([1c900c5](https://github.com/pavlovic265/265-gt/commit/1c900c50abf802b79ecf299c399567829d8892e1))


### Bug Fixes

* simplify TrimPrefix usage in GetRemoteBranches ([750fae5](https://github.com/pavlovic265/265-gt/commit/750fae595878153e508a4b4ea8baf44eed7e9dac))
* standardize error handling patterns ([986fd04](https://github.com/pavlovic265/265-gt/commit/986fd0436634646bd2b93c1f619fb466c0267d1e))

## [0.65.0](https://github.com/pavlovic265/265-gt/compare/v0.64.0...v0.65.0) (2026-02-06)


### Features

* add automated releases with release-please ([f4aa369](https://github.com/pavlovic265/265-gt/commit/f4aa36933a38b1f8fe6a32f038b7f69fbcdbe008))
* add SSH config integration and clone command ([75bd3af](https://github.com/pavlovic265/265-gt/commit/75bd3af145591da95df7df4076c259bf201c7561))
* auto-attach account after clone ([36a10e4](https://github.com/pavlovic265/265-gt/commit/36a10e4df2f872350c4d10dc9f8ca2cdd74bb5f6))


### Bug Fixes

* trigger release on GitHub Release publish ([b2c8b30](https://github.com/pavlovic265/265-gt/commit/b2c8b30110765e462718a812fd2333ad3ba59747))
* trigger release-please ([b21cbb5](https://github.com/pavlovic265/265-gt/commit/b21cbb55d644bdc2070a28c3684394090fbd6562))
* use authenticated GitHub API for version check ([02d29b6](https://github.com/pavlovic265/265-gt/commit/02d29b6aa021f4560a878c6b865cbc859434c6fe))
* use PAT for release-please to trigger other workflows ([28893ac](https://github.com/pavlovic265/265-gt/commit/28893ac35cf8ac89dc3696cafd9790d77f40cd65))
* use PAT for release-please workflow ([9bda931](https://github.com/pavlovic265/265-gt/commit/9bda93178fb8a3d201768c253178a3a4d0ff60b8))

## [0.64.0](https://github.com/pavlovic265/265-gt/compare/v0.63.0...v0.64.0) (2026-02-06)


### Features

* add automated releases with release-please ([f4aa369](https://github.com/pavlovic265/265-gt/commit/f4aa36933a38b1f8fe6a32f038b7f69fbcdbe008))


### Bug Fixes

* use PAT for release-please to trigger other workflows ([28893ac](https://github.com/pavlovic265/265-gt/commit/28893ac35cf8ac89dc3696cafd9790d77f40cd65))
* use PAT for release-please workflow ([9bda931](https://github.com/pavlovic265/265-gt/commit/9bda93178fb8a3d201768c253178a3a4d0ff60b8))
