#!/usr/bin/env bash

swagger generate server -f swagger.json -t .
swagger generate client -f swagger.json -t .