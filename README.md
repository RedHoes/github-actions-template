# GitHub Actions Workflow Templates

> **👋 New to setting up CI/CD?** Start here! This guide will help you set up GitHub Actions workflows for your project
> in minutes.

## 📁 Repository Structure

This repository contains:

- **`templates/`** - Reusable workflow templates that use composite actions
    - These are optimized for reusability and maintainability
    - Use composite actions from `.github/actions/` for modularity
    - Copy these to your project and customize as needed

- **`projects/`** - Sample test projects with fully expanded workflows
    - `backend-kotlin/` - Kotlin/Gradle test project with detailed workflows
    - `backend-python/` - Python/Flask test project with detailed workflows
    - `frontend/` - React/Vite test project with CloudFront workflows
    - `infra/` - Terraform test project with deployment workflows
    - **Note**: These workflows have all steps written out explicitly (no composite action references)
    - Use these as reference examples to understand workflow structure

- **`.github/actions/`** - Reusable composite actions (at repository root)
    - Used by templates for modularity
    - Can be referenced in your own workflows if needed

## 🧪 Sample Projects

The `projects/` directory contains ready-to-use test codebases with **fully expanded workflows**. Each project includes:

- Source code and build configuration
- Dockerfiles (for backend projects)
- Pre-configured GitHub Actions workflows with **all steps written out explicitly**
- No composite action references - easier to understand and modify
- Ready to test CI/CD pipelines

**Key Difference from Templates:**

- **Templates** (`templates/`): Use composite actions for reusability
- **Sample Projects** (`projects/`): Have all steps expanded for clarity

To test workflows:

1. Navigate to a project: `cd projects/backend-kotlin`
2. Initialize Git: `git init && git add . && git commit -m "Initial commit"`
3. Push to GitHub and configure secrets/variables
4. Test: Create PR → triggers PR checks, Push to main → triggers DEV deployment

**When to use which:**

- **Use templates** if you want modular, maintainable workflows that reference composite actions
- **Use sample projects** as reference to understand the detailed steps, then copy and customize for your needs

---

## 🚀 Quick Start: I'm Starting a New Project

**First, answer these questions:**

1. **What type of project is this?**
    - [ ] Frontend (React, Vue, Angular, etc.)
    - [ ] Backend API (Kotlin/Java or Python)
    - [ ] Infrastructure (Terraform)

2. **For Frontend/Backend: What's your deployment strategy?**
    - [ ] **Default**: Deploy when code is pushed to `main` or `master` branch

   > **Note**: If you need a trunk-based release strategy, consider using a release branch (e.g., `release/v1.0.0`)
   instead. You can modify the workflow trigger to include `release/*` branches in the `on.push.branches` section.

3. **For Backend: What's your deployment target?**
    - [ ] **EKS** (Kubernetes) - Use if you have Kubernetes clusters
    - [ ] **ECS** (AWS Container Service) - Use if you prefer simpler AWS-native deployments

4. **For Frontend: What's your deployment method?**
    - [ ] **CloudFront/S3** - Static site hosting on AWS
    - [ ] **Amplify** - AWS Amplify hosting

---

## 📋 Step-by-Step Setup Guide

### Step 1: Choose Your Templates

Based on your answers above, find your template path below and copy the files.

#### For Frontend Projects

**CloudFront Deployment:**

```bash
# 1. Copy PR check workflow (always needed)
cp templates/frontend/cloudfront/pr-check.yml <your-project>/.github/workflows/

# 2. Copy deployment workflows
cp templates/frontend/cloudfront/deploy-*.yml <your-project>/.github/workflows/

# 3. Optional: Copy E2E tests
cp templates/frontend/e2e-tests.yml <your-project>/.github/workflows/
```

**Amplify Deployment:**

```bash
# 1. Copy PR check workflow
cp templates/frontend/amplify/pr-check.yml <your-project>/.github/workflows/

# 2. Copy deployment workflows
cp templates/frontend/amplify/deploy-*.yml <your-project>/.github/workflows/
```

#### For Backend Projects

**Kotlin/Java Projects:**

