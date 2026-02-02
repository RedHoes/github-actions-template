# GitHub Actions Best Practices - Working Examples

This folder contains real, working GitHub Actions workflows that demonstrate best practices for CI/CD pipelines. Each
workflow is self-contained and can be used as a reference or starting point for your own workflows.

## 📋 Table of Contents

- [Overview](#overview)
- [Workflow Examples](#workflow-examples)
    - [Matrix Strategies](#matrix-strategies)
    - [Caching Strategies](#caching-strategies)
    - [Workflow Optimization](#workflow-optimization)
    - [Security Practices](#security-practices)
    - [Error Handling](#error-handling)
    - [Performance Optimization](#performance-optimization)
- [Quick Reference](#quick-reference)
- [Additional Resources](#additional-resources)

## Overview

These workflows demonstrate practical implementations of the best practices documented
in [BEST_PRACTICES.md](BEST_PRACTICES.md). Each workflow is:

- **Standalone**: Self-contained and doesn't require external composite actions
- **Well-documented**: Extensive inline comments explain each best practice
- **Real-world ready**: Uses realistic scenarios and can be adapted for actual use
- **Safe**: Deployment workflows use `workflow_dispatch` to prevent accidental runs

## Workflow Examples

### Matrix Strategies

#### 1. `matrix-multi-environment.yml`

**What it demonstrates:**

- Deploying to multiple environments in parallel using matrix strategies
- Building once and deploying to multiple targets
- Using `fail-fast: false` to continue deploying even if one environment fails
- Always sending notifications regardless of success/failure

**When to use:**

- You need to deploy the same artifact to multiple production environments
- You want to deploy to all environments simultaneously (faster than sequential)
- You want deployments to continue even if one environment fails

**How to use:**

1. Copy the workflow to your `.github/workflows/` directory
2. Update environment names in the matrix (e.g., `[prod-1, prod-2, prod-3]`)
3. Configure AWS credentials and deployment steps
4. Trigger manually via `workflow_dispatch` or adapt triggers as needed

**Key features:**

- Builds image once in `prepare-image` job
- Deploys to multiple environments in parallel using matrix
- Each environment deployment is independent (fail-fast: false)
- Notifications sent even on failure (if: always())

#### 2. `matrix-multi-version-test.yml`

**What it demonstrates:**

- Testing against multiple versions of dependencies (Python, Node.js)
- Combining multiple matrix dimensions
- Excluding specific combinations
- Parallel test execution for faster feedback

**When to use:**

- You need to ensure compatibility across multiple dependency versions
- You want to test your application against different runtime versions
- You want comprehensive test coverage across versions

**How to use:**

1. Copy the workflow to your `.github/workflows/` directory
2. Update version matrices to match your requirements
3. Configure test commands for your project
4. The workflow runs automatically on pull requests

**Key features:**

- Tests Python against versions 3.9, 3.10, 3.11, 3.12
- Tests Node.js against versions 18, 20, 22
- Multi-dimensional matrix (Python × OS)
- Exclude/include patterns for fine-grained control

### Caching Strategies

#### 3. `caching-strategies.yml`

**What it demonstrates:**

- Built-in caching for Node.js (yarn/npm), Python (pip), and Java (Gradle)
- Manual caching with `actions/cache@v4` for custom paths
- Various cache key strategies (hash-based, branch-based, SHA-based)
- Restore keys for partial cache matches
- Cache scope control (read-only for PRs)

**When to use:**

- You want to speed up workflow execution by caching dependencies
- You have custom build artifacts that should be cached
- You need fine-grained control over cache invalidation

**How to use:**

1. Review the examples in the workflow
2. Copy relevant caching patterns to your workflows
3. Adapt cache keys and paths to your project structure
4. Test cache effectiveness by running workflows multiple times

**Key features:**

- 10 different caching examples
- Built-in caching for common package managers
- Manual caching for Docker layers and custom paths
- Time-based cache invalidation strategies
- Cache read-only mode for PR branches

### Workflow Optimization

#### 4. `workflow-optimization.yml`

**What it demonstrates:**

- Parallel job execution (lint and test run simultaneously)
- Job dependencies (build waits for lint and test)
- Artifact sharing between jobs
- Concurrency control (cancel in-progress runs)
- Conditional steps and jobs

**When to use:**

- You want to maximize parallelization in your workflows
- You need to share build artifacts between jobs
- You want to prevent multiple workflow runs from interfering

**How to use:**

1. Copy the workflow structure to your project
2. Adapt job names and steps to your build process
3. Configure artifact paths and retention
4. Set up concurrency groups appropriately

**Key features:**

- Lint and test jobs run in parallel
- Build job waits for both (using `needs: [lint, test]`)
- Artifacts uploaded and downloaded between jobs
- Concurrency control prevents duplicate runs
- Conditional deployment based on branch/tag

### Security Practices

#### 5. `security-practices.yml`

**What it demonstrates:**

- OIDC for AWS authentication (instead of long-lived credentials)
- GitHub App tokens (instead of personal access tokens)
- Minimal permissions (only grant what's needed)
- Input validation
- Proper secrets usage

**When to use:**

- You need to authenticate with AWS or other cloud providers
- You want to follow security best practices
- You need to validate workflow inputs
- You want to minimize security attack surface

**How to use:**

1. Copy security patterns to your workflows
2. Configure OIDC roles in AWS
3. Set up GitHub App tokens
4. Review and minimize permissions
5. Add input validation for workflow_dispatch inputs

**Key features:**

- OIDC authentication examples
- GitHub App token generation
- Minimal permissions at workflow and job level
- Input validation with proper error messages
- Secure secret usage patterns

### Error Handling

#### 6. `error-handling-notifications.yml`

**What it demonstrates:**

- Always send notifications (even on failure) using `if: always()`
- Proper exit codes with `set -e` and `set -o pipefail`
- Job status checking before proceeding
- `continue-on-error` for non-critical steps
- Proper error messages and context

**When to use:**

- You need reliable notifications for workflow status
- You want to handle errors gracefully
- You have non-critical steps that shouldn't fail the job
- You need to clean up resources even on failure

**How to use:**

1. Copy error handling patterns to your workflows
2. Add notification steps with `if: always()`
3. Use proper exit codes in scripts
4. Mark non-critical steps with `continue-on-error`
5. Add cleanup steps that always run

**Key features:**

- Notifications always sent (success or failure)
- Proper bash error handling (`set -e`, `set -o pipefail`)
- Job and step status checking
- Non-critical steps with `continue-on-error`
- Cleanup steps that always run

### Performance Optimization

#### 7. `performance-optimization.yml`

**What it demonstrates:**

- Docker layer caching with buildx
- Minimal checkout depth (`fetch-depth: 1`)
- Path filters to skip workflows on irrelevant changes
- Artifact retention periods
- Parallel test execution with matrix
- Self-hosted runners for resource-intensive builds

**When to use:**

- You want to optimize workflow execution time
- You have large repositories (minimal checkout helps)
- You want to skip workflows when only docs change
- You need to optimize Docker build times
- You have resource-intensive builds

**How to use:**

1. Add path filters to your workflow triggers
2. Set `fetch-depth: 1` for faster checkouts
3. Enable Docker layer caching
4. Set appropriate artifact retention periods
5. Consider self-hosted runners for large builds

**Key features:**

- Minimal checkout depth for faster runs
- Docker layer caching with GitHub Actions cache
- Path filters to skip irrelevant changes
- Artifact retention to save storage
- Parallel test execution
- Conditional execution to skip unnecessary work

## Quick Reference

### Best Practices Summary

| Practice          | Workflow                                                        | Key Benefit                                          |
|-------------------|-----------------------------------------------------------------|------------------------------------------------------|
| Matrix strategies | `matrix-multi-environment.yml`, `matrix-multi-version-test.yml` | Parallel execution, faster deployments               |
| Caching           | `caching-strategies.yml`                                        | Faster builds, reduced network usage                 |
| Parallelization   | `workflow-optimization.yml`                                     | Faster overall workflow execution                    |
| Security          | `security-practices.yml`                                        | Reduced attack surface, better credential management |
| Error handling    | `error-handling-notifications.yml`                              | Better visibility, graceful failures                 |
| Performance       | `performance-optimization.yml`                                  | Faster workflows, lower costs                        |

### Common Patterns

**Always send notifications:**

```yaml
- name: Send notification
  if: always()  # Runs even on failure
  run:
    echo "Status: ${{ job.status }}"
```

**Minimal permissions:**

```yaml
permissions:
  contents: read
  id-token: write  # Required for OIDC
```

**Concurrency control:**

```yaml
concurrency:
  group: ${{ github.workflow }}-${{ github.ref }}
  cancel-in-progress: true
```

**Built-in caching:**

```yaml
- uses: actions/setup-node@v4
  with:
    node-version: '18'
    cache: 'yarn'  # Automatic caching
```

**Matrix strategy:**

```yaml
strategy:
  matrix:
    environment: [ prod-1, prod-2, prod-3 ]
  fail-fast: false
```

## Additional Resources

- **[BEST_PRACTICES.md](BEST_PRACTICES.md)** - Detailed documentation of all best practices
- **[README.md](../README.md)** - Main templates documentation
- **[WORKFLOW_STEPS_INDEX.md](../WORKFLOW_STEPS_INDEX.md)** - Reference for composite actions
- **[VARIABLES_AND_SECRETS.md](../VARIABLES_AND_SECRETS.md)** - Variables and secrets reference
- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [GitHub Actions Best Practices](https://docs.github.com/en/actions/learn-github-actions/best-practices)

## Contributing

When adding new workflow examples:

1. Keep workflows standalone and self-contained
2. Add extensive inline comments explaining best practices
3. Include a header comment describing what the workflow demonstrates
4. Use realistic scenarios that can be adapted for actual use
5. Update this README with the new workflow's description

## Notes

- All workflows are examples and may need adaptation for your specific use case
- Replace placeholder values (e.g., `${{ vars.* }}`, `${{ secrets.* }}`) with your actual values
- Test workflows in a safe environment before using in production
- Review security implications before copying authentication patterns
