# Brew CI files

Files in this directory are all related to `brew` package manager CI integration.

## Brew Release

For `brew` CI to complete, you **must** use a tag that has an existing branch. The brew package file will be commited to that branch. If you don't, the brew release will fail.

## Token

Your must create a token named `CI_BREW_TOKEN` with a **Developer** role and the following permissions: `read_api`, `read_repository`, `write_repository`.

You must then create CI/CD variable named `CI_BREW_TOKEN` and activate the **Masked** option.

## Pipelines

The last part of the automation create a new commit with the final brew installation file. **If this triggers a pipeline** use the `-o ci.skip` with `git push`. Otherwise, you'll have inifitine CI loop.
