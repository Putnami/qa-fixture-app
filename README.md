# qa-fixture-app

A Putnami workspace. Each project lives in its own directory and is listed in
putnami.workspace.json.

## Commands

```bash
putnami lint,test,build --impacted   # check the projects a change touches
putnami projects list                # list the projects
putnami projects create <name> --template <template>
```

## Documentation

Each project documents its purpose and its commands in its own README.md.
AGENTS.md tells coding agents how to work in this workspace.
