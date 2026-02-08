# Integrations

## Powerlevel10k Prompt Integration

If you use [Powerlevel10k](https://github.com/romkatv/powerlevel10k), you can display the active gt account in your prompt. Add the following to your `~/.p10k.zsh`:

```zsh
function prompt_gt_account() {
  local cfg="$HOME/.config/gt/public.yml"
  [[ -r "$cfg" ]] || return

  local user platform icon icon_color text_color

  user="$(sed -nE 's/^user:[[:space:]]*(.+)$/\1/p' "$cfg")"
  platform="$(sed -nE 's/^platform:[[:space:]]*(.+)$/\1/p' "$cfg")"

  [[ -z "$user" || -z "$platform" ]] && return

  case "$platform" in
    GitHub|github)
      icon_color=39
      icon=''
      ;;
    GitLab|gitlab)
      icon_color=208
      icon=''
      ;;
    *)
      icon_color=240
      icon=''
      ;;
  esac

  # Adaptive text color
  if [[ $P9K_BACKGROUND == light ]]; then
    text_color=238   # dark gray (good on white)
  else
    text_color=252   # light gray (good on dark)
  fi

  p10k segment -f $icon_color -i "$icon"
  p10k segment -f $text_color -t "$user"
}
```

Then add `gt_account` to your `POWERLEVEL9K_LEFT_PROMPT_ELEMENTS` or `POWERLEVEL9K_RIGHT_PROMPT_ELEMENTS` array:

```zsh
typeset -g POWERLEVEL9K_RIGHT_PROMPT_ELEMENTS=(
  # ... your other segments ...
  gt_account
)
```
