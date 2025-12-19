# File: config.gd
# Purpose: Global configuration singleton. Contains DEV_MODE flag to toggle
# development-only features like reset buttons and debug endpoints. Set to
# false for production builds to automatically hide all dev UI. Single source
# of truth for environment configuration across all scenes.
# Path: client/scripts/autoload/config.gd
# All Rights Reserved. Arc-Pub.

extends Node

## Set to false for production builds
const DEV_MODE: bool = true

## Backend URL (change for production)
const API_URL: String = "http://localhost:8080/api/v1"
