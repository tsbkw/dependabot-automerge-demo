# dependabot-automerge-demo
demo of https://docs.github.com/en/code-security/dependabot/working-with-dependabot/automating-dependabot-with-github-actions#enable-auto-merge-on-a-pull-request

In this repository, auto-merge on a pull request feature is configured.

What I want to implement is a function that merge pull requests created by dependabot automatically if and only if
 * dependency update is patch update (=low risk to break current behavior)
 * all checks are succressfully met


Note: 
This setting is dangerous if some of dependencies don't follow semantic versioning.
If only specific dependency don't follow, you can filter them like below.
Unfortunately in some ecosystem many of dependencies not following semantic version.

```
      - name: Enable auto-merge for Dependabot PRs
        if: !${{contains(steps.metadata.outputs.dependency-names, 'dependency-not-follow-semver') && steps.metadata.outputs.update-type == 'version-update:semver-patch'}}
        run: gh pr merge --auto --merge "$PR_URL"
```

