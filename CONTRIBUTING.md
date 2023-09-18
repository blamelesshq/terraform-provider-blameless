# Welcome to the Blameless Terraform Provider contributing guide

Thank you for your interest in contributing to our project!

Read our [Code of Conduct](./CODE_OF_CONDUCT.md) to keep our community approachable and respectable.

In this guide you will get an overview of the contribution workflow from opening an issue, creating a PR, reviewing, and merging the PR.

### Issues

#### Create a new issue

If you find a problem or have a suggestion, check if there is already an existing issue. If the isn't an existing issue, open a new issue and fill out the issue template with all of the relevant information.

#### Solve an issue

Scan through our existing issues to find one that interests you. If you find an issue to work on, you are welcome to open a PR with a fix.

### Local Testing

1. Run `make install` in the root directory
2. Go to the `modules` directory
3. Run `terraform init`
4. Run `terraform plan`, `terraform validate`, or `terraform apply`

### Documentation

Before committing your changes, remember to run `make doc` to capture any changes to the provider or resources in the documentation.

### Pull Request

When you're finished with the changes, create a pull request, also known as a PR.
- Fill the PR template so that we can review your PR. This template helps reviewers understand your changes as well as the purpose of your pull request.
- Don't forget to [link PR to issue](https://docs.github.com/en/issues/tracking-your-work-with-issues/linking-a-pull-request-to-an-issue) if you are solving one.
- Enable the checkbox to [allow maintainer edits](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/allowing-changes-to-a-pull-request-branch-created-from-a-fork) so the branch can be updated for a merge.
Once you submit your PR, a team member will review your proposal. We may ask questions or request additional information.
- We may ask for changes to be made before a PR can be merged, either using suggested changes or pull request comments. 
- As you update your PR and apply changes, mark each conversation as resolved.

### Your PR is merged!

Congratulations :tada::tada: The Blameless team thanks you :sparkles:.
