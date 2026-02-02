# GitHub Actions Best Practices

This document covers recommended patterns, optimizations, and best practices for GitHub Actions workflows. Use these
guidelines to improve workflow performance, maintainability, and reliability.

## Table of Contents

- [Matrix Strategies](#matrix-strategies)
- [GitHub Actions Caching](#github-actions-caching)
- [Workflow Optimization Patterns](#workflow-optimization-patterns)
- [Security Best Practices](#security-best-practices)
- [Error Handling and Notifications](#error-handling-and-notifications)
- [Performance Optimization](#performance-optimization)

---

## Matrix Strategies

Matrix strategies allow you to run a job multiple times with different configurations. This is particularly useful for:

- Deploying to multiple environments in parallel
- Testing against multiple versions of dependencies
- Building for multiple platforms or architectures

### Multi-Environment Deployment Pattern

When you need to deploy to multiple production environments (e.g., `prod-1`, `prod-2`, `prod-3`), use a matrix strategy
to deploy in parallel:

```yaml
jobs:
  promote-images:
    outputs:
      IMAGE_TAG: ${{ steps.prepare_image.outputs.app_version }}
    runs-on: ubuntu-latest
    steps:
      # Build or promote image once
      - name: Prepare image
        id: prepare_image
        run: |
          IMAGE_TAG=${GITHUB_SHA::8}
          echo "app_version=${IMAGE_TAG}" >> $GITHUB_OUTPUT

  deploy:
    needs: promote-images
    runs-on: ubuntu-latest
    strategy:
      matrix:
        environment: [ prod-1, prod-2, prod-3 ]
      fail-fast: false  # Continue deploying to other environments even if one fails

    name: Deploy to ${{ matrix.environment }}

    steps:
      - uses: ./.github/actions/setup-github-auth
        with:
          app_id: ${{ vars.GH_APP_ID }}
          private_key: ${{ secrets.GH_APP_PRIVATE_KEY }}

      - uses: ./.github/actions/setup-aws-credentials
        with:
          role_to_assume: ${{ secrets.AWS_ROLE_TO_ASSUME_PROD }}
          aws_region: ${{ vars.AWS_REGION }}

      - uses: ./.github/actions/deploy-ecs
        with:
          aws_region: ${{ vars.AWS_REGION }}
          ecs_cluster_name: ${{ vars.ECS_CLUSTER_NAME_PROD }}-${{ matrix.environment }}
          service_name: ${{ vars.SERVICE_NAME_PROD }}-${{ matrix.environment }}
          task_family: ${{ vars.TASK_FAMILY_PROD }}-${{ matrix.environment }}
          container_name: ${{ vars.CONTAINER_NAME }}
          image_uri: ${{ vars.DOCKER_REPO_PROD }}:${{ needs.promote-images.outputs.IMAGE_TAG }}
```

### Key Matrix Strategy Features

1. **`fail-fast: false`**: By default, if one matrix job fails, all other jobs are cancelled. Set `fail-fast: false` to
   allow other environments to continue deploying even if one fails.

2. **Dynamic Matrix**: You can generate matrix values dynamically:

```yaml
strategy:
  matrix:
    environment: ${{ fromJson(vars.ENVIRONMENTS) }}  # Read from variable
```

3. **Exclude Specific Combinations**: Skip certain combinations:

```yaml
strategy:
  matrix:
    environment: [ prod-1, prod-2, prod-3 ]
    include:
      - environment: staging
    exclude:
      - environment: prod-2  # Skip prod-2
```

### Testing Multiple Versions

Use matrix strategies to test against multiple dependency versions:

```yaml
jobs:
  test:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        python-version: [ '3.9', '3.10', '3.11' ]
        node-version: [ '18', '20' ]
    steps:
      - uses: actions/setup-python@v5
        with:
          python-version: ${{ matrix.python-version }}
      - uses: actions/setup-node@v4
        with:
          node-version: ${{ matrix.node-version }}
```

---

## GitHub Actions Caching

Caching dependencies significantly speeds up workflow execution by reusing previously downloaded dependencies.

### Built-in Caching

Most setup actions have built-in caching support:

#### Node.js/Yarn

```yaml
- uses: actions/setup-node@v4
  with:
    node-version: '18'
    cache: 'yarn'  # Automatically caches yarn dependencies
```

#### Python/pip

```yaml
- uses: actions/setup-python@v5
  with:
    python-version: '3.11'
    cache: 'pip'  # Automatically caches pip dependencies
```

#### Java/Gradle

```yaml
- uses: actions/setup-java@v4
  with:
    java-version: '17'
    distribution: 'temurin'
    cache: 'gradle'  # Automatically caches Gradle dependencies
```

### Manual Caching with actions/cache

For custom caching needs, use `actions/cache@v4`:

```yaml
- name: Cache Docker layers
  uses: actions/cache@v4
  with:
    path: /tmp/.buildx-cache
    key: ${{ runner.os }}-buildx-${{ github.sha }}
    restore-keys: |
      ${{ runner.os }}-buildx-

- name: Build Docker image
  uses: docker/build-push-action@v5
  with:
    cache-from: type=local,src=/tmp/.buildx-cache
    cache-to: type=local,dest=/tmp/.buildx-cache-new,mode=max
```

### Cache Key Strategies

1. **Branch-based caching**: Main branch writes to primary cache, PRs use isolated caches:

```yaml
- name: Cache dependencies
  uses: actions/cache@v4
  with:
    path: |
      ~/.gradle/caches
      ~/.gradle/wrapper
    key: ${{ runner.os }}-gradle-${{ hashFiles('**/*.gradle*', '**/gradle-wrapper.properties') }}
    restore-keys: |
      ${{ runner.os }}-gradle-
```

2. **SHA-based caching**: Cache per commit (useful for reproducible builds):

```yaml
key: ${{ runner.os }}-deps-${{ github.sha }}
```

3. **Time-based invalidation**: Add date to force cache refresh periodically:

```yaml
key: ${{ runner.os }}-deps-${{ hashFiles('**/package-lock.json') }}-${{ github.run_number }}
```

### Cache Best Practices

- **Cache size limits**: GitHub Actions caches are limited to 10GB per repository. Clean up old caches regularly.
- **Cache scope**: Use `cache_read_only: true` for PR branches to prevent cache pollution.
- **Cache paths**: Be specific with cache paths to avoid unnecessary cache hits/misses.
- **Restore keys**: Use restore keys to find partial matches when exact key doesn't exist.

---

## Workflow Optimization Patterns

### 1. Job Dependencies and Parallelization

Structure workflows to maximize parallelization:

```yaml
jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - name: Run linters
        run: npm run lint

  test:
    runs-on: ubuntu-latest
    steps:
      - name: Run tests
        run: npm test

  build:
    needs: [ lint, test ]  # Wait for both to complete
    runs-on: ubuntu-latest
    steps:
      - name: Build
        run: npm run build
```

### 2. Conditional Steps and Jobs

Use conditions to skip unnecessary work:

```yaml
jobs:
  deploy:
    if: github.ref == 'refs/heads/main'  # Only run on main branch
    runs-on: ubuntu-latest
    steps:
      - name: Deploy
        if: github.event_name == 'push'  # Skip on manual triggers
        run: echo "Deploying..."
```

### 3. Reusable Workflows

Extract common workflow patterns into reusable workflows:

```yaml
# .github/workflows/deploy.yml (reusable workflow)
on:
  workflow_call:
    inputs:
      environment:
        required: true
        type: string

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - name: Deploy to ${{ inputs.environment }}
        run: echo "Deploying to ${{ inputs.environment }}"

# .github/workflows/deploy-prod.yml (calling workflow)
on:
  push:
    tags: [ 'v*.*.*' ]

jobs:
  deploy:
    uses: ./.github/workflows/deploy.yml
    with:
      environment: 'prod'
```

### 4. Composite Actions

Use composite actions for reusable step sequences (already used in these templates):

```yaml
# .github/actions/my-action/action.yml
name: 'My Custom Action'
description: 'Does something useful'
inputs:
  my-input:
    description: 'An input'
    required: true
runs:
  using: 'composite'
  steps:
    - run: echo "Doing something with ${{ inputs.my-input }}"
      shell: bash
```

### 5. Concurrency Control

Prevent multiple workflow runs from interfering with each other:

```yaml
concurrency:
  group: ${{ github.workflow }}-${{ github.ref }}
  cancel-in-progress: true  # Cancel older runs when new one starts
```

### 6. Artifact Management

Use artifacts to share files between jobs:

```yaml
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Build
        run: npm run build
      - name: Upload artifacts
        uses: actions/upload-artifact@v4
        with:
          name: dist
          path: dist/

  deploy:
    needs: build
    runs-on: ubuntu-latest
    steps:
      - name: Download artifacts
        uses: actions/download-artifact@v4
        with:
          name: dist
      - name: Deploy
        run: echo "Deploying artifacts"
```

---

## Security Best Practices

### 1. Use OIDC for AWS Authentication

Always use OIDC instead of long-lived credentials:

```yaml
- uses: aws-actions/configure-aws-credentials@v4
  with:
    role-to-assume: ${{ secrets.AWS_ROLE_TO_ASSUME_DEV }}
    role-session-name: gh-actions
    aws-region: ${{ vars.AWS_REGION }}
```

### 2. Use GitHub App Tokens

Use GitHub App tokens instead of personal access tokens:

```yaml
- uses: tibdex/github-app-token@v1
  with:
    app_id: ${{ vars.GH_APP_ID }}
    private_key: ${{ secrets.GH_APP_PRIVATE_KEY }}
```

### 3. Minimize Permissions

Only grant the minimum permissions required:

```yaml
permissions:
  contents: read  # Only read, not write
  id-token: write  # Required for OIDC
```

### 4. Use Secrets for Sensitive Data

Never hardcode secrets. Always use GitHub secrets:

```yaml
env:
  API_KEY: ${{ secrets.API_KEY }}  # ✅ Good
  # API_KEY: "hardcoded-value"     # ❌ Bad
```

### 5. Validate Inputs

Validate workflow inputs and environment variables:

```yaml
jobs:
  deploy:
    if: vars.ENVIRONMENT == 'dev' || vars.ENVIRONMENT == 'prod'
    steps:
      - name: Validate environment
        run: |
          if [[ ! "${{ vars.ENVIRONMENT }}" =~ ^(dev|prod)$ ]]; then
            echo "Invalid environment"
            exit 1
          fi
```

---

## Error Handling and Notifications

### 1. Always Send Notifications

Use `if: always()` to send notifications even when jobs fail:

```yaml
- name: Send notification
  if: always()  # Run even if previous steps failed
  uses: ./.github/actions/send-slack-notification
  with:
    status: ${{ job.status }}
    service_name: ${{ vars.SERVICE_NAME }}
    environment: ${{ vars.ENVIRONMENT }}
```

### 2. Fail Fast for Critical Steps

Use `continue-on-error` sparingly and only for non-critical steps:

```yaml
- name: Optional step
  continue-on-error: true
  run: echo "This won't fail the job"
```

### 3. Proper Exit Codes

Ensure scripts return proper exit codes:

```bash
#!/bin/bash
set -e  # Exit on error
set -o pipefail  # Exit on pipe failure

# Your script here
```

### 4. Job Status Aggregation

Check overall job status before proceeding:

```yaml
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - name: Run tests
        run: npm test

  deploy:
    needs: test
    if: needs.test.result == 'success'
    runs-on: ubuntu-latest
    steps:
      - name: Deploy
        run: echo "Deploying"
```

---

## Performance Optimization

### 1. Use Self-Hosted Runners for Large Builds

For resource-intensive builds, use self-hosted runners:

```yaml
jobs:
  build:
    runs-on: builder-16-32  # Self-hosted runner with 16 cores, 32GB RAM
    steps:
      - name: Build
        run: npm run build
```

### 2. Optimize Docker Builds

Use Docker layer caching and buildx:

```yaml
- name: Set up Docker Buildx
  uses: docker/setup-buildx-action@v3

- name: Build and push
  uses: docker/build-push-action@v5
  with:
    cache-from: type=gha
    cache-to: type=gha,mode=max
```

### 3. Minimize Checkout Depth

Only checkout the minimum required history:

```yaml
- uses: actions/checkout@v4
  with:
    fetch-depth: 1  # Only latest commit, not full history
```

### 4. Parallel Test Execution

Run tests in parallel when possible:

```yaml
jobs:
  test:
    strategy:
      matrix:
        test-suite: [ unit, integration, e2e ]
    runs-on: ubuntu-latest
    steps:
      - name: Run ${{ matrix.test-suite }} tests
        run: npm run test:${{ matrix.test-suite }}
```

### 5. Skip Unnecessary Steps

Use path filters to skip workflows when irrelevant files change:

```yaml
on:
  push:
    paths:
      - 'src/**'
      - 'package.json'
      - '.github/workflows/**'
```

### 6. Use Workflow Artifacts Efficiently

Only upload necessary files and set retention periods:

```yaml
- uses: actions/upload-artifact@v4
  with:
    name: test-results
    path: test-results/
    retention-days: 7  # Auto-delete after 7 days
```

---

## Additional Resources

- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [GitHub Actions Best Practices](https://docs.github.com/en/actions/learn-github-actions/best-practices)
- [Caching dependencies](https://docs.github.com/en/actions/using-workflows/caching-dependencies-to-speed-up-workflows)
- [Matrix strategies](https://docs.github.com/en/actions/using-jobs/using-a-matrix-for-your-jobs)

---

## Working Examples

For real, working GitHub Actions workflows that demonstrate these best practices, see the *
*[best-practices](./best-practices/)** folder. This folder contains:

- **Standalone workflow examples** for each best practice category
- **Well-documented code** with extensive inline comments
- **Real-world scenarios** that can be adapted for your projects
- **Comprehensive README** explaining each workflow's purpose and usage

Each workflow is self-contained and demonstrates specific patterns you can copy and adapt for your own CI/CD pipelines.

---

## Summary

Key takeaways:

1. **Use matrix strategies** for parallel deployments to multiple environments
2. **Enable caching** for all dependency management (npm, pip, Gradle, etc.)
3. **Structure workflows** to maximize parallelization
4. **Use OIDC and GitHub App tokens** for secure authentication
5. **Optimize performance** with self-hosted runners, Docker caching, and minimal checkouts
6. **Handle errors gracefully** with proper notifications and status checks

For workflow templates and examples, see [README.md](../README.md).  
For working examples of these best practices, see [best-practices/](./best-practices/).
