#!/bin/sh
set -eu
echo("run core hooks")
hook_name="$(basename "$0")"
hook_script=".git/hooks/$hook_name"
[ -e "$hook_script" ] && "$hook_script"
exit 0