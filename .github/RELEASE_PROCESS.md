## Release Process for Go CLI Application with Embedded Version

Here's a release process for your Go CLI application with embedded version information, suitable for both DevOps and development teams:

**Preparation:**

1. **Versioning:**
    * Implement a versioning scheme (e.g., semantic versioning).
    * Ensure the version information can be easily accessed within the code or a extenal git repository. 

2. **Version Embedding:**
    * Use a build flag (`-ldflags`) during compilation to embed the version information into the binary. Here's an example:

    ```
    go build -ldflags=...
    ```

    Tthe running script is in ./scripts/build

3. **Changelog:**
    * Maintain a changelog file (e.g., `CHANGELOG.md`) following the "Keep a Changelog" format ([https://github.com/olivierlacan/keep-a-changelog](https://github.com/olivierlacan/keep-a-changelog)). Here's the basic structure:

    ```
    ## Unreleased

    ### Added
    * Briefly describe the added feature.
    * Additional details if needed.

    ### Changed
    * Briefly describe the change made.
    * Impact of the change.

    ### Fixed
    * Briefly describe the bug fix.
    * Reference any related issue number.

    ### Deprecated
    * Briefly describe the deprecated feature.
    * Recommended alternative or explanation for deprecation.

    ## [Version number] - [Release date]

    ... (Repeat the sections above for each version)
    ```

**Development Team Responsibilities:**

1. **Development and Testing:**
    * Develop new features and bug fixes.
    * Write unit and integration tests for changes made.
    * Update documentation and release notes reflecting changes in the new version.

2. **Pull Requests and Code Review:**
    * Submit pull requests for code changes and new features.
    * Participate in code reviews for other developers' work.

3. **Version Update and Changelog:**
    * Update the version number in the codebase when necessary.
    * **Before merging pull requests:**
        * Update the appropriate section (Added, Changed, Fixed, Deprecated) in the `Unreleased` section of the changelog.
        * Use clear and concise language, following the "Keep a Changelog" format.

**DevOps Team Responsibilities:**

1. **Continuous Integration/Continuous Delivery (CI/CD):**
    * Set up a CI/CD pipeline to automate the build, test, and deployment process upon code changes.
    * Integrate automated testing tools (e.g., Go testing framework) into the pipeline.
    * Configure the pipeline to use the appropriate `-ldflags` during the build stage to embed the version information.

2. **Release Deployment:**
    * Deploy the new version to the production environment using appropriate deployment methods (e.g., package managers, container orchestration tools).
    * Update any configuration files or environment variables needed for the new version.

3. **Monitoring and Rollbacks:**
    * Monitor the application health and performance after deployment.
    * Have a rollback plan in place in case of unforeseen issues with the new version.

