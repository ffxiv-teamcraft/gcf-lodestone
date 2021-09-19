# GCFLodestone

A set of ready-to-deploy Google Cloud Functions to parse Lodestone. 

## How to use

First of all, make sure you have a Google Cloud account with a local CLI available (https://cloud.google.com/sdk)

Deploy the desired function using:

`gcloud functions deploy <name of the function to deploy (LodestoneCharacter, SearchCharacter, ...)> --runtime go113 --trigger-http --allow-unauthenticated`
