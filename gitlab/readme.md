# GitLab CI scripts

Files in this directory are bash helper scripts ran from the `script:` directive in the .gitlab-ci.yml.

## Release

Here are the steps to publish a new release.

First, `git push` to whatever branch your new release is in. Then, create a new tag with `git tag -a tag_name`. Tag name should be short, like `1.0.0` or `1.1.0`.

Git will then open your editor so that you can enter the tag message. The first line will become the release title, and the rest will become the release message.

Here's what it should look like:

```
Release 1.0.0 - First release
This is the first release off my awesome software.

Changelog:

  - #1: This is the first issue that we fixed. 
  - #2: Also this feature request.
  - #3: And this this.
  - #4: And that.

Which conclude this release!
```

Once your message is saved, push your tag with `git push --tags` then wait for the pipeline to finish and review your project release page to see if it worked as expected.