```bash
# 1. Always copy these base workflows
cp templates/backend/kotlin/pr-check.yml <your-project>/.github/workflows/

# 2. Choose your deployment method:

# Option A: EKS/Helm
cp templates/backend/kotlin/eks/deploy-*.yml <your-project>/.github/workflows/

# Option B: ECS
cp templates/backend/kotlin/ecs/deploy-*.yml <your-project>/.github/workflows/
```

**Python Projects:**

```bash
# 1. Always copy these base workflows
cp templates/backend/python/pr-check.yml <your-project>/.github/workflows/

# 2. Choose your deployment method (same options as Kotlin above, but use python/ instead of kotlin/)
# Example for EKS:
cp templates/backend/python/eks/deploy-*.yml <your-project>/.github/workflows/
```

#### For Infrastructure (Terraform) Projects

```bash
# 1. Copy PR check workflow
cp templates/infra/pr-check.yml <your-project>/.github/workflows/

# 2. Copy deployment workflows
cp templates/infra/deploy-*.yml <your-project>/.github/workflows/
```

---

### Step 2: Customize Your Workflows

After copying the templates, open each workflow file and look for `TODO:` comments. Update these values:

#### Frontend Customization

In each workflow file, update:

- [ ] **Node.js version**: Find `NODE_VERSION` and set your version (e.g., `18.18.2`)
- [ ] **Environment variables**: Add your `VITE_*` variables in the `env:` section
- [ ] **Package manager**: If not using `yarn`, update cache and install commands

#### Backend Customization

**For Kotlin/Java:**

- [ ] **Java version**: Find `JAVA_VERSION` (usually `17`)
- [ ] **Gradle build commands**: Update build steps if your project structure differs
- [ ] **Docker image paths**: Update `context:` and `file:` paths in Docker build steps
- [ ] **Service name**: Update `SERVICE_NAME` variable

**For Python:**

- [ ] **Python version**: Find `PYTHON_VERSION` (usually `3.11`)
- [ ] **Dependencies**: Uncomment and configure pip or Poetry steps
- [ ] **Docker image paths**: Update `context:` and `file:` paths
- [ ] **Service name**: Update `SERVICE_NAME` variable

**For All Backend:**

- [ ] **Helm configuration**: Update `HELM_REPO`, `HELM_SPARTAN_VERSION` if using EKS
- [ ] **ECS configuration**: Update `ECS_CLUSTER_NAME_*`, `TASK_FAMILY_*`, `CONTAINER_NAME` if using ECS

#### Infrastructure Customization

- [ ] **Terraform version**: Update `TF_VERSION` if specified
- [ ] **Working directory**: Update if Terraform files are in a subdirectory
- [ ] **Backend configuration**: Set up `TF_BACKEND_CONFIG_*` secrets if needed

---

### Step 3: Configure GitHub Variables and Secrets

Go to your GitHub repository → **Settings** → **Secrets and variables** → **Actions**

#### Required for All Projects

**Variables:**

- [ ] `GH_APP_ID` - Get from your organization's GitHub App
- [ ] `AWS_REGION` - Your AWS region (e.g., `us-east-1`)

**Secrets:**

- [ ] `GH_APP_PRIVATE_KEY` - Get from your organization's GitHub App
- [ ] `SLACK_WEBHOOK_URL` - Your Slack webhook for notifications
- [ ] `AWS_ROLE_TO_ASSUME_DEV` - AWS IAM role ARN for DEV environment
- [ ] `AWS_ROLE_TO_ASSUME_PROD` - AWS IAM role ARN for PROD environment

#### Frontend-Specific

**For CloudFront:**

- [ ] `AWS_S3_BUCKET_DEV` (variable)
- [ ] `AWS_CLOUDFRONT_ID_DEV` (variable)
- [ ] `AWS_S3_BUCKET_PROD` (variable)
- [ ] `AWS_CLOUDFRONT_ID_PROD` (variable)
- [ ] `NODE_VERSION` (variable)
- [ ] `VITE_*` variables (your frontend environment variables)

**For Amplify:**

