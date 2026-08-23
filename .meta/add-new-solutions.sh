#!/usr/bin/env bash

set -euo pipefail

files=()
while IFS= read -r file; do
    [[ -n "$file" ]] && files+=("$file")
done < <(
    {
        git diff --diff-filter=A --cached --name-only -- 'solutions/[0-9]*.go'
        git ls-files --others --exclude-standard -- 'solutions/[0-9]*.go'
    } | sort -u
)

if [[ ${#files[@]} -eq 0 ]]; then
    echo "No new staged or untracked problem files found in solutions/"
    exit 0
fi

count=${#files[@]}
problem_lines=""

for file in "${files[@]}"; do
    filename=$(basename "$file")

    number="${filename%%.*}"
    without_num="${filename#*.}"
    shortname="${without_num%.go}"

    problem_lines+="- ${number} ${shortname}"$'\n'
done

commit_msg="problems: Add ${count} new solution(s)

new solutions:
${problem_lines}"

git add "${files[@]}"
git commit -m "$commit_msg" --only "${files[@]}"
