# go-example-tigris

Forked from the minimal Go application [repository](https://github.com/fly-apps/go-example) for [fly.io Getting Started](https://fly.io/docs/getting-started/golang/).Includes setup for setting up a Tigris connection using AWS SDK for Go.


To get started:

1. clone this repo
2. `flyctl launch`
3. view the deployed app with flyctl open
4. Create a Tigris Bucket by running `fly storage create` in the base directory of the repo. 
5. Update the Bucket name specified in the `UploadText()` function to use the Bucket name of your Tigris App
6. Run `fly deploy`