- [ ] `AMPLIFY_WEBHOOK_URL_DEV` (secret)
- [ ] `AMPLIFY_WEBHOOK_URL_PROD` (secret)
- [ ] `NODE_VERSION` (variable)

#### Backend-Specific

**For Kotlin/Java:**

- [ ] `ARTIFACTORY_URL` (variable)
- [ ] `ARTIFACTORY_USERNAME` (variable)
- [ ] `AWS_CODE_ARTIFACT_DOMAIN_DEV` (variable)
- [ ] `AWS_ACCOUNT_ID_DEV` (secret)
- [ ] `AWS_ACCOUNT_ID_PROD` (secret)

**For EKS:**

- [ ] `HELM_REPO` (variable)
- [ ] `HELM_SPARTAN_VERSION` (variable)
- [ ] `DOCKER_REPO_DEV` (variable)
- [ ] `DOCKER_REPO_PROD` (variable)
- [ ] `SERVICE_NAME` (variable)

**For ECS:**

- [ ] `ECS_CLUSTER_NAME_DEV` (variable)
- [ ] `ECS_CLUSTER_NAME_PROD` (variable)
- [ ] `TASK_FAMILY_DEV` (variable)
- [ ] `TASK_FAMILY_PROD` (variable)
- [ ] `CONTAINER_NAME` (variable)
- [ ] `SERVICE_NAME_DEV` (variable)
- [ ] `SERVICE_NAME_PROD` (variable)
- [ ] `DOCKER_REPO` (variable)

#### Infrastructure-Specific

- [ ] `TF_BACKEND_CONFIG_DEV` (secret, optional)
- [ ] `TF_BACKEND_CONFIG_PROD` (secret, optional)
- [ ] Project-specific `TF_VAR_*` variables

