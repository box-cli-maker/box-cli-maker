#!/usr/bin/env bash
# Regenerate every showcase screenshot in img/ (and the hero) from source.
set -euo pipefail
cd "$(dirname "$0")/../.."

subjects=(single single_double double double_single bold round hidden classic block
  top bottom top_center top_right bottom_center bottom_right
  inside_left inside_right left right padding margin wrap ansi_safe)

for s in "${subjects[@]}"; do
  ./scripts/screenshots/shoot.sh "go run ./scripts/screenshots/showcase $s" "img/$s.png" --plain
  echo "img/$s.png"
done

# The version-labelled variant of the ANSI-safety figure carries shorter
# labels, so it is padded to the first one's dimensions — the two are shown
# in different documents and should not differ in size.
./scripts/screenshots/shoot.sh "go run ./scripts/screenshots/showcase ansi_safe_versions" \
  img/ansi_safe_versions.png --plain --match img/ansi_safe.png
echo "img/ansi_safe_versions.png"

./scripts/screenshots/shoot.sh "go run ./examples/readme/" img/hero.png
echo "img/hero.png"
