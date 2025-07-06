#!/bin/sh

while [ $# -gt 0 ]; do
  case "$1" in
    --config_file=*)
      config_file="${1#*=}"
      ;;
    --update_file=*)
      update_file="${1#*=}"
      ;;
    *)
      printf "***************************\n"
      printf "* Error: Invalid argument.*\n"
      printf "***************************\n"
      exit 1
  esac
  shift
done

/bin/fast-rss-translator --config "$config_file" --update-file "$update_file" >> running.log
if [ $? -eq 0 ]
then
  cat running.log
else
  echo "Update rss feed files failed"
  cat running.log
  exit 1
fi

git config --global --add safe.directory /github/workspace
git branch
git remote -v
git status -s
