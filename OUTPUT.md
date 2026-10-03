# Lesson 5 — final result format

Use this file only after the independent review and the learner's verification.
It specifies report formatting, not findings to look for. Do not read the
learner's human-first notes before completing the independent review.

On an explicit final-export request, read the supplied observations and write
root `report.json`. Read-only commands may inspect this instruction, repository,
and Git state. Do not modify the candidate, install/resolve modules, run the
experimental renderer, create probe files, commit, push, approve, or merge.
Request missing observations. Do not invent model findings, external checks,
learner decisions, command results, or evidence you have not received.

Write UTF-8 JSON with `schema_version: 1`, `lesson: 5`, fixed keys below, and
free-text explanations in the learner's language. No Markdown fences, comments,
duplicate keys, or placeholders. Use the actual data, not the empty example:

```json
{
  "schema_version": 1,
  "lesson": 5,
  "pr": {"url": "", "base": "", "head": "", "changed_files": []},
  "human_questions": [{"question": "", "observation": ""}],
  "model_findings": [
    {"id": "F1", "title": "", "location": {"path": "", "line": 1, "quote": ""}, "assessment": "", "reason": ""}
  ],
  "missed_findings": [{"title": "", "evidence": ""}],
  "checks": {
    "normal": {"command": "", "exit_code": 0, "output": ""},
    "probes": {"command": "", "exit_code": 1, "output": "", "results": {}}
  },
  "risks": [
    {"class": "", "location": {"path": "", "line": 1, "quote": ""}, "execution_path": "", "evidence": "", "consequence": "", "mitigation": "", "certainty": ""}
  ],
  "contributor_comment": "",
  "decision": "",
  "uncertainty": ""
}
```

## Field rules

- `pr`: the field name is retained for compatibility, but this exercise reviews a branch comparison, not an existing GitHub PR. Set `url` to `https://github.com/nebius-academy-templates/text-stats-codex-go/compare/codex/lesson-5-base...codex/lesson-5`, `base` to `codex/lesson-5-base`, and `head` to `codex/lesson-5`. Record the exact changed-file list from that comparison, not the later submission commit. Do not invent a PR number, author, open status, or posted review.
- `human_questions`: the learner's four questions and initial observations; collect them only after the independent review.
- `model_findings`: one entry per actual finding, with a distinct ID, changed location, and learner assessment (`Useful`, `Weak`, or `False`) plus evidence-based reason. If the model returned nothing, use `[]`; do not invent findings or call an empty run a miss without independent evidence.
- `missed_findings`: only separately verified omissions from that actual model list. Use `[]` if there were none. An unverified concern belongs in `uncertainty`, not an invented category.
- Every `location` contains a repository-relative `path`, one-based `line`, and exact nonempty code excerpt starting on that line in the reviewed candidate. Use a short relevant excerpt, not just a filename. If submitting from the base, cite the separately reviewed candidate, not base line numbers.
- `checks.normal`: actual approved offline command, integer exit code, and observed output. `checks.probes`: the actual targeted command, integer exit code, observed output, and one entry per probe in `results`, using the test name as key and `PASS` or `FAIL` as value. Record what happened; investigate a mismatch with the exercise instead of editing results to satisfy the checker.
- `risks`: one entry for each class covered by the practice: `correctness`, `resource_safety`, `dependency_api`, `provenance`. Explain the reachable path (or optional/installation path), evidence, consequence, and requested mitigation. `certainty` is `confirmed` or `uncertain`; an evidence gap is not proof of a security or licensing violation. These are the learner's final verified conclusions, not necessarily four model findings.
- `contributor_comment`: the learner-checked review draft from the practice. Saving it does not post it to GitHub.
- `decision`: the learner's conclusion, `mergeable` or `not-mergeable-yet`.
- `uncertainty`: remaining limitations, or an explicit statement that no additional uncertainty was recorded. Do not treat all four classes as automatically proven defects.

## Learner handoff

Ask the learner to compare the report with the actual review and command output,
then commit only report.json on their submission branch. The checker verifies
scope, structured observations, citation locations, and intact files; it cannot
authenticate external actions or fully judge the prose. No code changes are
part of this exercise's solution.