> **📋 Complete Reference**: See the [Variables and Secrets](#variables-and-secrets-reference) section below for the full
> list.

---

### Step 4: Test Your Setup

1. **Create a test branch:**
   ```bash
   git checkout -b test-ci-setup
   git push origin test-ci-setup
   ```

2. **Create a Pull Request** - This will trigger the `pr-check.yml` workflow

3. **Check the Actions tab** in GitHub to see if workflows run successfully

4. **Fix any errors** - Common issues:
    - Missing variables/secrets
    - Incorrect file paths in Docker builds
    - Wrong service names or cluster names

---

## 🎯 What Each Workflow Does

### Frontend Workflows

- **`pr-check.yml`**: Runs on every Pull Request
    - Lints your code
    - Runs tests
    - Builds your app (to catch build errors early)

- **`deploy-dev-*.yml`**: Deploys to DEV environment
    - Triggers: Push to `main`/`master` branch
    - Builds your app
    - Deploys to DEV (CloudFront S3 bucket or Amplify)

- **`deploy-prod-*.yml`**: Deploys to PROD environment
    - Triggers: When you create a git tag (e.g., `v1.0.0`)
    - Builds your app
    - Deploys to PROD

- **`e2e-tests.yml`**: Optional end-to-end tests
    - Runs Playwright tests
    - Can be triggered manually or on schedule

### Backend Workflows

- **`pr-check.yml`**: Runs on every Pull Request
    - Builds your code
    - Runs tests
    - Runs database migrations (Flyway for Kotlin)

- **`deploy-dev.yml`**: Deploys to DEV
    - Builds Docker image
    - Pushes to ECR
    - Deploys to DEV environment (EKS/ECS)

- **`deploy-prod.yml`**: Deploys to PROD
    - Promotes DEV image to PROD (no rebuild)
    - Deploys to PROD environment

### Infrastructure Workflows

- **`pr-check.yml`**: Runs on Pull Requests
    - Validates Terraform syntax
    - Runs `terraform plan` (shows what would change)

- **`deploy-dev.yml`**: Applies Terraform to DEV
    - Triggers: Push to `main`/`master`
    - Runs `terraform apply`

- **`deploy-prod.yml`**: Applies Terraform to PROD
    - Triggers: Git tag `v*.*.*`
    - Runs `terraform apply`

---

## 🔧 Reusable Workflow Steps (Composite Actions)

The **templates** use **composite actions** - reusable, modular steps that make workflows easier to understand and
maintain.
These actions are located in `.github/actions/` and can be referenced in any workflow.

**Note:** The sample projects in `projects/` have all steps expanded explicitly for easier understanding and
modification.
If you prefer detailed, explicit workflows, use the sample projects as a reference.

### Architecture: 3-Step Structure

Modern workflows use a simplified **3-step structure** for maximum clarity:

1. **Setup Environment** - Authentication, AWS credentials, and language setup
2. **Build and Package** - Build code and create Docker images (or build frontend bundle)
3. **Deploy and Notify** - Deploy to target and send notifications

### High-Level Consolidated Actions

These actions combine multiple operations for simplified workflows:

#### `setup-environment`

Combines GitHub auth, AWS credentials, and language setup (Java/Kotlin or Python).

**Inputs:**

- `app_id`, `private_key` - GitHub App authentication
- `aws_role_to_assume`, `aws_region` - AWS configuration
- `language` - `java-kotlin` or `python`
- Language-specific inputs (Java version, Python version, artifact registry config, etc.)

**Used in:** Backend deployment workflows

#### `setup-frontend-environment`

Combines GitHub auth, AWS credentials, and Node.js setup.

**Inputs:**

- `app_id`, `private_key` - GitHub App authentication
- `aws_role_to_assume`, `aws_region` - AWS configuration
- `node_version`, `package_manager` - Node.js configuration

**Used in:** Frontend deployment workflows

#### `build-and-package`

Combines build operations and Docker image creation/push.

**Inputs:**

- `language` - `java-kotlin` or `python`
- `docker_repo`, `environment`, `github_sha` - Docker configuration
- `docker_context`, `dockerfile` - Docker build configuration
- Language-specific build inputs

**Outputs:**

- `image_tag` - Generated Docker image tag

**Used in:** Backend deployment workflows

#### `build-frontend-bundle`

Builds frontend application bundle.

**Inputs:**

- `build_command`, `package_manager`

**Used in:** Frontend deployment workflows

#### `build-and-test`

Combines environment setup, build, lint, and test operations.

**Inputs:**

- `app_id`, `private_key` - GitHub App authentication
- `language` - `java-kotlin`, `python`, or `nodejs`
- `test_command`, `lint_command` - Test and lint commands
- `run_linters`, `run_build` - Feature flags
- Language-specific configuration

**Used in:** PR check workflows

#### `deploy-and-notify`

Combines deployment and notification operations. Supports multiple deployment types.

**Inputs:**

- `deployment_type` - `ecs`, `eks`, `cloudfront`, `amplify`, or `terraform`
- `aws_region`, `service_name`, `environment` - Common deployment config
- `slack_webhook_url` - Notification configuration
- Type-specific inputs (ECS cluster, Helm config, S3 bucket, etc.)

**Used in:** All deployment workflows

### Individual Actions (For Advanced Use Cases)

These actions are still available for workflows that need more granular control:

#### Authentication & Setup

- `setup-github-auth` - Generate GitHub App token and checkout source
- `setup-aws-credentials` - Configure AWS credentials using OIDC
- `setup-java-kotlin` - Setup JDK, Gradle, and artifact registry
- `setup-python` - Setup Python version and install dependencies

#### Build Steps

- `build-java-kotlin` - Run Gradle build commands

#### Docker Operations

- `docker-ecr-login` - Login to Amazon ECR
- `docker-prepare-tags` - Prepare Docker image tags from SHA/version
- `docker-build-push` - Build and push Docker image to ECR

#### Utilities

- `extract-git-info` - Extract git author, revision, and commit messages
- `extract-app-version` - Extract application version from manifest.json, pyproject.toml, or git tag
    - **Inputs:** `version_source` (manifest, pyproject, or git-tag), `version_file`
    - **Outputs:** `app_version`
    - **Used in:** Version extraction scenarios
- `send-slack-notification` - Send deployment notification to Slack

### Example: Using Consolidated Actions (Templates)

**Backend Deployment (3 steps) - Used in templates:**

```yaml
steps:
  # Step 1: Setup Environment
  - uses: ./.github/actions/setup-environment
    with:
      app_id: ${{ vars.GH_APP_ID }}
      private_key: ${{ secrets.GH_APP_PRIVATE_KEY }}
      aws_role_to_assume: ${{ secrets.AWS_ROLE_TO_ASSUME_DEV }}
      aws_region: ${{ vars.AWS_REGION }}
      language: 'java-kotlin'
      java_version: '17'

  # Step 2: Build and Package
  - uses: ./.github/actions/build-and-package
    id: package
    with:
      language: 'java-kotlin'
      docker_repo: ${{ vars.DOCKER_REPO }}
      environment: 'dev'
      aws_region: ${{ vars.AWS_REGION }}

  # Step 3: Deploy and Notify
  - uses: ./.github/actions/deploy-and-notify
    with:
      deployment_type: 'ecs'
      aws_region: ${{ vars.AWS_REGION }}
      service_name: ${{ vars.SERVICE_NAME }}
      environment: 'dev'
      ecs_cluster_name: ${{ vars.ECS_CLUSTER_NAME }}
      task_family: ${{ vars.TASK_FAMILY }}
      container_name: ${{ vars.CONTAINER_NAME }}
      image_uri: ${{ vars.DOCKER_REPO }}:${{ steps.package.outputs.image_tag }}
      slack_webhook_url: ${{ secrets.SLACK_WEBHOOK_URL }}
```

**Frontend Deployment (3 steps) - Used in templates:**

```yaml
steps:
  # Step 1: Setup Environment
  - uses: ./.github/actions/setup-frontend-environment
    with:
      app_id: ${{ vars.GH_APP_ID }}
      private_key: ${{ secrets.GH_APP_PRIVATE_KEY }}
      aws_role_to_assume: ${{ secrets.AWS_ROLE_TO_ASSUME_DEV }}
      aws_region: ${{ vars.AWS_REGION }}
      node_version: '18.18.2'

  # Step 2: Build Bundle
  - uses: ./.github/actions/build-frontend-bundle

  # Step 3: Deploy and Notify
  - uses: ./.github/actions/deploy-and-notify
    with:
      deployment_type: 'cloudfront'
      aws_region: ${{ vars.AWS_REGION }}
      service_name: ${{ github.event.repository.name }}
      environment: 'dev'
      s3_bucket: ${{ vars.AWS_S3_BUCKET }}
      cloudfront_id: ${{ vars.AWS_CLOUDFRONT_ID }}
      slack_webhook_url: ${{ secrets.SLACK_WEBHOOK_URL }}
```

> **💡 Tip:** If you prefer to see all steps written out explicitly (without composite actions), check the sample
> projects in `projects/` directory. They contain fully expanded workflows that are easier to understand and modify.

---

## 📁 Directory Structure

```
.github/
└── actions/                         # Reusable composite actions
    ├── setup-environment/           # Combined: GitHub auth + AWS + language setup (backend)
    ├── setup-frontend-environment/  # Combined: GitHub auth + AWS + Node.js setup
    ├── build-and-package/          # Combined: Build + Docker operations
    ├── build-frontend-bundle/      # Frontend build
    ├── build-and-test/            # Combined: Setup + build + lint + test
    ├── deploy-and-notify/         # Combined: Deploy + extract git info + Slack
    │
    ├── setup-github-auth/         # Individual: GitHub authentication
    ├── setup-aws-credentials/    # Individual: AWS credentials
    ├── setup-java-kotlin/         # Individual: Java/Kotlin setup
    ├── setup-python/              # Individual: Python setup
    ├── build-java-kotlin/         # Individual: Java/Kotlin build
    ├── docker-ecr-login/         # Individual: Docker ECR login
    ├── docker-prepare-tags/      # Individual: Docker tag preparation
    ├── docker-build-push/        # Individual: Docker build & push
    ├── extract-git-info/         # Individual: Git info extraction
    ├── extract-app-version/      # Individual: App version extraction
    └── send-slack-notification/  # Individual: Slack notifications

templates/
├── frontend/                      # Frontend templates
│   ├── cloudfront/               # CloudFront/S3 deployment
│   ├── amplify/                  # AWS Amplify deployment
│   └── e2e-tests.yml            # End-to-end tests
├── backend/                       # Backend templates
│   ├── kotlin/                   # Kotlin/Java
│   │   ├── pr-check.yml         # PR validation
│   │   ├── eks/                 # EKS/Kubernetes
│   │   │   ├── deploy-dev.yml
│   │   │   └── deploy-prod.yml
│   │   └── ecs/                 # ECS
│   │       ├── deploy-dev.yml
│   │       └── deploy-prod.yml
│   └── python/                   # Python (same structure)
├── infra/                         # Infrastructure (Terraform)
│   ├── pr-check.yml
│   ├── deploy-dev.yml
│   └── deploy-prod.yml
└── best-practices/               # Best practice examples
    ├── BEST_PRACTICES.md
    ├── matrix-multi-environment.yml
    ├── caching-strategies.yml
    └── ...
```

---

## 📋 Variables and Secrets Reference

### Common Variables and Secrets (All Templates)

**Variables:**

- `GH_APP_ID` - GitHub App ID for authentication
- `AWS_REGION` - AWS region for deployments

**Secrets:**

- `GH_APP_PRIVATE_KEY` - GitHub App private key for authentication
- `SLACK_WEBHOOK_URL` - Slack webhook URL for notifications
- `AWS_ROLE_TO_ASSUME_DEV` - AWS IAM role ARN for DEV environment
- `AWS_ROLE_TO_ASSUME_PROD` - AWS IAM role ARN for PROD environment

### Frontend Templates

**CloudFront Deployment:**

**Variables:**

- `AWS_S3_BUCKET_DEV` - S3 bucket name for DEV
- `AWS_CLOUDFRONT_ID_DEV` - CloudFront distribution ID for DEV
- `AWS_S3_BUCKET_PROD` - S3 bucket name for PROD
- `AWS_CLOUDFRONT_ID_PROD` - CloudFront distribution ID for PROD
- `NODE_VERSION` - Node.js version (e.g., `18.18.2`)
- `VITE_*` - Frontend environment variables (project-specific)

**Amplify Deployment:**

**Variables:**

- `NODE_VERSION` - Node.js version

**Secrets:**

- `AMPLIFY_WEBHOOK_URL_DEV` - AWS Amplify webhook URL for DEV
- `AMPLIFY_WEBHOOK_URL_PROD` - AWS Amplify webhook URL for PROD

### Backend Templates

**Common (All Backend):**

**Variables:**

- `SERVICE_NAME` - Service name for deployments
- `BASE_DOMAIN_DEV` - Base domain for DEV environment
- `BASE_DOMAIN_PROD` - Base domain for PROD environment

**Secrets:**

- `AWS_ACCOUNT_ID_DEV` - AWS account ID for DEV
- `AWS_ACCOUNT_ID_PROD` - AWS account ID for PROD

**Kotlin/Java Specific:**

**Variables:**

- `JAVA_VERSION` - Java version (e.g., `17`)
- `ARTIFACTORY_URL` - Maven artifact repository URL
- `ARTIFACTORY_USERNAME` - Maven artifact repository username
- `AWS_CODE_ARTIFACT_DOMAIN_DEV` - AWS CodeArtifact domain for DEV

**EKS/Helm Deployment:**

**Variables:**

- `DOCKER_REPO_DEV` - Docker repository URL for DEV
- `DOCKER_REPO_PROD` - Docker repository URL for PROD
- `HELM_REPO` - Helm chart repository URL
- `HELM_SPARTAN_VERSION` - Helm chart version

**ECS Deployment:**

**Variables:**

- `DOCKER_REPO` - Docker repository URL
- `ECS_CLUSTER_NAME_DEV` - ECS cluster name for DEV
- `ECS_CLUSTER_NAME_PROD` - ECS cluster name for PROD
- `SERVICE_NAME_DEV` - ECS service name for DEV
- `SERVICE_NAME_PROD` - ECS service name for PROD
- `TASK_FAMILY_DEV` - ECS task family name for DEV
- `TASK_FAMILY_PROD` - ECS task family name for PROD
- `CONTAINER_NAME` - Container name in task definition

**Python Specific:**

**Variables:**

- `PYTHON_VERSION` - Python version (e.g., `3.11`)

### Infrastructure (Terraform) Templates

**Secrets:**

- `TF_BACKEND_CONFIG_DEV` - Terraform backend configuration for DEV (optional)
- `TF_BACKEND_CONFIG_PROD` - Terraform backend configuration for PROD (optional)

**Variables:**

- `TF_VAR_*` - Terraform input variables (project-specific)

---

## ❓ Common Questions

### Which deployment strategy should I choose?

- **Default**: Workflows deploy directly from `main`/`master` branch to DEV and from git tags to PROD
- **Release branches**: If you need a trunk-based release strategy, consider using release branches (e.g.,
  `release/v1.0.0`). You can modify the workflow trigger to include `release/*` branches in the `on.push.branches`
  section of your workflow file.

### EKS vs ECS?

- **EKS**: Choose if you already have Kubernetes clusters or need Kubernetes features
- **ECS**: Choose if you want simpler AWS-native container orchestration

### How do I deploy to multiple environments?

- Use **matrix strategies** to deploy to multiple environments in parallel
- See [best-practices/BEST_PRACTICES.md](best-practices/BEST_PRACTICES.md) for detailed guidance
- Matrix strategies allow you to build once and deploy to multiple environments simultaneously

---

## 🔧 Troubleshooting

### Workflow fails with "Permission denied"

- Check that all required secrets are set
- Verify `GH_APP_ID` and `GH_APP_PRIVATE_KEY` are correct
- Ensure AWS roles have correct permissions

### Docker build fails

- Check Dockerfile path in workflow
- Verify build context is correct
- Check if all required files exist

### Deployment fails

- Verify all environment variables are set
- Check AWS credentials and roles
- Verify cluster/service names are correct
- Check ECR repository exists

### PR check fails

- Review test failures
- Check linting errors
- Verify build commands match your project structure

---

## 🎓 Key Concepts

### Concurrency

Workflows automatically cancel older runs when a new one starts on the same branch. This prevents conflicts.

### Image Promotion

For PROD deployments, we promote the DEV image instead of rebuilding. This ensures consistency and saves time.

### Security

All workflows use GitHub App tokens (not personal tokens) and OIDC for AWS authentication. This is more secure.

### Caching

Workflows cache dependencies to speed up builds. PRs get isolated caches, while main branch refreshes the primary cache.

**GitHub-hosted runners**: Caching is automatically handled by GitHub Actions. Dependencies are cached between runs.

**Self-hosted runners**: Caching works the same way, but cache storage is on your runner's disk. Make sure your runner
has sufficient disk space for caches.

### Runners

Workflows support both GitHub-hosted and self-hosted runners:

- **GitHub-hosted** (`ubuntu-latest`): Default for most workflows. No setup required, but limited resources.
- **Self-hosted** (e.g., `builder-16-32`): Used for resource-intensive builds. Requires setting up your own runner.

To use self-hosted runners, update the `runs-on` value in your workflow file. See the Kotlin templates for examples.

---

## 📚 Additional Resources

- [best-practices/BEST_PRACTICES.md](best-practices/BEST_PRACTICES.md) - GitHub Actions best practices, matrix
  strategies, caching, and optimization patterns
- [best-practices/](./best-practices/) - **Working examples** of best practices with real, runnable workflows
- [GitHub Actions Documentation](https://docs.github.com/en/actions)

---

**Need help?**

- Check the workflow files for `TODO:` comments - these indicate what needs to be customized for your project
- Review the sample projects in `projects/` to see fully expanded workflows with all steps written out
- Use templates from `templates/` if you prefer modular workflows with composite actions
