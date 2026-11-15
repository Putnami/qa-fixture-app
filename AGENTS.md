# AGENTS.md

<!-- putnami:guidance v2 begin sha256:a960af4ac268c69ba83ed81d56f6438b624a702764e27ea7f4498d8d78f9d3c4 -->
This is a Putnami workspace.
- Build, test, lint, install, add dependencies, and generate with `./putnamiw` when present, otherwise `putnami`. Never call npm, bun, go, pip, or cargo directly for those.
- While iterating, select the projects you changed with `--projects <a>,<b>`; unlike `--impacted`, it skips every project that depends on them. Before declaring work complete, run `putnami lint,test,build,validate --impacted --enforce-coverage` once, when feasible.
- Discover code through the CLI, then grep: `putnami projects list|describe|owner|deps|why` for the project graph, `putnami context map|pack` to orient on the workspace or a project, and `putnami <verb> --impacted --dry-run` for what a change affects. Use grep for literal text.
- Skills and agents that an extension's agent content installs under `.agents/`, `.claude/` and `.codex/` are managed by Putnami. To customize one, copy it under another name; edited managed files are preserved and reported, never overwritten.
- Read `.agents/constraints.md` when it exists.
<!-- putnami:guidance v2 end -->

## Repository rules

This repository is the fixture workspace for Putnami test accounts. Keep it small and representative.

- `services/greeter` is a Go HTTP service. `libs/greeting` is the Go library it calls.
- Run `putnami lint,test,build,validate --impacted --enforce-coverage` before you push.
- `putnami.ci.json` runs the same commands on pushes to `main` and on pull requests.
- Commit messages and pull request bodies name no agent and no model.
