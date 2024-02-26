#!/usr/bin/env bash
set -xe
# echo $PWD
# echo $HOME
. "$HOME/.asdf/asdf.sh"

export PATH=$PATH:/usr/local/texlive/2023/bin/x86_64-linux

SUCCESS_SYMLINK="last_built.pdf"

echo "Removing out/..."
rm -rf out
echo "Removing symlink"
rm -f $SUCCESS_SYMLINK
echo "Removed."
echo "Running generate step - this will take a while..."
PLANNER_CONFIG="config/kindlescribe_left-hand_dotted.yaml"
time bundle exec latex_yearly_planner generate $PLANNER_CONFIG
echo Finished.
PLANNER_NAME="$(basename $PLANNER_CONFIG .yaml)"
echo "Renaming output index.pdf to $PLANNER_NAME.pdf"
mv out/index.pdf out/$PLANNER_NAME.pdf
ln -s out/$PLANNER_NAME.pdf $SUCCESS_SYMLINK
echo "Build done, output:"
echo "out/$PLANNER_NAME.pdf"