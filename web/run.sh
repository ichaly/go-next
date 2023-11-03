#!/bin/bash

# This script is used to run the development server.
# It is intended to be used with the "dev" command in the "bun" tool.
# It is not intended to be run directly.

set -e

# Set the NODE_ENV to "development"
export NODE_ENV=development

bun --bun run dev