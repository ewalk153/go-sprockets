go-sprockets
============

This is an example project to show how you can use ruby sprockets as part of your development environment when building a Go web application.

There are still many things to tackle:
1. the sprockets assets manager should probably be seperated into its own package for reuse.
2. there is not build out to be able to reference a static manifest file (would be useful for production environments)
3. this should have some sense of which enviornment it's running in and change the loading behavior based on that
4. it would be great if the hashed asset name references would change automatically any time the assets change (lazy update references in go)
