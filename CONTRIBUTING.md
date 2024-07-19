# Contributing

Our deployment process is powered by GitHub Actions.

## After Merging to Main

Once your changes have been merged into the main branch:

### Updating Arrow Files

If you've made changes to the arrow files, follow these steps:

#### Creating and Publishing a Docker Image

1. **Find the relevant workflow**: Look for "Create and publish a Docker image for Otel Collectors with Arrow".
2. **Run the workflow**: Click on 'Run workflow' to initiate the Docker image creation and publishing process.

### Updating Charts

If you've made changes to the chart files, follow these steps:

#### Releasing Helm Chart Changes

For releasing changes to Helm charts, follow these steps:

1. **Find the relevant workflow**: Locate "Release Charts".
2. **Run the workflow**: Click on 'Run workflow' to trigger the release process for Helm charts.
