#!/bin/bash

buildScript/lib/core/init.sh
cd ..
rm -rf sing-box
rm -rf sing
git clone -b sing-box-b-t-h https://github.com/White12352/sing-box.git sing-box
git clone -b dev-tl https://github.com/White12352/sing.git sing
cd sjnh/libcore
go mod tidy
cd ..
buildScript/lib/core/build.sh
