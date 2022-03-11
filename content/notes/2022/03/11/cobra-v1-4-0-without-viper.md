---
title: "Cobra v1.4.0 without Viper"
likeOf: "https://github.com/spf13/cobra/releases/tag/v1.4.0"
date: 2022-03-11T21:33:34+0100
tags:
- golang
---
Yesterday, the Cobra maintainers released version 1.4.0 of the popular command-line library for Go. This release is a huge one as it removes the configuration library Viper as a dependency which also means that you no longer pull dozens of indirect dependencies due to that when using Cobra! The recently released [muesli/coral](https://github.com/muesli/coral) already has a note about possibly moving back to Cobra due to that change 🙂