# GCFLodestone

A set of ready-to-deploy Google Cloud Functions to parse Lodestone. 

## How to use

First of all, make sure you have a Google Cloud account with a local CLI available (https://cloud.google.com/sdk)

Deploy the desired function using:

`gcloud functions deploy <The name you want to give to your function> --runtime go113 --trigger-http --allow-unauthenticated`
