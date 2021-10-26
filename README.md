# Cloud Run Router

This is a work in progress (probably a long way off)

# Goal

To provide a simple means of using Google Cloud Run as a reverse proxy to Cloud Run Services and Cloud Functions.

# Configuration

## Environment variables

**EXPOSE_ROUTES** - Boolean, default = false.  If true, the route config will be shown at /_routes

**ENABLE_FIREBASE_CONFIG** - Boolean, default = false.  If true, the routes configuration will be loaded from FIREBASE_CONFIG_PATH

**FIREBASE_CONFIG_PATH** - String, default = null. Required if ENABLE_FIREBASE_CONFIG is true.

**ROUTE_CONFIG** - JSON string, default = null.  Required if ENABLE_FIREBASE_CONFIG is false.

