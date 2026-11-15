# qa-fixture-app

The repository of the `QA-fixture` workspace on Putnami. It exists so the
console has representative data to show: changes, CI runs and verdicts in
several states, and missions.

Its content comes from `putnami operator qa-fixture seed`, which pushes it
from the template in the `putnami-cloud` repository. A run brings every
branch back to the template: a commit made here by hand is replaced.

| Branch | Pull request | What its CI does |
| --- | --- | --- |
| `main` | | passes |
| `qa/passing-change` | adds a farewell | passes |
| `qa/failing-test` | stops trimming the name | fails on purpose: a test expects the trim |
| `qa/declaration-change` | capitalizes the name and declares it | passes, and waits for the declaration's approval |
| `qa/draft` | documents how to run the tests | a draft: no run |

Run the gate locally with `putnami lint,test,build,validate --all`.
