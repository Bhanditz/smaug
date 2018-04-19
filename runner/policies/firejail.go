package policies

// This templates generates a sandbox policy file suitable for
// running relatively-untrusted apps via itch.
//
// TODO: figure a better way — blacklists aren't so good.
// whitelist doesn't seem to work with exclusions, though?

const FirejailTemplate = `
blacklist ~/.config/itch/users
blacklist ~/.config/itch/butler_creds
blacklist ~/.config/itch/marketdb
blacklist ~/.config/itch/Cookies
blacklist ~/.config/itch/Partitions

blacklist ~/.config/kitch/users
blacklist ~/.config/kitch/butler_creds
blacklist ~/.config/kitch/marketdb
blacklist ~/.config/kitch/Cookies
blacklist ~/.config/kitch/Partitions

blacklist ~/.config/chromium
blacklist ~/.config/chrome
blacklist ~/.mozilla
`
