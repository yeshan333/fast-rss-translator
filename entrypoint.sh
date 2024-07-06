#!/bin/sh

while [ $# -gt 0 ]; do
  case "$1" in
    --pattern=*)
      pattern="${1#*=}"
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

fast-rss-translator --update-file "$update_file" > "$output"

if [ $? -eq 0 ]
then
  cat "$output"
else
  echo "Generate $output failed"
  cat "$output"
  exit 1
fi
