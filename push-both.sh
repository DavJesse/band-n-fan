#!/bin/sh
git push gitea "$1"
git push github "$1"
gitea push search "$1"