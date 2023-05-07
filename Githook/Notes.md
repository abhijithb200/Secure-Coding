```sh
#!/usr/bin/env bash
set -e

#=== 'prev-commit' solution by o_O Tync
#commit_hash=$(git rev-parse --verify HEAD)
commit=$(git log -1 --pretty="%H%n%ci") # hash \n date
commit_hash=$(echo "$commit" | head -1)
commit_date=$(echo "$commit" | head -2 | tail -1) # 2010-12-28 05:16:23 +0300

branch_name=$(git symbolic-ref -q HEAD) # http://stackoverflow.com/questions/1593051/#1593487
branch_name=${branch_name##refs/heads/}
branch_name=${branch_name:-HEAD} # 'HEAD' indicates detached HEAD situation

# Write it


go run ../layer2/diffmain.go $commit_hash

echo "written"
```
