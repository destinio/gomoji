# Welcome to Gomoji!

Gomoji is a simple emoji CLI tool to quickly find and copy emojis to your clipboard.

## Installation

```bash
./install.sh
```

```bash
make install
```

## Usage

```bash
# Find all emojis that match the search term
# This only return the default emoji
# (Not all emojis have a default e.g. doctor)
gomoji [search term]

# Find all emojis that match the search term
# This returns all emojis including genders and skin tones
gomoji --all [search term]
```

## Examples

```bash
gomoji golf

🏌 U+1F3CC :person_golfing:
⛳ U+26F3 :golf:

gomoji --all golf

🏌 U+1F3CC :person_golfing:
🏌🏻 U+1F3CC U+1F3FB :person_golfing_tone1:
🏌🏼 U+1F3CC U+1F3FC :person_golfing_tone2:
🏌🏽 U+1F3CC U+1F3FD :person_golfing_tone3:
🏌🏾 U+1F3CC U+1F3FE :person_golfing_tone4:
🏌🏿 U+1F3CC U+1F3FF :person_golfing_tone5:
🏌️‍♂️ U+1F3CC U+FE0F U+200D U+2642 U+FE0F :man_golfing:
🏌🏻‍♂️ U+1F3CC U+1F3FB U+200D U+2642 U+FE0F :man_golfing_tone1:
🏌🏼‍♂️ U+1F3CC U+1F3FC U+200D U+2642 U+FE0F :man_golfing_tone2:
🏌🏽‍♂️ U+1F3CC U+1F3FD U+200D U+2642 U+FE0F :man_golfing_tone3:
🏌🏾‍♂️ U+1F3CC U+1F3FE U+200D U+2642 U+FE0F :man_golfing_tone4:
🏌🏿‍♂️ U+1F3CC U+1F3FF U+200D U+2642 U+FE0F :man_golfing_tone5:
🏌️‍♀️ U+1F3CC U+FE0F U+200D U+2640 U+FE0F :woman_golfing:
🏌🏻‍♀️ U+1F3CC U+1F3FB U+200D U+2640 U+FE0F :woman_golfing_tone1:
🏌🏼‍♀️ U+1F3CC U+1F3FC U+200D U+2640 U+FE0F :woman_golfing_tone2:
🏌🏽‍♀️ U+1F3CC U+1F3FD U+200D U+2640 U+FE0F :woman_golfing_tone3:
🏌🏾‍♀️ U+1F3CC U+1F3FE U+200D U+2640 U+FE0F :woman_golfing_tone4:
🏌🏿‍♀️ U+1F3CC U+1F3FF U+200D U+2640 U+FE0F :woman_golfing_tone5:
⛳ U+26F3 :golf:
```

## Coming Soon
- [ ] Add support for configuring gender and skin tone
- [ ] Add support for favorite emojis