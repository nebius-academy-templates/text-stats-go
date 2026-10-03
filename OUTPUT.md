# Lesson 7 — final result format

Use this file only at the final export, after the bounded read-only lookup and
the learner's source check, or after recording why that task was unavailable.
It is a format instruction, not an answer to the Git question.

Write root `report.json` when explicitly asked. Read-only commands to read this
instruction and inspect local Git state are permitted during export. Do not
execute the Git examples, run the application, repeat the external lookup,
add MCP configuration, modify application files, commit, or push.

Ask the learner for policy and independent-source observations not visible in
the conversation. A tool response is not proof that the learner opened its
source. Never claim a call, approval, or source verification that did not occur.

Use UTF-8 JSON, `schema_version: 1`, and `lesson: 7`. Keep machine keys and
choices unchanged; free text may use the learner's language. No Markdown fences,
comments, duplicate keys, or placeholders. The empty example describes shape:

```json
{
  "schema_version": 1,
  "lesson": 7,
  "outcome": "",
  "why_external": "",
  "policy": {"decision": "", "basis": ""},
  "provider": "Context7",
  "library": "/git/htmldocs",
  "tool": null,
  "model_conclusion": null,
  "verification": null,
  "unavailable": null
}
```

## Verified outcome

Set `outcome` to `verified` only when the permitted lookup and independent source
check were actually completed. `unavailable` stays null. Fill these objects:

```json
{
  "tool": {"name": "", "query": "", "library": "", "excerpt": "", "source_url": ""},
  "model_conclusion": "",
  "verification": {"source_url": "", "source_observation": "", "verdict": "", "lowercase_c": "", "uppercase_c": "", "learner_confirmed": true}
}
```

- `why_external`: why the repository cannot supply the needed evidence.
- `policy.decision`: `approved`, `not-approved`, or `unknown`; `basis`: the learner's actual applicable policy decision. A connected provider does not itself mean policy approval.
- `tool`: the visible call name, actual query, resolved library, relevant returned excerpt, and exact returned source URL. Quotes may be short; do not include credentials or private data.
- `model_conclusion`: what the model actually answered, even if it was wrong.
- `verification.source_url`: the Git documentation page independently opened by the learner; not a search-result page or a secondary blog.
- `verification.source_observation`: what the learner found in the relevant option text, including any qualification; this can be a paraphrase.
- `verification.verdict`: `supported` or `contradicted`. If evidence is still insufficient, use the unavailable outcome instead of claiming verification.
- `lowercase_c` and `uppercase_c`: source-checked results for an already existing branch. Each value is one of `refuses-existing-branch`, `resets-to-start-point`, `switches-without-reset`, or `unknown`. Do not substitute the model's unsupported answer for the learner's checked conclusion.
- `learner_confirmed`: true only when the learner has actually supplied the independent source-check result; otherwise ask for it.

## Unavailable outcome

Set `outcome` to `unavailable`. `verification` must be null; `tool` and
`model_conclusion` may be null or contain actual partial observations. Fill:

```json
{
  "unavailable": {"step": "", "observation": "", "checked": "", "gap": "", "next_step": ""}
}
```

`step` is `policy`, `connection`, `library`, `tool-call`, `source`, or
`verification`. State the observed limitation, checks actually made, the
remaining evidence gap, and a safe next step within the original boundary.
If policy is unknown or not approved, stop at `policy`; do not claim a call.
An unavailable record is not a verified answer and must not invent one.

## Learner handoff

Ask the learner to verify the saved report, then commit only report.json.
The checker validates structure, source scope, option conclusions for verified
records, and unchanged project files. It does not contact Context7, open URLs,
authenticate UI history, or fully evaluate the explanatory prose.
