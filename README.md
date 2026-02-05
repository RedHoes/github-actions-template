# GitHub Actions Workflow Templates

> **Reusable CI/CD workflows for backend, frontend, and infrastructure projects**

This repository provides production-ready GitHub Actions workflow templates that can be copied and customized for your
projects. All workflows use a modular architecture with composite actions for maintainability and reusability.

## What's Included

- **Workflow Templates** (`templates/`) - Reusable workflows using composite actions
- **Sample Projects** (`projects/`) - Complete examples with fully expanded workflows
- **Composite Actions** (`.github/actions/`) - Reusable workflow steps
- **Best Practices** (`templates/best-practices/`) - Example workflows demonstrating best practices

## Quick Start

1. **Choose your project type:**
    - Backend (Kotlin/Java, Python) → `templates/backend/`
    - Frontend (React, Vue, etc.) → `templates/frontend/`
    - Infrastructure (Terraform) → `templates/infra/`

2. **Copy templates:**
   ```bash
   cp templates/backend/kotlin/pr-check.yml your-project/.github/workflows/
   cp templates/backend/kotlin/ecs/deploy-*.yml your-project/.github/workflows/
   cp -r .github/actions your-project/.github/
   ```

3. **Configure:**
    - Set GitHub secrets and variables (see workflow files for `TODO:` comments)
    - Update project-specific values in workflows

4. **Test:**
    - Create PR → triggers PR checks
    - Push to main → triggers DEV deployment
    - Create tag → triggers PROD deployment

## Key Features

- **3-Step Structure**: Setup → Build → Deploy
- **Multiple Deployment Targets**: EKS, ECS, CloudFront, Amplify, Terraform
- **Security**: OIDC for AWS, GitHub App tokens
- **Notifications**: Automatic Slack notifications
- **Modular**: Composite actions for easy maintenance

## Documentation

For detailed information, see **[GUIDE.md](GUIDE.md)** - Comprehensive guide covering repository structure,
architecture, approach, setup, and best practices.

## Support

- Check workflow files for `TODO:` comments indicating customization points
- Review sample projects in `projects/` for fully expanded examples
- See `templates/best-practices/` for optimization examples
