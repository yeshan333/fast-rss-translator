#!/bin/sh

while [ $# -gt 0 ]; do
  case "$1" in
    --pattern=*)
      pattern="${1#*=}"
      ;;
    --update_file=*)
      update_file="${1#*=}"
      ;;
    --username=*)
      username="${1#*=}"
      ;;
    --push=*)
      push="${1#*=}"
      ;;
    --org=*)
      org="${1#*=}"
      ;;
    --repo=*)
      repo="${1#*=}"
      ;;
    --token=*)
      token="${1#*=}"
      ;;
    *)
      printf "***************************\n"
      printf "* Error: Invalid argument.*\n"
      printf "***************************\n"
      exit 1
  esac
  shift
done

ls -al /usr/bin/ >> running.log

/bin/fast-rss-translator --update-file "$update_file" >> running.log

if [ $? -eq 0 ]
then
  cat running.log
else
  echo "Update rss feed files failed"
  cat running.log
  exit 1
fi

if [ "$push" = "true" ]
then
  git config --local user.email "${username}@users.noreply.github.com"
  git config --local user.name "${username}"
  git status -s
  git add .

  git commit -m "Auto commit by bot, ci skip"
  git push https://${username}:${token}@github.com/${org}/${repo}.git
fi
