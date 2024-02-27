## [Unreleased]

## [0.1.0] - 2024-02-25
### Added
- Capability to pass version information as an argument to the build script. This allows the build process to dynamically include version details in the binary. (./script/build.sh)
- Validation for Git information (commit hash, tree state) to ensure accuracy in version reporting.
- Default information printout if no specific version information is provided during the build process. This ensures that the application always has some version info to display, even in default scenarios.

### Fixed
- Updated the version display logic in /cmd/bean/main.go to include comprehensive version information, including Git commit, tree state, build date, and Go environment details.