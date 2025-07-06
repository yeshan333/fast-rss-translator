#!/bin/sh

while [ $# -gt 0 ]; do
  case "$1" in
    --config_file=*)
      config_file="${1#*=}"
      ;;
    --update_file=*)
      update_file="${1#*=}"
      ;;
    --cloudflare_account_id=*)
      cloudflare_account_id="${1#*=}"
      ;;
    --cloudflare_api_key=*)
      cloudflare_api_key="${1#*=}"
      ;;
    *)
      printf "***************************\n"
      printf "* Error: Invalid argument.*\n"
      printf "***************************\n"
      exit 1
  esac
  shift
done

export CLOUDFLARE_ACCOUNT_ID="${cloudflare_account_id}"
export CLOUDFLARE_API_KEY="${cloudflare_api_key}"
/bin/fast-rss-translator --config "$config_file" --update-file "$update_file" >> running.log

if [ $? -eq 0 ]
then
  cat running.log
else
  echo "Update rss feed files failed"
  cat running.log
  exit 1
fi

git branch
git remote -v
git status -s
