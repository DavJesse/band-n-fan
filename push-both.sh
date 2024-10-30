#!/bin/sh
git push gitea "$1"
git push github "$1"
git push search "$1"
git push visual "$1"