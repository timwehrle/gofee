<pre>
_________     ________         
__  ____/________  __/_________ 
_  / __ _  __ \_  /_ _  _ \  _ \
/ /_/ / / /_/ /  __/ /  __/  __/
\____/  \____//_/    \___/\___/ 
</pre>

# Gofee - The Go Password Generator

## Overview

Gofee is a secure CLI password generator in Go. This CLI is based on the `crypto/rand` Go package, which leverages the system's random number generator (RNG).

### What does the name mean?

Originally, the name results out of "Go" + "Safe" = "Gofe", but since this name was already taken, an additional "e" was added, resulting in "Gofee".

### Why another Password Generator?

While there are already many password generators available, we felt the need for one that improves usability and customizability. Some existing generators lack flexibility, which is why we created Gofee.

You can also contribute additional features ([Contributing Guidelines][contributing]).

## Building and Usage

## Security and Data Protection

Gofee uses a random number generator to ensure that the generated passwords have no discernible patterns and are free from bias, making them highly secure against attacks.

### Security Recommendations

For optimal security, always use long, complex passwords. Once you have generated a password, ensure you copy it securely and store it in a reputable password manager. This practice will help protect your credentials from unauthorized access and make it easier to manage them effectively.

## Contributing

If something feels off, you see an opportunity to improve performance, or think some functionality is missing, we’d love to hear from you! Please review our [contributing docs][contributing] for detailed instructions on how to provide feedback or submit a pull request. Thank you!

[contributing]: ./.github/CONTRIBUTING.md
