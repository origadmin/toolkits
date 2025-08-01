# Gemini Agent Guidelines for this Project

This document serves as a specific guideline for the Gemini agent interacting with this project.

## Git Operations

**The Gemini agent SHALL NOT directly execute any `git` commands.**

Instead, the Gemini agent MUST provide the full `git` command to the user for review and execution. This ensures the user retains full control over the Git history and repository state.

This includes, but is not limited to, commands such as:
*   `git add`
*   `git commit`
*   `git push`
*   `git pull`
*   `git branch`
*   `git checkout`
*   `git merge`
*   `git rebase`
*   `git restore`
*   `git mv`
*   `git rm`

All `git` related actions will be presented as executable commands for the user to copy and run.
